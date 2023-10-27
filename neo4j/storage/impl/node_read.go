package impl

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"sea9.org/go/neo4j/common"
	"sea9.org/go/neo4j/nodes"
)

// CYPHER_READ_NEXT cypher statement to read all child nodes of a given node
const CYPHER_READ_NODE = "MATCH (n:N) WHERE elementId(n)=$elmId RETURN n{.*} as prop"

// ReadNext read all child nodes of the given node
func (conn *Neo4jConn) ReadNode(curr nodes.Nid) (node *nodes.Node, err error) {
	if conn.Connected() {
		var result *neo4j.EagerResult
		result, err = neo4j.ExecuteQuery(
			conn.ctx, conn.driver,
			CYPHER_READ_NODE,
			map[string]any{
				"elmId": curr.(string),
			},
			neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(conn.dbName),
		)
		if err != nil {
			return
		}

		if len(result.Records) != 1 {
			err = common.NewError(true, "DATA", "Invalid number of node(s) found")
			return
		}

		node = &nodes.Node{}
		for k, v := range result.Records[0].AsMap() {
			switch k {
			case "prop":
				nxt := &nodes.Node{}
				for j, u := range v.(map[string]any) {
					switch j {
					case "d":
						nxt.D = uint32(u.(int64))
					case "u":
						nxt.V = uint16(u.(int64))
					case "r":
						nxt.R = uint64(u.(int64))
					case "w":
						nxt.W = uint64(u.(int64))
					}
				}
			}
		}
	}
	return
}
