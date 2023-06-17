package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/anukul/xmon/backend/common"
	"github.com/anukul/xmon/backend/config"
)

type Database struct {
	level *leveldb.DB
	scope string
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	db, err := leveldb.OpenFile(cfg.DatabasePath, nil)
	if err != nil {
		return nil, err
	}
	return &Database{level: db}, nil
}

func (d *Database) WithScope(scope string) common.DatabaseClient {
	return &Database{level: d.level, scope: scope}
}

func (d *Database) getScopedKey(key string) string {
	if d.scope != "" {
		return fmt.Sprintf("%s_%s", d.scope, key)
	} else {
		return key
	}
}

func (d *Database) GetScope() string {
	return d.scope
}
