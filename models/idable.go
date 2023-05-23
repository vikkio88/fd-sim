package models

type idable struct {
	Id string
}

func NewIdable(id string) idable {
	return idable{id}
}
