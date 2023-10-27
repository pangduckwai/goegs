package storage

import (
	"context"

	"sea9.org/go/neo4j/config"
	"sea9.org/go/neo4j/nodes"
	"sea9.org/go/neo4j/storage/impl"
)

type Conn interface {
	GetContext() context.Context
	Disconnect()
	Connected() bool
	BeginTrx() error
	Commit() error
	Rollback() error
	TrxReady() bool

	ReadTree(
		variant string, // game variant string
	) (
		*nodes.Tree, // tree read from db
		error,
	)

	ReadNext(
		curr nodes.Nid, // the node which child nodes are to be read
	) (
		[]*nodes.Node, // list of sorted child nodes
		[]nodes.Nid, // list of node IDs of the sorted child nodes
		error,
	)

	AddNode(
		parent nodes.Nid, // the node which the new node is to be added to
		node *nodes.Node, // new node to be added to the tree
		winner uint8, // winner of the current simulation
	) (
		nodes.Nid, // the newly added leaf node
		error,
	)

	TestId(elmId nodes.Nid) string
	ExecuteQuery(query string, params map[string]any) ([]map[string]any, error)
	ExecuteTrx(query string, params map[string]any) ([]map[string]any, error)
}

func Connect(cfg *config.Config, ctx context.Context) (conn Conn, err error) {
	return impl.Connect(cfg, ctx)
}
