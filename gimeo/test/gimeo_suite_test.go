package gimeo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGimeo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gimeo Suite")
}
