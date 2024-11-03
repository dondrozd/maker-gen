package plan

import "dondrozd/maker-gen/examples"

type ExampleOneModifier func(*examples.ExampleOne)

type ExampleOneTemplate interface {
	With(modifiers ...ExampleOneModifier) ExampleOneTemplate
	Make() *examples.ExampleOne
}

type exampleOneTemplate struct {
	subject *examples.ExampleOne
}

func NewExampleOneTemplate() ExampleOneTemplate {
	return &exampleOneTemplate{
		subject: new(examples.ExampleOne),
	}
}

func NewExampleOneTemplateFrom(subject *examples.ExampleOne) ExampleOneTemplate {
	return &exampleOneTemplate{
		subject: subject,
	}
}

func (t *exampleOneTemplate) With(modifiers ...ExampleOneModifier) ExampleOneTemplate {
	return &exampleOneTemplate{
		subject: t.apply(*t.subject, modifiers...),
	}
}

func (t *exampleOneTemplate) apply(subject examples.ExampleOne, modifiers ...ExampleOneModifier) *examples.ExampleOne {
	subjectPtr := &subject
	for _, modifier := range modifiers {
		modifier(subjectPtr)
	}

	return subjectPtr
}

func (t *exampleOneTemplate) Make() *examples.ExampleOne {
	return t.subject
}

func WithPublicString(value string) ExampleOneModifier {
	return func(subject *examples.ExampleOne) {
		subject.PublicString = value
	}
}

func WithPublicInt(value int64) ExampleOneModifier {
	return func(subject *examples.ExampleOne) {
		subject.PublicInt = value
	}
}
