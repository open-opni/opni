package prometheusrule_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/pkg/rules"
	"github.com/open-panoptes/opni/pkg/rules/prometheusrule"
	"github.com/open-panoptes/opni/pkg/test/testk8s"
	"github.com/open-panoptes/opni/pkg/test/testlog"
	"github.com/open-panoptes/opni/pkg/util/notifier"
	monitoringcoreosv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/samber/lo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Prometheus Rule Group Discovery", Ordered, Label("integration", "slow"), func() {
	alertGroup := []monitoringv1.RuleGroup{
		{
			Name:     "test-0",
			Interval: lo.ToPtr[monitoringv1.Duration]("1m"),
			Rules: []monitoringv1.Rule{
				{
					Record:      "",
					Alert:       "nothing",
					Expr:        intstr.FromString("burger"),
					For:         lo.ToPtr[monitoringv1.Duration]("5m"),
					Labels:      map[string]string{},
					Annotations: map[string]string{},
				},
			},
		},
	}
	testGroups1 := []monitoringv1.RuleGroup{
		{
			Name: "test",
			Rules: []monitoringv1.Rule{
				{
					Record: "foo",
					Expr:   intstr.FromString("foo"),
					Labels: map[string]string{"foo": "bar"},
				},
			},
			Interval: lo.ToPtr[monitoringv1.Duration]("1m"),
		},
		{
			Name: "test2",
			Rules: []monitoringv1.Rule{
				{
					Record: "bar",
					Expr:   intstr.FromString("bar"),
					Labels: map[string]string{"bar": "baz"},
				},
			},
			Interval: lo.ToPtr[monitoringv1.Duration]("2m"),
		},
	}
	testGroups2 := make([]monitoringv1.RuleGroup, len(testGroups1))
	for i, group := range testGroups1 {
		testGroups2[i] = *group.DeepCopy()
	}
	testGroups2[0].Name = "test3"
	testGroups2[1].Name = "test4"
	testGroups3 := []monitoringv1.RuleGroup{
		{
			Name: "test5",
			Rules: []monitoringv1.Rule{
				{
					Record: "foo",
					Expr:   intstr.FromString("foo"),
					For:    lo.ToPtr[monitoringv1.Duration]("invalid"),
				},
				{
					Record: "bar",
					Expr:   intstr.FromString("bar"),
					For:    lo.ToPtr[monitoringv1.Duration]("1m"), // not allowed in recording rule
				},
				{
					Record: "baz",
					Expr:   intstr.FromString("baz"),
				},
			},
			Interval: lo.ToPtr[monitoringv1.Duration]("1m"),
		},
		{
			Name: "test6",
			Rules: []monitoringv1.Rule{
				{
					Record: "baz",
					Expr:   intstr.FromString("baz"),
					For:    lo.ToPtr[monitoringv1.Duration]("2m"),
				},
			},
			Interval: lo.ToPtr[monitoringv1.Duration]("invalid"),
		},
	}

	var k8sClient client.Client
	var finder notifier.Finder[rules.RuleGroup]
	BeforeAll(func() {
		ctx, ca := context.WithCancel(context.Background())

		s := scheme.Scheme
		monitoringcoreosv1.AddToScheme(s)

		restConfig, _, err := testk8s.StartK8s(ctx, []string{"testdata/crds"}, s)
		Expect(err).NotTo(HaveOccurred())

		k8sClient, err = client.New(restConfig, client.Options{
			Scheme: s,
		})
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(ca)

		Expect(k8sClient.Create(context.Background(), &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test1",
			},
		})).To(Succeed())
		Expect(k8sClient.Create(context.Background(), &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test2",
			},
		})).To(Succeed())

		finder = prometheusrule.NewPrometheusRuleFinder(k8sClient, prometheusrule.WithLogger(testlog.Log))
	})

	It("should initially find no groups", func() {
		groups, err := finder.Find(context.Background())
		Expect(err).NotTo(HaveOccurred())
		Expect(groups).To(BeEmpty())
	})

	When("creating a PrometheusRule", func() {
		It("should find the groups in the PrometheusRule", func() {
			Expect(k8sClient.Create(context.Background(), &monitoringv1.PrometheusRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test1",
				},
				Spec: monitoringv1.PrometheusRuleSpec{
					Groups: testGroups1,
				},
			})).To(Succeed())

			By("finding the groups in the new PrometheusRule")
			groups, err := finder.Find(context.Background())
			Expect(err).NotTo(HaveOccurred())
			Expect(groups).To(HaveLen(2))
			Expect([]string{
				groups[0].Name,
				groups[1].Name,
			}).To(ConsistOf("test", "test2"))
		})
	})
	When("creating a PrometheusRule in a different namespace", func() {
		It("should find PrometheusRules in both namespaces", func() {
			Expect(k8sClient.Create(context.Background(), &monitoringv1.PrometheusRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test2",
				},
				Spec: monitoringv1.PrometheusRuleSpec{
					Groups: testGroups2,
				},
			})).To(Succeed())
			groups, err := finder.Find(context.Background())

			Expect(err).NotTo(HaveOccurred())
			Expect(groups).To(HaveLen(4))

			Expect([]string{
				groups[0].Name,
				groups[1].Name,
				groups[2].Name,
				groups[3].Name,
			}).To(ContainElements("test", "test2", "test3", "test4"))
		})
	})
	When("creating a PrometheusRule with invalid contents", func() {
		It("should skip invalid rules or groups", func() {
			Expect(k8sClient.Create(context.Background(), &monitoringv1.PrometheusRule{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-invalid",
					Namespace: "test2",
				},
				Spec: monitoringv1.PrometheusRuleSpec{
					Groups: testGroups3,
				},
			})).To(Succeed())

			groups, err := finder.Find(context.Background())
			Expect(err).NotTo(HaveOccurred())

			// It should skip test6 and 2 of the rules in test5
			Expect(groups).To(HaveLen(5))
			Expect([]string{
				groups[0].Name,
				groups[1].Name,
				groups[2].Name,
				groups[3].Name,
				groups[4].Name,
			}).To(ConsistOf("test", "test2", "test3", "test4", "test5"))
			for _, group := range groups {
				if group.Name == "test5" {
					Expect(group.Rules).To(HaveLen(1))
					Expect(group.Rules[0].Record.Value).To(Equal("baz"))
				}
			}
		})
	})
	It("should allow specifying namespaces to search in", func() {
		finder1 := prometheusrule.NewPrometheusRuleFinder(k8sClient, prometheusrule.WithNamespaces("test1"))
		groups, err := finder1.Find(context.Background())
		Expect(err).NotTo(HaveOccurred())
		Expect(groups).To(HaveLen(2))
		Expect([]string{
			groups[0].Name,
			groups[1].Name,
		}).To(ContainElements("test", "test2"))

		finder2 := prometheusrule.NewPrometheusRuleFinder(k8sClient, prometheusrule.WithNamespaces("test2"))
		groups, err = finder2.Find(context.Background())
		Expect(err).NotTo(HaveOccurred())
		Expect(groups).To(HaveLen(3))
		Expect([]string{
			groups[0].Name,
			groups[1].Name,
			groups[2].Name,
		}).To(ContainElements("test3", "test4", "test5"))

		// these should all match both namespaces
		for _, namespaces := range [][]string{
			{""},
			{"test1", "test2"},
			{"test2", "test1", ""},
			{"test1", "test2", "test3"},
		} {
			finder := prometheusrule.NewPrometheusRuleFinder(k8sClient,
				prometheusrule.WithLogger(testlog.Log), prometheusrule.WithNamespaces(namespaces...))
			groups, err = finder.Find(context.Background())
			Expect(err).NotTo(HaveOccurred())
			Expect(groups).To(HaveLen(5))
		}
	})

	It("should ignore alerting rules all together", func() {
		Expect(k8sClient.Create(context.Background(), &monitoringv1.PrometheusRule{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "alerting",
				Namespace: "test2",
			},
			Spec: monitoringv1.PrometheusRuleSpec{
				Groups: alertGroup,
			},
		})).To(Succeed())

		groups, err := finder.Find(context.Background())
		Expect(err).To(Succeed())
		Expect(groups).To(HaveLen(5))
		Expect([]string{
			groups[0].Name,
			groups[1].Name,
			groups[2].Name,
			groups[3].Name,
			groups[4].Name,
		}).To(ConsistOf([]string{"test", "test2", "test3", "test4", "test5"}))
	})
})
