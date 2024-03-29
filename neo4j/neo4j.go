package main

import (
	"fmt"
	"sort"

	"sea9.org/go/neo4j/config"
	"sea9.org/go/neo4j/nodes"
	"sea9.org/go/neo4j/storage"
)

func main() {
	cfg := &config.Config{
		[]string{"localhost"},
		[]string{"7687"},
		[]string{"neo4j"}, // DB name
		[]string{"neo4j"}, // User name
		[]string{"12345678"},
	}

	conn, err := storage.Connect(cfg, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Disconnect()

	if conn.Connected() {
		fmt.Println("Connected")
	} else {
		panic("Connection failed")
	}

	// var rslt []map[string]any

	// Test ID
	// fmt.Println(conn.TestId("helloThereHowAreYou???"))

	// Query
	// rslt, err = conn.ExecuteQuery("MATCH (t:T)-[r:C*]-(c) RETURN t,r,c", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// Write
	// if err := conn.BeginTrx(); err == nil {
	// 	rslt, err = conn.ExecuteTrx("MATCH (p:N) WHERE ID(p)=$nid CREATE (c:N {r:1})<-[:C]-(p) WITH c MATCH (c)<-[:C*]-(b) SET b.r = b.r + 1 RETURN c", map[string]any{"nid": 8})
	// 	if err != nil {
	// 		conn.Rollback()
	// 		panic(err)
	// 	}
	// 	conn.Commit()
	// 	fmt.Println(rslt)
	// } else {
	// 	panic(err)
	// }

	// Print results
	// for _, rst := range rslt {
	// 	for k, v := range rst {
	// 		if k == "r" {
	// 			fmt.Printf("%v :\n", k)
	// 			for _, m := range v.([]any) {
	// 				fmt.Printf("   -%+v\n", m)
	// 			}
	// 		} else {
	// 			fmt.Printf("%v : %+v\n", k, v)
	// 		}
	// 	}
	// }

	// Read tree
	tree, err := conn.ReadTree("TESTING-3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Root: %v (%v)\n", tree, tree.R.ID)
	fmt.Printf("%v\n\n", tree.R)

	// Add node
	var nid nodes.Nid
	if err := conn.BeginTrx(); err == nil {
		nid, err = conn.AddNode("4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:13", nodes.New(nil, 0, 2, []uint8{1}, 3), 0)
		if err != nil {
			conn.Rollback()
			panic(err)
		}
		conn.Commit()
		fmt.Println(nid)
		fmt.Println()
	} else {
		panic(err)
	}

	// Read child
	next, _, err := conn.ReadNext("4:96d0ac42-1e05-4930-9584-b243aa9c8b8a:13")
	if err != nil {
		panic(err)
	}
	for _, n := range next {
		fmt.Printf(" - %v\n", n)
	}

	fmt.Printf("CFG: %+v\n", cfg)

	if sort.IsSorted(nodes.Sorted(next)) {
		fmt.Println("Sorted")
	} else {
		fmt.Println("Not sorted")
	}
}

/*
 * Impl

// Count count number of nodes in the database.
func (conn *PsqlConn) Count() (count uint64, err error) {
	var ca any
	err = conn.pool.QueryRow(conn.ctx, SQL_NODES_CNT).Scan(&ca)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = nil
		}
		return
	}

	c, ok := FromAny64(ca)
	if !ok {
		err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of count(*) from nodes", ca))
		return
	}
	count = uint64(c)
	return
}

// List retrieve all rows of node from the database.
// Has to read all nodes at once, because the current design it is impossible to select nodes with any
// meaningful criteria
func (conn *PsqlConn) List(cr chan *storage.Nrow) (cnt int, err error) {
	if cr != nil {
		defer close(cr)
	}

	var ma, na, da, va, ra, wa any
	var rows pgx.Rows

	rows, err = conn.pool.Query(conn.ctx, SQL_NODES_ALL)
	if err != nil {
		return
	}

	var v int16
	var d int32
	var r, w int64
	var ok bool
	for rows.Next() {
		err = rows.Scan(&ma, &na, &da, &va, &ra, &wa)
		if err != nil {
			return
		}
		row := &storage.Nrow{}
		row.M, ok = FromAny64(ma)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.m", ma))
			return
		}
		row.N, ok = FromAny64(na)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.n", na))
			return
		}

		d, ok = FromAny32(da)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.d", da))
			return
		}
		row.D = uint32(d)

		v, ok = FromAny16(va)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.v", va))
			return
		}
		row.V = uint16(v)

		r, ok = FromAny64(ra)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.r", ra))
			return
		}
		row.R = uint64(r)

		w, ok = FromAny64(wa)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.w", wa))
			return
		}
		row.W = uint64(w)

		cnt++
		if cr != nil {
			cr <- row
		}
	}

	return
}

// ReadNode read a node by the node ID (the given hash) from the database.
// Also return the current node's ID
func (conn *PsqlConn) ReadNode(id nodes.Nid) (
	c *nodes.Node, // the node read from the database
	err error,
) {
	var d uint32
	var v uint16
	var r, w uint64 // since zero values are stored as NULL in database
	var ok bool

	ok, d, v, r, w, err = conn.ReadRow(id)
	if err != nil || !ok {
		return
	}

	c = &nodes.Node{
		D: d,
		V: v,
		R: r,
		W: w,
	}

	return
}
*/
