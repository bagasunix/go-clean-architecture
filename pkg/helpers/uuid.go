package helpers

import (
	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

func GenerateUUIDV1(logs *zap.Logger) uuid.UUID {
	id, err := uuid.NewV1()
	errors.HandlerWithLoggerReturnedVoid(logs, err, "uuid", "generator")
	return id
}

func GenerateUUIDV4(logs *zap.Logger) uuid.UUID {
	id, err := uuid.NewV4()
	errors.HandlerWithLoggerReturnedVoid(logs, err, "uuid", "generator")
	return id
}
