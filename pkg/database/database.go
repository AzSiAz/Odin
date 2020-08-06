// Package database contains all action related to database access
package database

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/v3/codec/msgpack"
)

type StormDB struct {
	DB *storm.DB
}

func New(addr string) (*StormDB, error) {
	db, err := storm.Open(addr, storm.Codec(msgpack.Codec))
	if err != nil {
		return nil, err
	}

	return &StormDB{
		DB: db,
	}, nil
}
