package model

import "github.com/google/uuid"

type Unit struct {
	Id   string
	Name string
}

func NewUnit(n string) *Unit {
	return &Unit{Id: uuid.NewString(), Name: n}
}
