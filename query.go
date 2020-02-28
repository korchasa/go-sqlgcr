package sqlgcr

import (
	"database/sql/driver"
	"testing"
)

type Query struct {
	IsUnexpected bool `json:"-"`
	CassettePath string `json:"-"`
	Text         string
	Args         []driver.Value
	Rows         Rows
	Result       Result
}

func (q *Query) ToStatement(t *testing.T) driver.Stmt {
	return NewStmt(t, q)
}
