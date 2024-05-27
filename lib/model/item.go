package model

import "github.com/google/uuid"

type item struct {
	Id    string
	Name  string
	Price float32
}

func NewItem(n string, p float32) *item {
	i := new(item)
	i.Id = uuid.NewString()
	i.Name = n
	i.Price = p
	return i
}
