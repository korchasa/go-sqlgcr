package sqlgcr

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
)

type Stmt struct {
	t *testing.T
	q *Query
}

func NewStmt(t *testing.T, q *Query) *Stmt {
	return &Stmt{t: t, q: q}
}

func (s *Stmt) Close() error {
	return nil
}

func (s *Stmt) NumInput() int {
	return -1
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	fmt.Printf("\n\nExec: %v\n\n", args)
	return &s.q.Result, nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	s.q.Args = args
	js, err := json.MarshalIndent(s.q, "", "  ")
	if err != nil {
		s.t.Errorf("Internal sqlgcr error: can't marshal query: %v", err)
	}
	if s.q.IsUnexpected {
		s.t.Errorf("Unexpected query in %s: %s", s.q.CassettePath, string(js))
	}
	return &s.q.Rows, nil
}
