package sqlgcr

import (
	"encoding/json"
	"fmt"
)

type Cassette struct {
	Queries []*Query
	current int
}

func UnmarshalCassette(content []byte) (Cassette, error) {
	cas := new(Cassette)
	if err := json.Unmarshal(content, cas); err != nil {
		return Cassette{}, fmt.Errorf("can't unmarshal cassette: %v", err)
	}
	return *cas, nil
}

func (c Cassette) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func (c *Cassette) CurrentQuery() *Query {
	if c.current + 1 > len(c.Queries) {
		return nil
	}
	q := c.Queries[c.current]
	c.current++
	return q
}
