package aiops_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/apis"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"github.com/open-panoptes/opni/pkg/test/testk8s"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	k8sClient  client.Client
	restConfig *rest.Config
	scheme     *runtime.Scheme
	k8sManager ctrl.Manager
)

func TestAPIs(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	SetDefaultEventuallyTimeout(30 * time.Second)
	SetDefaultEventuallyPollingInterval(100 * time.Millisecond)
	RegisterFailHandler(Fail)
	RunSpecs(t, "AI Plugin Suite")
}

var _ = BeforeSuite(func() {
	var err error
	ctx, ca := context.WithCancel(context.Background())
	restConfig, scheme, err = testk8s.StartK8s(ctx, []string{
		"../../../config/crd/bases",
		"../../../config/crd/opensearch",
		"../../../test/resources",
	}, apis.NewScheme())
	Expect(err).NotTo(HaveOccurred())

	DeferCleanup(ca)

	k8sClient, err = client.New(restConfig, client.Options{
		Scheme: scheme,
	})
	Expect(err).NotTo(HaveOccurred())

	k8sManager = testk8s.StartManager(ctx, restConfig, scheme)
})
