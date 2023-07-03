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

func unserialiseAction(actions *string) *models.Actionable {
	if actions == nil {
		return nil
	}

	var result models.Actionable
	data := *actions
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("error while loading actions", err)
		return nil
	}

	return &result
}
func serialiseAction(action *models.Actionable) *string {
	if action == nil {
		return nil
	}

	var result string
	data, _ := json.Marshal(action)
	result = string(data)

	return &result
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
	Expires *time.Time
	Links   *string
	Action  *string
}

func DtoFromEmail(email *models.Email) EmailDto {
	return EmailDto{
		Id:      email.Id,
		Read:    email.Read,
		Sender:  email.Sender,
		Subject: email.Subject,
		Body:    email.Body,
		Date:    email.Date,
		Expires: email.Expires,
		Links:   serialiseLinks(email.Links),
		Action:  serialiseAction(email.Action),
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
		Expires: email.Expires,
		Links:   unserialiseLinks(email.Links),
		Action:  unserialiseAction(email.Action),
	}
}
