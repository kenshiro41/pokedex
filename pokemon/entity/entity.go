package entity

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Pokemon struct {
	ID   int            `json:"id"`
	Name postgres.Jsonb `json:"name"`
	Type pq.StringArray `gorm:"type:varchar(255)[]" json:"type"`
	Base postgres.Jsonb `json:"base"`
}
type Name struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}
type Base struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}
