package db

import (
	"testing"

	"github.com/anukul/xmon/backend/config"

	"github.com/stretchr/testify/assert"
)

var db *Database

func setup() func() {
	var err error
	db, err = NewDatabase(&config.Config{DatabasePath: "./scratch/database"})
	if err != nil {
		panic(err)
	}
	return func() {}
}

func TestDatabase_Set(t *testing.T) {
	defer setup()()
	if err := db.Set("xxx", "yyy"); err != nil {
		panic(err)
	}
}

func TestDatabase_Get(t *testing.T) {
	defer setup()()
	var data string
	if err := db.Get("xxx", &data); err != nil {
		panic(err)
	}
	assert.Equal(t, "yyy", data)
}
