package storage

import (
	"context"

	"sea9.org/go/neo4j/config"
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

	ExecuteQuery(query string, params map[string]any) ([]string, error)
	ExecuteTrx(query string, params map[string]any) ([]string, error)
}

func Connect(cfg *config.Config, ctx context.Context) (conn Conn, err error) {
	return impl.Connect(cfg, ctx)
}
