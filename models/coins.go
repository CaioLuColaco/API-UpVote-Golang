package models

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	Name 		string 	`json:"name"`
	ShortName	string	`json:"shortname"`
	Votes 		int    	`json:"votes"`
	Price 		float64	`json:"price"`
}

var Coins []Coin