package models

type Link struct {
	Label string
	Id    *string
	Route string
}

func NewLink(label, route string, id *string) Link {
	return Link{
		Label: label,
		Route: route,
		Id:    id,
	}
}
