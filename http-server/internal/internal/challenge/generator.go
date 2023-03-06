package challenge

import (
	"crypto/rand"

	"github.com/google/uuid"
)

type generator interface {
	generate() (uuid.UUID, []byte, error)
}

type randomGenerator struct {
	bytesLen int
}

func (g *randomGenerator) generate() (id uuid.UUID, suffix []byte, err error) {
	id = uuid.New()
	suffix = make([]byte, g.bytesLen)
	_, err = rand.Read(suffix)

	return id, suffix, err
}
