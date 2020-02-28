package sqlgcr

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Player struct {
	path string
	t    *testing.T
	Cas  *Cassette
}

func NewPlayer(t *testing.T) *Player {
	return &Player{t:t}
}

func (p *Player) Load(path string) error {
	p.path = path
	info, err := os.Stat(p.path)
	var cas Cassette
	if !os.IsNotExist(err) {
		if info.IsDir() {
			return fmt.Errorf("cassette file `%s` is a directory", p.path)
		}
		js, err := ioutil.ReadFile(p.path)
		if err != nil {
			return fmt.Errorf("can't read cassette file `%s`: %v", p.path, err)
		}
		cas, err = UnmarshalCassette(js)
		if err != nil {
			return fmt.Errorf("can't unmarshal cassette file `%s`: %v", p.path, err)
		}
	}
	p.Cas = &cas
	return nil
}