package caching_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
)

func TestCaching(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Caching Suite")
}
