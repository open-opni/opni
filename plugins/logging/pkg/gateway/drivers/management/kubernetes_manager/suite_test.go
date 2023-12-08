package kubernetes_manager_test

import (
	"context"
	"testing"

	"github.com/kralicky/kmatch"
	"github.com/nats-io/nats.go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/apis"
	"github.com/open-panoptes/opni/pkg/test"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"github.com/open-panoptes/opni/pkg/test/testk8s"
	"github.com/open-panoptes/opni/pkg/test/testruntime"
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
	nc         *nats.Conn
)

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logging Plugin Suite")
}

var _ = BeforeSuite(func() {
	testruntime.IfIntegration(func() {
		env := test.Environment{}

		ctx, ca := context.WithCancel(context.Background())
		var err error

		restConfig, scheme, err = testk8s.StartK8s(ctx, []string{
			"../../../../../../../config/crd/bases",
			"../../../../../../../config/crd/opensearch",
			"../../../../../../../test/resources",
		}, apis.NewScheme())
		Expect(err).NotTo(HaveOccurred())

		nc, err = env.StartEmbeddedJetstream()
		Expect(err).NotTo(HaveOccurred())

		k8sClient, err = client.New(restConfig, client.Options{
			Scheme: scheme,
		})
		Expect(err).NotTo(HaveOccurred())
		kmatch.SetDefaultObjectClient(k8sClient)

		k8sManager = testk8s.StartManager(ctx, restConfig, scheme)

		DeferCleanup(func() {
			By("tearing down the test environment")
			nc.Close()
			err := env.Stop()
			Expect(err).NotTo(HaveOccurred())

			ca()
		})
	})
})
