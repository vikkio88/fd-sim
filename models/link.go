package models

type Link struct {
	Label       string
	Route       string
	Id          *string
	SubtabIndex *int
}

func NewLink(label, route string, id *string) Link {
	return Link{
		Label: label,
		Route: route,
		Id:    id,
	}
}

func NewLinkSubTab(label, route string, id *string, subTabIndex *int) Link {
	return Link{
		Label:       label,
		Route:       route,
		Id:          id,
		SubtabIndex: subTabIndex,
	}
}
