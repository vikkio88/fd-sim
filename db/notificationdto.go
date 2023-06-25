package db

import (
	"encoding/json"
	"fdsim/models"
	"fmt"
	"time"
)

func unserialiseLinks(links *string) []models.Link {
	if links == nil {
		return []models.Link{}
	}

	var result []models.Link
	data := *links
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading link", err)
		return []models.Link{}
	}

	return result
}

func serialiseLinks(links []models.Link) *string {
	var result string
	data, _ := json.Marshal(links)
	result = string(data)

	return &result
}

func unserialiseActions(actions *string) []models.Action {
	return []models.Action{}
}
func serialiseActions(actions []models.Action) *string {
	return nil
}

type NewsDto struct {
	Id        string `gorm:"primarykey"`
	Date      time.Time
	Title     string
	NewsPaper string
	Body      string
	Links     *string
	Read      bool
}

func DtoFromNews(news *models.News) NewsDto {
	return NewsDto{
		Id:        news.Id,
		Date:      news.Date,
		Title:     news.Title,
		NewsPaper: news.NewsPaper,
		Body:      news.Body,
		Links:     serialiseLinks(news.Links),
		Read:      news.Read,
	}
}

func (news *NewsDto) News() *models.News {
	return &models.News{
		Id:        news.Id,
		Date:      news.Date,
		Title:     news.Title,
		NewsPaper: news.NewsPaper,
		Body:      news.Body,
		Links:     unserialiseLinks(news.Links),
		Read:      news.Read,
	}
}

type EmailDto struct {
	Id      string `gorm:"primarykey"`
	Read    bool
	Sender  string
	Subject string
	Body    string
	Date    time.Time
	Links   *string
	Actions *string
}

func DtoFromEmail(email *models.Email) EmailDto {
	return EmailDto{
		Id:      email.Id,
		Read:    email.Read,
		Sender:  email.Sender,
		Subject: email.Subject,
		Body:    email.Body,
		Date:    email.Date,
		Links:   serialiseLinks(email.Links),
		Actions: serialiseActions(email.Actions),
	}
}

func (email *EmailDto) Email() *models.Email {
	return &models.Email{
		Id:      email.Id,
		Read:    email.Read,
		Sender:  email.Sender,
		Subject: email.Subject,
		Body:    email.Body,
		Date:    email.Date,
		Links:   unserialiseLinks(email.Links),
		Actions: unserialiseActions(email.Actions),
	}
}
