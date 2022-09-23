package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Coin struct {
	gorm.Model
	Name 		string 	`json:"name" validate:"nonzero"`
	ShortName	string	`json:"shortname" validate:"len=3"`
	Votes 		int    	`json:"votes" validate:"min=0"`
	Price 		float64	`json:"price" validate:"min=0"`
}

func ValidateCoinData(coin *Coin) error {
	if err := validator.Validate(coin); err != nil {
		return err
	}
	return nil
}