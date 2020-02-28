package sqlgcr

import (
	"database/sql"
	"testing"
)

func init() {
	sql.Register("sqlgcr", new(Driver))
}

var currentPlayer *Player

func New(t *testing.T, path string) (*sql.DB, error) {
	db, err := sql.Open("sqlgcr", path)
	if err != nil {
		return nil, err
	}
	currentPlayer = NewPlayer(t)
	if err := currentPlayer.Load(path); err != nil {
		return nil, err
	}
	return db, db.Ping()
}
