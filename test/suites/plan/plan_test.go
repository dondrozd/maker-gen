package plan_test

import (
	"github.com/dondrozd/maker-gen/example"
	"github.com/dondrozd/maker-gen/plan"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plan", func() {
	Describe("NewExampleOneTemplate", func() {
		var (
			subject      plan.ExampleOneTemplate
			childSubject plan.ExampleOneTemplate
			actual       *example.ExampleOne
		)
		BeforeEach(func() {
			subject = plan.NewExampleOneTemplate()
			actual = nil
		})
		It("should successfully construct", func() {
			Expect(subject).ShouldNot(BeNil())
		})
		It("should successfully make", func() {
			actual = subject.Make()

			Expect(actual).Should(
				HaveField("PublicString", Equal("")),
			)
		})
		It("should successfully set data", func() {
			actual = subject.ButWith(
				plan.WithPublicString("TEST_DATA"),
			).Make()

			Expect(actual).Should(
				HaveField("PublicString", Equal("TEST_DATA")),
			)
		})
		Describe("with derived template", func() {
			BeforeEach(func() {
				subject = subject.ButWith(
					plan.WithPublicString("TEST_DATA"),
				)
				childSubject = subject.ButWith(plan.WithPublicInt(2))
			})
			It("parent template should have public int 0", func() {
				actual = subject.Make()
				Expect(actual).Should(SatisfyAll(
					HaveField("PublicString", Equal("TEST_DATA")),
					HaveField("PublicInt", Equal(int64(0))),
				))
			})
			It("child template should have public int 1", func() {
				actual = childSubject.Make()
				Expect(actual).Should(SatisfyAll(
					HaveField("PublicString", Equal("TEST_DATA")),
					HaveField("PublicInt", Equal(int64(2))),
				))
			})
		})
	})
})
