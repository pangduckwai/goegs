package impl

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"sea9.org/go/neo4j/common"
	"sea9.org/go/neo4j/nodes"
)

// CYPHER_READ_TREE cypher statement to read the tree object
const CYPHER_READ_TREE = "MERGE (t:T {v: $variant}) ON CREATE SET t:N, t.r=1, t.d=0 RETURN elementId(t) as id, t{.*} as prop"

// ReadTree read the tree object
func (conn *Neo4jConn) ReadTree(variant string) (tree *nodes.Tree, err error) {
	if conn.Connected() {
		var result *neo4j.EagerResult
		result, err = neo4j.ExecuteQuery(
			conn.ctx, conn.driver,
			CYPHER_READ_TREE,
			map[string]any{
				"variant": variant,
			},
			neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(conn.dbName),
		)
		if err != nil {
			return
		}

		if len(result.Records) != 1 {
			err = common.NewError(true, "DATA", "Invalid number of tree(s) found")
			return
		}

		tree = &nodes.Tree{
			V: variant,
			R: &nodes.Node{},
		}
		for k, v := range result.Records[0].AsMap() {
			switch k {
			case "id":
				tree.R.ID = v
			case "prop":
				for j, u := range v.(map[string]any) {
					switch j {
					case "d":
						tree.R.D = uint32(u.(int64))
					case "r":
						tree.R.R = uint64(u.(int64))
					case "w":
						tree.R.W = uint64(u.(int64))
					}
				}
			}
		}
	}
	return
}
