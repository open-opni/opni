package notifier_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"go.uber.org/mock/gomock"
)

var ctrl *gomock.Controller

func TestNotifier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Notifier Suite")
}

var _ = BeforeSuite(func() {
	ctrl = gomock.NewController(GinkgoT())
})
