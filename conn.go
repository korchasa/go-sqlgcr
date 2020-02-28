package sqlgcr

import (
	"database/sql/driver"
)

type Conn struct {
	player *Player
}

func NewConn(player *Player) *Conn {
	return &Conn{player: player}
}

func (c *Conn) Prepare(queryText string) (driver.Stmt, error) {
	query := c.player.Cas.CurrentQuery()
	if query == nil {
		query = &Query{IsUnexpected:true, CassettePath:c.player.path, Text:queryText}
	}
	return query.ToStatement(c.player.t), nil
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return nil, nil
}
