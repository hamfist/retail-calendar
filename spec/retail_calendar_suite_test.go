package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRetailCalendar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Retail-Calendar Suite")
}
