package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("From has to contains @")
	}

	e.email.from = from
	return e
}

func (e *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("To has to contains @")
	}

	e.email.to = to
	return e
}

func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

func (e *EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}

func sendEmailMock(email *email) {
	fmt.Println(email)
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailMock(&builder.email)
}

func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.
			From("aa.com").
			To("b@b.com").
			Subject("Asunto").
			Body("lalalalala")
	})
}
