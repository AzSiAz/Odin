// Package database contains all action related to database access
package database

type Database interface {
	Insert() error
	Update() error
	Upsert() error
	Delete() error
	List() error
	Get() error
}

type DB struct{}

func New(addr string) *DB {
	return &DB{}
}

func (d *DB) Insert() error {
	return nil
}

func (d *DB) Update() error {
	return nil
}

func (d *DB) Upsert() error {
	return nil
}

func (d *DB) Delete() error {
	return nil
}

func (d *DB) List() error {
	return nil
}

func (d *DB) Get() error {
	return nil
}
