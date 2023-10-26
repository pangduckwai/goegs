package impl

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"sea9.org/go/neo4j/common"
	"sea9.org/go/neo4j/config"
	"sea9.org/go/neo4j/nodes"
)

// Neo4jConn connection object for neo4j
type Neo4jConn struct {
	// trxReady bool
	err     error
	dbName  string
	ctx     context.Context
	driver  neo4j.DriverWithContext
	session neo4j.SessionWithContext
	trx     neo4j.ExplicitTransaction
}

// Connect connect to the database
func Connect(cfg *config.Config, ctx context.Context, idx ...int) (conn *Neo4jConn, err error) {
	j := 0
	if len(idx) > 0 {
		j = idx[0]
	}

	if j >= len(cfg.DbHost) ||
		j >= len(cfg.DbPort) && len(cfg.DbPort) > 0 ||
		j >= len(cfg.DbName) ||
		j >= len(cfg.DbUser) ||
		j >= len(cfg.Dbpass) && len(cfg.Dbpass) > 0 {
		err = common.NewError(true, "CONN", fmt.Sprintf("Invalid index of database connection params specified: %v", idx))
		return
	}

	conn = &Neo4jConn{}

	if ctx == nil {
		conn.ctx = context.Background()
	} else {
		conn.ctx = ctx
	}

	conn.driver, conn.err = neo4j.NewDriverWithContext(fmt.Sprintf("neo4j://%v:%v", cfg.DbHost[j], cfg.DbPort[j]), neo4j.BasicAuth(cfg.DbUser[j], cfg.Dbpass[j], ""))
	if conn.err == nil {
		conn.dbName = cfg.DbName[j]
	}

	return
}

// GetContext get context
func (conn *Neo4jConn) GetContext() context.Context {
	return conn.ctx
}

func (conn *Neo4jConn) sessionEnd() {
	conn.trx = nil
	if conn.session != nil {
		err := conn.session.Close(conn.ctx)
		if err != nil {
			log.Printf("[SSSN] %v\n", err)
		} else {
			conn.session = nil
		}
	}
}

// Disconnect disconnect from the database
func (conn *Neo4jConn) Disconnect() {
	conn.sessionEnd()
	err := conn.driver.Close(conn.ctx)
	if err != nil {
		log.Printf("[DRVR] %v\n", err)
	}
}

// Connected return true if connected to the database
func (conn *Neo4jConn) Connected() bool {
	if conn.err != nil {
		log.Printf("[DRVR] %v\n", conn.err)
	}
	err := conn.driver.VerifyConnectivity(conn.ctx)
	if err != nil {
		log.Printf("[CONN] %v\n", err)
		return false
	}
	return true
}

// TrxReady return true if transaction is ready
func (conn *Neo4jConn) TrxReady() bool {
	return conn.trx != nil
}

// BeginTrx begin a transaction
func (conn *Neo4jConn) BeginTrx() (err error) {
	conn.session = conn.driver.NewSession(conn.ctx, neo4j.SessionConfig{DatabaseName: conn.dbName})
	conn.trx, err = conn.session.BeginTransaction(conn.ctx)
	return
}

// Commit commit a transaction
func (conn *Neo4jConn) Commit() error {
	defer conn.sessionEnd()
	return conn.trx.Commit(conn.ctx)
}

// Rollback rollback a transaction
func (conn *Neo4jConn) Rollback() error {
	defer conn.sessionEnd()
	return conn.trx.Rollback(conn.ctx)
}

// CYPHER_ADD
const CYPHER_ADD = "MATCH (p:N) WHERE elementId(p)=$elmId CREATE (c:N {r:1, d:$data, v:0})<-[:C]-(p) " +
	"WITH c MATCH (c)<-[:C*]-(b) WITH c, b, " +
	"CASE WHEN apoc.bitwise.op(b.d, \"&\", 7) = $winner THEN 1 WHEN b.w IS NULL THEN null ELSE 0 END AS w, " +
	"CASE WHEN apoc.bitwise.op(c.d, \"&\", 7) = $winner THEN 1 ELSE null END AS v " +
	"SET b.r = b.r + 1, b.w = coalesce(b.w, 0) + w, c.w = v RETURN elementId(c)"

// AddNode add a new leaf node to the tree
func (conn *Neo4jConn) AddNode(
	elmId nodes.Nid, // elementId of the node which the leaf node is to be added to
	leaf *nodes.Node, // leaf node to be added to the tree
	winner uint8, // winner of the current simulation
) (
	nid []map[string]any, // TEMP
	err error,
) {
	if conn.TrxReady() {
		var result neo4j.ResultWithContext
		result, err = conn.trx.Run(
			conn.ctx,
			CYPHER_ADD,
			map[string]any{
				"elmId":  elmId.(string),
				"data":   leaf.D,
				"winner": winner,
			},
		)
		if err != nil {
			return
		}
		var records []*db.Record
		records, err = result.Collect(conn.ctx)
		if err != nil {
			return
		}
		for _, rcrd := range records {
			nid = append(nid, rcrd.AsMap())
		}
	}
	return
}

// //////////////////////////////////////////
// ExecuteQuery TEMP function for testing
func (conn *Neo4jConn) ExecuteQuery(query string, params map[string]any) (out []map[string]any, err error) {
	if conn.Connected() {
		var result *neo4j.EagerResult
		result, err = neo4j.ExecuteQuery(conn.ctx, conn.driver, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(conn.dbName))
		if err != nil {
			return
		}
		for _, rcrd := range result.Records {
			out = append(out, rcrd.AsMap())
		}
	}
	return
}

// ExecuteTrx TEMP function for testing
func (conn *Neo4jConn) ExecuteTrx(query string, params map[string]any) (out []map[string]any, err error) {
	if conn.TrxReady() {
		var result neo4j.ResultWithContext
		result, err = conn.trx.Run(conn.ctx, query, params)
		if err != nil {
			return
		}
		var records []*db.Record
		records, err = result.Collect(conn.ctx)
		if err != nil {
			return
		}
		for _, rcrd := range records {
			out = append(out, rcrd.AsMap())
		}
	}
	return
}

func (conn *Neo4jConn) TestId(elmId nodes.Nid) string {
	return elmId.(string)
}
