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

	AddNode(
		elmId nodes.Nid, // elementId of the node which the leaf node is to be added to
		leaf *nodes.Node, // leaf node to be added to the tree
		winner uint8, // winner of the current simulation
	) (
		[]map[string]any, // TEMP
		error,
	)

	TestId(elmId nodes.Nid) string
	ExecuteQuery(query string, params map[string]any) ([]map[string]any, error)
	ExecuteTrx(query string, params map[string]any) ([]map[string]any, error)
}

func Connect(cfg *config.Config, ctx context.Context) (conn Conn, err error) {
	return impl.Connect(cfg, ctx)
}
