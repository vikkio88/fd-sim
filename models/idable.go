package models

type Idable struct {
	Id string
}

func NewIdable(id string) Idable {
	return Idable{id}
}
