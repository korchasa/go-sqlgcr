package sqlgcr

import (
	"database/sql/driver"
	"io"
)

type Rows struct {
	Cols []string
	Data [][]driver.Value
	pos int
}

func (r *Rows) Columns() []string {
	return r.Cols
}

func (r *Rows) Close() error {
	return nil
}

func (r *Rows) Next(dest []driver.Value) error {
	if r.pos + 1 > len(r.Data) {
		return io.EOF
	}
	res := r.Data[r.pos]
	r.pos++
	if r.pos > len(r.Data) {
		return io.EOF
	}

	for i, col := range res {
		dest[i] = col
	}

	return nil
}

