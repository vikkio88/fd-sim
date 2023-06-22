package models

import "fmt"

type Email struct {
	Read    bool
	Sender  string
	Subject string
	Body    string
	Links   []Link
	Actions []Action
}

func NewEmail(from, subject, body string, links []Link) *Email {
	return &Email{
		Sender:  from,
		Subject: subject,
		Body:    body,
		Links:   links,
	}
}

func (e Email) String() string {
	return fmt.Sprintf("%s - %s", e.Sender, e.Subject)
}
