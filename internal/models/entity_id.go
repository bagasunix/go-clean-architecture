package models

import (
	"encoding/json"

	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type EntityId struct {
	Id any `json:"id"`
}

func (c EntityId) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Id, validation.Required.Error("id harus diisi")),
	)
}

func (s *EntityId) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}
