package models

import (
	"fdsim/conf"
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
	Action  *Actionable
}

func NewEmailNoLinks(from, subject, body string, date time.Time) *Email {
	return &Email{
		Id:      emailIdGenerator(),
		Sender:  from,
		Subject: subject,
		Body:    body,
		Date:    date,
		Links:   []Link{},
	}
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

func NewEmailWithAction(from, subject, body string, date time.Time, links []Link, action *Actionable) *Email {
	return &Email{
		Id:      emailIdGenerator(),
		Sender:  from,
		Subject: subject,
		Body:    body,
		Date:    date,
		Links:   links,
		Action:  action,
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

func (e *Email) Answer(choice *Choosable) {
	if e.Action == nil {
		return
	}

	e.Action.Decide(choice)
}

func (e Email) String() string {
	return fmt.Sprintf("%s - %s - %s", e.Sender, e.Date.Format(conf.DateFormatShort), e.Subject)
}
