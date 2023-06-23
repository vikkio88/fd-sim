package models

import (
	"fdsim/conf"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

const newsInMemoryId = "newsId"

func newsIdGenerator() string {
	return fmt.Sprintf("%s_%s", newsInMemoryId, ulid.Make())
}

type News struct {
	Id        string
	Date      time.Time
	Title     string
	NewsPaper string //TODO: Newspapers to country, also country to game
	Body      string
	Links     []Link
	Read      bool
}

func NewNews(title, newsPaper, body string, date time.Time, links []Link) *News {
	return &News{
		Id:        newsIdGenerator(),
		Date:      date,
		Title:     title,
		NewsPaper: newsPaper,
		Body:      body,
		Links:     links,
	}
}

func NewNewsWithId(id, title, newsPaper, body string, date time.Time, links []Link) *News {
	return &News{
		Id:        id,
		Date:      date,
		Title:     title,
		NewsPaper: newsPaper,
		Body:      body,
		Links:     links,
	}
}

func (n News) String() string {
	return fmt.Sprintf("%s - %s", n.Title, n.Date.Format(conf.DateFormatShort))
}
