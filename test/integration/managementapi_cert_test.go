package integration_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/emptypb"

	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/test"
	_ "github.com/open-panoptes/opni/plugins/example/test"
)

//#region Test Setup

var _ = Describe("Management API Cerificate Management Tests", Ordered, Label("integration"), func() {
	var environment *test.Environment
	var client managementv1.ManagementClient
	BeforeAll(func() {
		environment = &test.Environment{}
		Expect(environment.Start()).To(Succeed())
		DeferCleanup(environment.Stop)
		client = environment.NewManagementClient()
	})

	//#endregion

	//#region Happy Path Tests

	It("can retrieve full certification chain information", func() {
		certsInfo, err := client.CertsInfo(context.Background(), &emptypb.Empty{})
		Expect(err).NotTo(HaveOccurred())

		leaf := certsInfo.Chain[0]
		root := certsInfo.Chain[len(certsInfo.Chain)-1]

		Expect(root.Issuer).To(Equal("CN=Example Root CA"))
		Expect(root.Subject).To(Equal("CN=Example Root CA"))
		Expect(root.IsCA).To(BeTrue())
		Expect(root.Fingerprint).NotTo(BeEmpty())

		Expect(leaf.Issuer).To(Equal("CN=Example Root CA"))
		Expect(leaf.Subject).To(Equal("CN=leaf"))
		Expect(leaf.IsCA).To(BeFalse())
		Expect(leaf.Fingerprint).NotTo(BeEmpty())
	})

	//#endregion

	//#region Edge Case Tests

	//TODO: Add Certificate Edge Case Tests

	//#endregion
})
