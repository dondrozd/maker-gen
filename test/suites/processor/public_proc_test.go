package processor_test

import (
	"log/slog"

	"github.com/dondrozd/maker-gen/model"
	"github.com/dondrozd/maker-gen/processor"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PublicProc", func() {
	BeforeEach(func() {
		slog.SetDefault(slog.New(slog.NewTextHandler(ginkgo.GinkgoWriter, nil)))
	})

	Describe("RenderMaker", func() {
		var (
			actual model.MakerModel
			err    error
		)

		It("should return error when struct was not found", func() {
			fileModel := model.GoFileModel{}

			actual, err = processor.PublicProc(fileModel, model.GenerateParams{
				StructName: "",
			})

			Expect(err).Should(MatchError("could not find struct with name: "))
			Expect(actual).Should(HaveField("PackageName", Equal("")))
		})

		Describe("happy path", func() {
			BeforeEach(func() {
				fileModel := model.GoFileModel{
					PackageName: "example",
					ModulePath:  "github.com/dondrozd/maker-gen",
					Structs: []model.StructModel{
						{
							Name: "MyStruct",
							Properties: []model.StructPropertyModel{
								{
									Name: "privateProperty",
									Type: "string",
								},
								{
									Name: "PublicProperty",
									Type: "string",
								},
							},
						},
					},
				}

				actual, err = processor.PublicProc(fileModel, model.GenerateParams{
					StructName: "MyStruct",
				})
			})
			It("should not return error", func() {
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should copy package name", func() {
				Expect(actual).Should(HaveField("PackageName", Equal("example")))
			})
			It("should create import for the struct", func() {
				Expect(actual).Should(HaveField("Imports", ContainElement(
					HaveField("ImportPath", Equal("\"github.com/dondrozd/maker-gen/example\"")),
				)))
			})
			It("should have the struct", func() {
				Expect(actual).Should(HaveField("Structs", ContainElement(
					HaveField("Name", Equal("MyStruct")),
				)))
			})
			It("should have the struct with public properties", func() {
				Expect(actual).Should(HaveField("Structs", ContainElement(SatisfyAll(
					HaveField("Name", Equal("MyStruct")),
					HaveField("Properties", ContainElement(
						HaveField("Name", Equal("PublicProperty")),
					)),
				))))
			})
			It("should not have the struct with private properties", func() {
				Expect(actual).Should(HaveField("Structs", ContainElement(SatisfyAll(
					HaveField("Name", Equal("MyStruct")),
					HaveField("Properties", Not(ContainElement(
						HaveField("Name", Equal("privateProperty")),
					))),
				))))
			})
		})
	})
})
