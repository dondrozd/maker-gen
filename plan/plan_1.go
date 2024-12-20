package plan

// the purpose of this file is to provide a plan to test my scenarios

import (
	"database/sql"
	"time"

	"github.com/dondrozd/maker-gen/example"
)

type ExampleOneModifier func(*example.ExampleOne)

type ExampleOneTemplate interface {
	ButWith(modifiers ...ExampleOneModifier) ExampleOneTemplate
	Make() *example.ExampleOne
}

type templateExampleOne struct {
	subject *example.ExampleOne
}

func NewExampleOneTemplate() ExampleOneTemplate {
	return &templateExampleOne{
		subject: new(example.ExampleOne),
	}
}

func NewExampleOneTemplateFrom(subject *example.ExampleOne) ExampleOneTemplate {
	return &templateExampleOne{
		subject: subject,
	}
}

func (t *templateExampleOne) ButWith(modifiers ...ExampleOneModifier) ExampleOneTemplate {
	return &templateExampleOne{
		subject: t.apply(*t.subject, modifiers...),
	}
}

func (t *templateExampleOne) apply(subject example.ExampleOne, modifiers ...ExampleOneModifier) *example.ExampleOne {
	subjectPtr := &subject
	for _, modifier := range modifiers {
		modifier(subjectPtr)
	}

	return subjectPtr
}

func (t *templateExampleOne) Make() *example.ExampleOne {
	return t.subject
}

func WithPublicString(value string) ExampleOneModifier {
	return func(subject *example.ExampleOne) {
		subject.PublicString = value
	}
}

func WithPublicInt(value int64) ExampleOneModifier {
	return func(subject *example.ExampleOne) {
		subject.PublicInt = value
	}
}

func WithPublicTime(value time.Time) ExampleOneModifier {
	return func(subject *example.ExampleOne) {
		subject.PublicTime = value
	}
}

func WithPublicNullString(value sql.NullString) ExampleOneModifier {
	return func(subject *example.ExampleOne) {
		subject.PublicNullString = value
	}
}
