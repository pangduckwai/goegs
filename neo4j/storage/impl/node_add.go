package impl

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"sea9.org/go/neo4j/common"
	"sea9.org/go/neo4j/nodes"
)

// CYPHER_ADD_NODE cypher statement to add a node and perform back-propagate
const CYPHER_ADD_NODE = "MATCH (p:N) WHERE elementId(p)=$elmId AND NOT EXISTS {MATCH (:N {r:1, d:%v%v})<-[:C]-(p)} " +
	"CREATE (c:N {r:1, d:%v%v})<-[:C]-(p) " +
	"WITH c MATCH (c)<-[:C*]-(b) WITH c, b, " +
	"CASE WHEN apoc.bitwise.op(b.d, \"&\", 7) = $winner THEN 1 WHEN b.w IS NULL THEN null ELSE 0 END AS w, " +
	"CASE WHEN apoc.bitwise.op(c.d, \"&\", 7) = $winner THEN 1 ELSE null END AS v " +
	"SET b.r = b.r + 1, b.w = coalesce(b.w, 0) + w, c.w = v RETURN elementId(c)"

// AddNode add a new node to the tree
func (conn *Neo4jConn) AddNode(parent nodes.Nid, node *nodes.Node, winner uint8) (nid nodes.Nid, err error) {
	if conn.TrxReady() {
		var result neo4j.ResultWithContext
		var strv string
		if node.V > 0 {
			strv = fmt.Sprintf(", u:%v", node.V)
		}
		result, err = conn.trx.Run(
			conn.ctx,
			fmt.Sprintf(CYPHER_ADD_NODE, node.D, strv, node.D, strv),
			map[string]any{
				"elmId":  parent.(string),
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

		rslt := map[string]struct{}{}
		for _, rcrd := range records {
			for _, v := range rcrd.AsMap() {
				rslt[v.(string)] = struct{}{}
			}
		}

		var nids []string
		for k := range rslt {
			nids = append(nids, k)
		}
		switch len(nids) {
		case 0:
		case 1:
			nid = nids[0]
		default:
			err = common.NewError(true, "DATA", fmt.Sprintf("Invalid number of new node ID(s) returned: %v", nids))
		}
	}
	return
}
