package model

import ( 
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bank select {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name" valid:"notnull"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.Id = uuid.NewV4().String()
	bank.created_at = time.Now().UTC()

	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil
}