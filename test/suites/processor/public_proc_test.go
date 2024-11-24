package processor_test

import (
	"github.com/dondrozd/maker-gen/model"
	"github.com/dondrozd/maker-gen/processor"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PublicProc", func() {
	Describe("RenderMaker", func() {
		var (
			actual model.MakerModel
			err    error
		)

		It("should return error when struct was not found", func() {
			model := model.GoFileModel{}

			actual, err = processor.PublicProc(model, "")

			Expect(err).Should(MatchError("could not find struct with name: "))
			Expect(actual).Should(HaveField("PackageName", Equal("")))
		})
	})
})
