package impl

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"sea9.org/go/neo4j/common"
	"sea9.org/go/neo4j/config"
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
