package db

import (
	"fdsim/models"
	"time"
)

func unserialiseLinks(links *string) []models.Link {
	return []models.Link{}
}
func serialiseLinks(links []models.Link) *string {
	return nil
}

func unserialiseActions(actions *string) []models.Action {
	return []models.Action{}
}
func serialiseActions(actions []models.Action) *string {
	return nil
}

type NewsDto struct {
	Id        string
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

type EmailDto struct {
	Id      string
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
