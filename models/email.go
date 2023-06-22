package models

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

const emailInMemoryId = "emailId"

func emailIdGenerator() string {
	return fmt.Sprintf("%s_%s", emailInMemoryId, ulid.Make())
}

type Email struct {
	Id      string
	Read    bool
	Sender  string
	Subject string
	Body    string
	Date    time.Time
	Links   []Link
	Actions []Action
}

func NewEmail(from, subject, body string, date time.Time, links []Link) *Email {
	return &Email{
		Id:      emailIdGenerator(),
		Sender:  from,
		Subject: subject,
		Body:    body,
		Date:    date,
		Links:   links,
	}
}

func NewEmailWithId(id, from, subject, body string, date time.Time, links []Link) *Email {
	return &Email{
		Id:      id,
		Sender:  from,
		Subject: subject,
		Body:    body,
		Date:    date,
		Links:   links,
	}
}

func (e Email) String() string {
	return fmt.Sprintf("%s - %s - %s", e.Sender, e.Date, e.Subject)
}
