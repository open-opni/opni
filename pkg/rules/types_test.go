package rules_test

import (
	"unsafe"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/pkg/test/testdata"
	"github.com/prometheus/prometheus/model/rulefmt"

	"github.com/open-panoptes/opni/pkg/rules"
	"github.com/open-panoptes/opni/pkg/util/notifier"
)

var _ = Describe("Types", Label("unit"), func() {
	It("should clone rule groups", func() {
		rg, errs := rulefmt.Parse(testdata.TestData("prometheus/sample-rules.yaml"))
		Expect(errs).To(BeEmpty())
		groups := rg.Groups
		groupsToInterface := []rules.RuleGroup{}
		for _, group := range groups {
			groupsToInterface = append(groupsToInterface, rules.RuleGroup(group))
		}
		// clone := rules.CloneRuleGroupList(groups)
		clone := notifier.CloneList(groupsToInterface)
		Expect(unsafe.Pointer(&clone)).NotTo(Equal(unsafe.Pointer(&groups)))
		Expect(unsafe.Pointer(&clone[0].Rules)).NotTo(Equal(unsafe.Pointer(&groups[0].Rules)))
		Expect(unsafe.Pointer(&clone[0].Rules[0].Labels)).NotTo(Equal(unsafe.Pointer(&groups[0].Rules[0].Labels)))
		Expect(unsafe.Pointer(&clone[0].Rules[0].Annotations)).NotTo(Equal(unsafe.Pointer(&groups[0].Rules[0].Annotations)))
	})
})
