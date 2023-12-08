package k8sutil_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/pkg/test/testk8s"
	"github.com/open-panoptes/opni/pkg/util/k8sutil"
	"github.com/samber/lo"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

var _ = Describe("Cluster Utils", Ordered, Label("unit", "slow"), func() {
	var restConfig *rest.Config
	var kubeconfigPath string
	BeforeAll(func() {
		var err error
		ctx, ca := context.WithCancel(context.Background())
		restConfig, _, err = testk8s.StartK8s(ctx, nil, runtime.NewScheme())
		Expect(err).NotTo(HaveOccurred())
		tempFile, err := os.CreateTemp("", "test-kubeconfig")
		Expect(err).NotTo(HaveOccurred())
		kubeconfigPath = tempFile.Name()
		apiConfig := api.NewConfig()
		apiConfig.CurrentContext = "test"
		apiConfig.Clusters[apiConfig.CurrentContext] = &api.Cluster{
			Server:                   restConfig.Host,
			CertificateAuthorityData: restConfig.CAData,
		}
		apiConfig.AuthInfos[apiConfig.CurrentContext] = &api.AuthInfo{
			ClientCertificateData: restConfig.CertData,
			ClientKeyData:         restConfig.KeyData,
		}
		apiConfig.Contexts[apiConfig.CurrentContext] = &api.Context{
			Cluster:  apiConfig.CurrentContext,
			AuthInfo: apiConfig.CurrentContext,
		}
		err = clientcmd.WriteToFile(*apiConfig, kubeconfigPath)
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(ca)
	})
	Describe("NewK8sClient", func() {
		When("a kubeconfig is given", func() {
			It("should create the client from the kubeconfig", func() {
				_, err := k8sutil.NewK8sClient(k8sutil.ClientOptions{
					Kubeconfig: &kubeconfigPath,
				})
				Expect(err).NotTo(HaveOccurred())
			})
		})
		When("a REST config is given", func() {
			It("should create the client from the REST config", func() {
				_, err := k8sutil.NewK8sClient(k8sutil.ClientOptions{
					RestConfig: restConfig,
				})
				Expect(err).NotTo(HaveOccurred())
			})
		})
		When("neither a kubeconfig nor a REST config is given", func() {
			It("should create the client from the in-cluster config", func() {
				_, err := k8sutil.NewK8sClient(k8sutil.ClientOptions{})
				Expect(err).To(MatchError(rest.ErrNotInCluster))
			})
		})
		It("should handle errors", func() {
			_, err := k8sutil.NewK8sClient(k8sutil.ClientOptions{
				Kubeconfig: lo.ToPtr("/dev/null"),
			})
			Expect(err).To(HaveOccurred())
			_, err = k8sutil.NewK8sClient(k8sutil.ClientOptions{
				Kubeconfig: lo.ToPtr("/does/not/exist"),
			})
			Expect(err).To(HaveOccurred())
		})
	})
})
