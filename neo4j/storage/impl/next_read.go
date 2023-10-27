package impl

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"sea9.org/go/neo4j/nodes"
)

// CYPHER_READ_NEXT cypher statement to read all child nodes of a given node
const CYPHER_READ_NEXT = "MATCH (p:N)-[:C]->(c) WHERE elementId(p)=$elmId RETURN elementId(c) as id, c{.*} as prop ORDER BY c.d, coalesce(c.u, 0)"

// ReadNext read all child nodes of the given node
func (conn *Neo4jConn) ReadNext(curr nodes.Nid) (next []*nodes.Node, nids []nodes.Nid, err error) {
	if conn.Connected() {
		var result *neo4j.EagerResult
		result, err = neo4j.ExecuteQuery(
			conn.ctx, conn.driver,
			CYPHER_READ_NEXT,
			map[string]any{
				"elmId": curr.(string),
			},
			neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(conn.dbName),
		)
		if err != nil {
			return
		}

		for _, record := range result.Records {
			for k, v := range record.AsMap() {
				switch k {
				case "id":
					nids = append(nids, v)
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
					next = append(next, nxt)
				}
			}
		}
	}
	return
}
