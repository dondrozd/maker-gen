package renderer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRender(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Renderer Suite")
}
