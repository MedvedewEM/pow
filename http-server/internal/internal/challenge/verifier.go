package challenge

import (
	"crypto/sha256"

	"github.com/MedvedewEM/pow/pkg/slices"
)

type verifier interface {
	verify(expected []byte, actual string) bool
}

type sha256Verifier struct {
	bytesLen int
}

func (v *sha256Verifier) verify(expected []byte, actualStr string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(actualStr))
	actualBytes := hasher.Sum(nil)

	actual := actualBytes[len(actualBytes)-v.bytesLen:]

	return slices.Equal(actual, expected)
}
