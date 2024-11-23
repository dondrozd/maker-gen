package plan

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

type exampleOneTemplate struct {
	subject *example.ExampleOne
}

func NewExampleOneTemplate() ExampleOneTemplate {
	return &exampleOneTemplate{
		subject: new(example.ExampleOne),
	}
}

func NewExampleOneTemplateFrom(subject *example.ExampleOne) ExampleOneTemplate {
	return &exampleOneTemplate{
		subject: subject,
	}
}

func (t *exampleOneTemplate) ButWith(modifiers ...ExampleOneModifier) ExampleOneTemplate {
	return &exampleOneTemplate{
		subject: t.apply(*t.subject, modifiers...),
	}
}

func (t *exampleOneTemplate) apply(subject example.ExampleOne, modifiers ...ExampleOneModifier) *example.ExampleOne {
	subjectPtr := &subject
	for _, modifier := range modifiers {
		modifier(subjectPtr)
	}

	return subjectPtr
}

func (t *exampleOneTemplate) Make() *example.ExampleOne {
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
