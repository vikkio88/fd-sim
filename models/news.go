package models

import (
	"fmt"
	"time"
)

type News struct {
	Date      time.Time
	Title     string
	NewsPaper string //TODO: Newspapers to country, also country to game
	Body      string
	Links     []Link
	Read      bool
}

func NewNews(date time.Time, title, newsPaper, body string, links []Link) *News {
	return &News{
		Date:      date,
		Title:     title,
		NewsPaper: newsPaper,
		Body:      body,
		Links:     links,
	}
}

func (n News) String() string {
	return fmt.Sprintf("%s - %s", n.Title, n.NewsPaper)
}
