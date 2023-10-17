package model

import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	Id 				string 		`json:"id valid:"uuid"`
	CreatedAt time.Time `json:"created" valid:"-"`
	UpdatedAt time.Time `json:"updated" valid:"-"`
}