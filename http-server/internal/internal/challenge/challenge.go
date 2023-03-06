package challenge

import (
	"encoding/hex"
	"errors"

	"github.com/google/uuid"
)

const (
	MAX_RETRY_COUNT   = 3
	SUFFIX_BYTES_SIZE = 2

	errUnableGenerateTask = "errUnableGenerateTask"
	errUnknownTaskID      = "errUnknownTaskID"
)

func NewChallenge() *Service {
	return &Service{
		g: &randomGenerator{SUFFIX_BYTES_SIZE},
		v: &sha256Verifier{SUFFIX_BYTES_SIZE},

		tasks: make(map[uuid.UUID][]byte),
	}
}

type Service struct {
	g generator
	v verifier

	tasks map[uuid.UUID][]byte
}

func (c *Service) TryGenerate() (string, string, error) {
	retry := 0
	for retry < MAX_RETRY_COUNT {
		id, suffix, err := c.g.generate()
		if err != nil {
			retry++
			continue
		}

		c.tasks[id] = suffix

		return id.String(), hex.EncodeToString(suffix), nil
	}

	return "", "", errors.New(errUnableGenerateTask)
}

func (c *Service) TryVerify(id string, token string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return errors.New(errUnknownTaskID)
	}

	expected, ok := c.tasks[uuid]
	if !ok {
		return errors.New(errUnknownTaskID)
	}

	if !c.v.verify(expected, token) {
		return errors.New(errUnknownTaskID)
	}

	return nil
}
