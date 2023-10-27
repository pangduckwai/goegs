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

// ReadRow read a row of node from the database.
func (conn *PsqlConn) ReadRow(id nodes.Nid) (
	exists bool,
	data uint32,
	value uint16,
	runs, wins uint64,
	err error,
) {
	if conn.Connected() {
		var da, va, ra, wa any
		err = conn.pool.QueryRow(conn.ctx, SQL_NODES_SEL, id.ToAny()...).Scan(&da, &va, &ra, &wa)
		if err != nil {
			if err == pgx.ErrNoRows {
				err = nil
			}
			return
		}

		var d int32
		var v int16
		var r, w int64
		var ok bool
		d, ok = FromAny32(da)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.d", da))
			return
		}
		data = uint32(d)

		v, ok = FromAny16(va)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.v", va))
			return
		}
		value = uint16(v)

		r, ok = FromAny64(ra)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.r", ra))
			return
		}
		runs = uint64(r)

		w, ok = FromAny64(wa)
		if !ok {
			err = common.NewError(true, ERR_DATA, fmt.Sprintf("Invalid value %v of 'nodes.w", wa))
			return
		}
		wins = uint64(w)

		exists = true
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

// ReadTree read a tree and its root node from the database.
func (conn *PsqlConn) ReadTree(variant string) (
	nid nodes.Nid, // hash (a.k.a. pkey) of the root node
	tree *nodes.Tree, // the tree read from the database
	err error,
) {
	if conn.Connected() {
		var v string
		var t time.Time
		err = conn.pool.QueryRow(conn.GetContext(), SQL_TREES_SEL, variant).Scan(&v, &t)
		if err != nil {
			if err == pgx.ErrNoRows {
				err = nil
			}
			return
		}

		var r *nodes.Node
		tree = &nodes.Tree{
			V: v,
		}
		// hash = nodeutils.CalcHash(tree.V)
		nid = nodeId.Init(tree.V)
		r, err = conn.ReadNode(nid)
		if err != nil || r == nil {
			return
		}

		tree = &nodes.Tree{
			V: variant,
			R: r,
		}
		r.Parent = tree
	}
	return
}

// ReadChild search and read all child nodes of the given nodo
func (conn *PsqlConn) ReadChild(vrn string, pth []uint8) (
	next []*nodes.Node, // child nodes read from the database
	path [][]uint8, // path of each child node read
	err error,
) {
	if conn.Connected() {
		var n *nodes.Node
		for i := uint8(0); i <= 255; i++ {
			q := append(make([]uint8, 0), pth...)
			q = append(q, i)
			n, err = conn.ReadNode(nodeId.Init(vrn, q...))
			if err != nil {
				return
			}
			if n == nil {
				break
			}
			next = append(next, n)
			path = append(path, q)
		}
	}
	return
}

// FillNext read child nodes and merge (and sort) to Next[] of the given node
func (conn *PsqlConn) FillNext(vrn string, pth []uint8, curr *nodes.Node) (
	count int, // the number of child nodes read from the database
	err error,
) {
	var next []*nodes.Node
	var found bool

	next, _, err = conn.ReadChild(vrn, pth)
	if err != nil {
		return
	}
	for _, c := range next {
		found, _, err = curr.AddChild(c)
		if !found {
			count++
		}
	}
	return
}

// GetNext read child nodes of the given node while keeping the original order
func (conn *PsqlConn) GetNext(vrn string, pth []uint8, curr *nodes.Node) (
	count int, // the number of child nodes read from the database
	path [][]uint8, // path of each child node read
	err error,
) {
	if len(curr.Next) > 0 {
		err = common.NewError(true, ERR_DATA, fmt.Sprintf("Children already exists in %v", curr))
		return
	}

	var next []*nodes.Node

	next, path, err = conn.ReadChild(vrn, pth)
	if err != nil {
		return
	}
	curr.Next = append(curr.Next, next...)
	for _, c := range next {
		c.Parent = curr
	}
	count = len(next)
	return
}

// WriteNode write a single node in a larger transaction
func (conn *PsqlConn) WriteNode(
	id nodes.Nid,
	n *nodes.Node,
) (total, updated int, err error) {
	if conn.TrxReady() {
		n.ClearTransient()
		var d, v, r, w any
		if n.D != 0 { // since zero values are stored as NULL in database
			d = int32(n.D)
		}
		if n.V != 0 {
			v = int32(n.V)
		}
		if n.R != 0 { // since zero values are stored as NULL in database
			r = int64(n.R)
		}
		if n.W != 0 { // since zero values are stored as NULL in database
			w = int64(n.W)
		}

		args := append(make([]any, 0), id.ToAny()...)
		args = append(args, d, v, r, w)
		err = conn.trx.QueryRow(conn.GetContext(), SQL_NODES_UPS, args...).Scan(&total, &updated)
	}
	return
}

// WriteNodes write all nodes in the tree to the database
func (conn *PsqlConn) WriteNodes(tree *nodes.Tree) (icnt, ucnt int, err error) {
	if conn.TrxReady() {
		var pc []uint8
		var nc *nodes.Node
		var ttl, upd, vldt int
		var sum, xpt uint64
		ps := [][]uint8{nil}
		ns := []*nodes.Node{tree.R}

		for len(ns) > 0 {
			pc = ps[0]
			nc = ns[0]

			// Validate current node  TODO HERE!!! comment this for debugging
			vldt, sum, xpt = nc.Validate()
			if vldt < 0 {
				err = common.NewError(true, ERR_DATA, fmt.Sprintf("Validation failed (%v S:%v X:%v) writing node\n%v\n%v", vldt, sum, xpt, common.ShowPath(pc), nc.Tree(len(pc)+1, 99)))
				return
			}

			// Process current node
			ttl, upd, err = conn.WriteNode(nodeId.Init(tree.V, pc...), nc)
			if err != nil {
				err = common.NewError(true, ERR_DATA, fmt.Sprintf("%v\n%v\n%v", err, pc, nc))
				return
			} else if ttl != 1 {
				err = common.NewError(true, ERR_DATA, fmt.Sprintf("%v row inserted/updated, expecting 1\n%v\n%v", ttl, pc, nc))
			}
			if upd > 0 {
				ucnt++
			} else {
				icnt++
			}

			// Traverse tree
			ns = append(nc.Next, ns[1:]...)

			var px [][]uint8
			for i := range nc.Next {
				p := append(make([]uint8, 0), pc...)
				p = append(p, uint8(i))
				px = append(px, p)
			}
			ps = append(px, ps[1:]...)
		}

		_, err = conn.trx.Exec(conn.GetContext(), SQL_TREES_UPS, tree.V)
	}
	return
}

///////////////////////
// Utility functions
///////////////////////

// FromAny16 convert values of type 'any' to int64
func FromAny16(inp any) (out int16, okay bool) {
	okay = true
	if inp != nil {
		out, okay = inp.(int16)
		if !okay {
			out = 0
		}
	}
	return
}

// FromAny32 convert values of type 'any' to int64
func FromAny32(inp any) (out int32, okay bool) {
	okay = true
	if inp != nil {
		out, okay = inp.(int32)
		if !okay {
			out = 0
		}
	}
	return
}

// FromAny convert values of type 'any' to int64
func FromAny64(inp any) (out int64, okay bool) {
	okay = true
	if inp != nil {
		out, okay = inp.(int64)
		if !okay {
			out = 0
		}
	}
	return
}
*/
