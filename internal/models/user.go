package models

import (
	"encoding/json"
	"time"

	"github.com/bagasunix/go-clean-architecture/pkg/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

type CreateUser struct {
	FullName string `json:"full_name"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *CreateUser) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

func (c CreateUser) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FullName, validation.Required.Error("Nama harus diisi")),
		validation.Field(&c.Sex, validation.Required.Error("Jenis kelamin harus diisi")),
		validation.Field(&c.Email, is.Email),
		validation.Field(&c.Password, validation.Required.Error("Password harus diisi")),
	)
}

type UpdateUser struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Sex      string    `json:"sex"`
	Password string    `json:"password"`
}

func (s *UpdateUser) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

func (c UpdateUser) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FullName, validation.Required.Error("Nama harus diisi")),
		validation.Field(&c.Sex, validation.Required.Error("Jenis kelamin harus diisi")),
	)
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *LoginUser) ToJSON() []byte {
	j, err := json.Marshal(s)
	errors.HandlerReturnedVoid(err)
	return j
}

func (c LoginUser) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, is.Email.Error("Format email salah"), validation.Required.Error("Email harus diisi")),
		validation.Field(&c.Password, validation.Required.Error("Password harus diisi")),
	)
}

type ResponseUser struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
