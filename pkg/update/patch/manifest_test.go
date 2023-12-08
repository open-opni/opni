package patch_test

import (
	"path/filepath"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	controlv1 "github.com/open-panoptes/opni/pkg/apis/control/v1"
	"github.com/open-panoptes/opni/pkg/plugins"
	"github.com/open-panoptes/opni/pkg/test/memfs"
	"github.com/open-panoptes/opni/pkg/test/testlog"
	"github.com/open-panoptes/opni/pkg/update/patch"
	"github.com/open-panoptes/opni/pkg/util"
	"github.com/spf13/afero"
)

func NewDigest4() (string, string, string, string) {
	return uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()
}

type testLeftJoinData struct {
	leftConfig  *controlv1.UpdateManifest
	rightConfig *controlv1.UpdateManifest
	expected    *controlv1.PatchList
}

var _ = Describe("Patch Manifest Operations", Label("unit"), func() {
	When("Receiving two sets of manifests", func() {
		hash1, hash2, hash3, hash4 := NewDigest4()

		metrics := &controlv1.UpdateManifest{
			Items: []*controlv1.UpdateManifestEntry{
				{
					Package: "metrics",
					Path:    "plugin_metrics",
					Digest:  hash1,
				},
			},
		}
		metricsRenamed := func() *controlv1.UpdateManifest {
			t := util.ProtoClone(metrics)
			t.Items[0].Path = "renamed"
			return t
		}()

		metricsRenamedPackage := func() *controlv1.UpdateManifest {
			t := util.ProtoClone(metrics)
			t.Items[0].Package = "modeltraining"
			t.Items[0].Digest = hash4 //hash is also necessarily changed when Package is changed
			return t
		}()

		metricsUpdate := func() *controlv1.UpdateManifest {
			t := util.ProtoClone(metrics)
			t.Items[0].Digest = hash2
			return t
		}()

		_ = &controlv1.UpdateManifest{
			Items: []*controlv1.UpdateManifestEntry{
				{
					Package: "metrics",
					Path:    "plugin_metrics",
					Digest:  hash3,
				},
			},
		}

		addlogging := &controlv1.UpdateManifest{
			Items: []*controlv1.UpdateManifestEntry{
				metrics.Items[0],
				{
					Package: "logging",
					Path:    "plugin_logging",
					Digest:  hash4,
				},
			},
		}

		removeLogging := &controlv1.UpdateManifest{
			Items: []*controlv1.UpdateManifestEntry{
				addlogging.Items[0],
			},
		}
		It("should determine when to update plugins", func() {
			// only the hashes differ
			scenario1 := testLeftJoinData{
				leftConfig:  metricsUpdate,
				rightConfig: metrics,
				expected: &controlv1.PatchList{
					Items: []*controlv1.PatchSpec{
						{
							Op:        controlv1.PatchOp_Update,
							Package:   metricsUpdate.Items[0].GetPackage(),
							OldDigest: metrics.Items[0].GetDigest(),
							NewDigest: metricsUpdate.Items[0].GetDigest(),
							Path:      metricsUpdate.Items[0].GetPath(),
						},
					},
				},
			}
			res := patch.LeftJoinOn(scenario1.leftConfig, scenario1.rightConfig)
			By("checking that the output from the two manifests is well formed")
			Expect(res.Validate()).To(Succeed())
			Expect(len(res.Items)).To(Equal(len(scenario1.expected.Items)))
			for i, item := range res.Items {
				By("checking that the generated operation is update")
				Expect(item.Op).To(Equal(controlv1.PatchOp_Update))
				By("checking that the metadata is generated correctly")
				Expect(item.Package).To(Equal(scenario1.expected.Items[i].Package))
				Expect(item.NewDigest).To(Equal(scenario1.expected.Items[i].NewDigest))
				Expect(item.OldDigest).To(Equal(scenario1.expected.Items[i].OldDigest))
				Expect(item.Path).To(Equal(scenario1.expected.Items[i].Path))
			}
		})

		It("should determine when to rename plugins", func() {
			scenario2 := testLeftJoinData{
				leftConfig:  metricsRenamed,
				rightConfig: metrics,
				expected: &controlv1.PatchList{
					Items: []*controlv1.PatchSpec{
						{
							Op:        controlv1.PatchOp_Rename,
							Package:   metrics.Items[0].GetPackage(),
							OldDigest: metrics.Items[0].GetDigest(),
							NewDigest: metricsRenamed.Items[0].GetDigest(),
							Path:      "plugin_metrics",
							Data:      []byte("renamed"),
						},
					},
				},
			}
			res := patch.LeftJoinOn(scenario2.leftConfig, scenario2.rightConfig)
			By("checking that the output from the two manifests is well formed")
			Expect(res.Validate()).To(Succeed())
			Expect(len(res.Items)).To(Equal(len(scenario2.expected.Items)))
			for i, item := range res.Items {
				By("checking that the generated operation is rename")
				Expect(item.Op).To(Equal(controlv1.PatchOp_Rename))
				By("checking that the metadata is generated correctly")
				Expect(item.Package).To(Equal(scenario2.expected.Items[i].Package))
				Expect(item.NewDigest).To(Equal(scenario2.expected.Items[i].NewDigest))
				Expect(item.OldDigest).To(Equal(scenario2.expected.Items[i].OldDigest))
				Expect(item.Path).To(Equal(scenario2.expected.Items[i].Path))
			}
		})

		It("should determine when to add plugins", func() {
			scenario3 := testLeftJoinData{
				leftConfig:  addlogging,
				rightConfig: metricsUpdate,
				expected: &controlv1.PatchList{
					Items: []*controlv1.PatchSpec{
						{
							Op:        controlv1.PatchOp_Update,
							Package:   addlogging.Items[0].GetPackage(),
							OldDigest: metricsUpdate.Items[0].GetDigest(),
							NewDigest: addlogging.Items[0].GetDigest(),
							Path:      addlogging.Items[0].GetPath(),
						},
						{
							Op:        controlv1.PatchOp_Create,
							Package:   addlogging.Items[1].GetPackage(),
							OldDigest: "",
							NewDigest: addlogging.Items[1].GetDigest(),
							Path:      addlogging.Items[1].GetPath(),
						},
					},
				},
			}
			res := patch.LeftJoinOn(scenario3.leftConfig, scenario3.rightConfig)
			By("checking that the output from the two manifests is well formed")
			Expect(res.Validate()).To(Succeed())
			Expect(len(res.Items)).To(Equal(len(scenario3.expected.Items)))
			for i, item := range res.Items {
				By("checking that the generated operation is create or update")
				Expect(item.Op).To(Equal(scenario3.expected.Items[i].Op))
				By("checking that the metadata is generated correctly")
				Expect(item.Package).To(Equal(scenario3.expected.Items[i].Package))
				Expect(item.NewDigest).To(Equal(scenario3.expected.Items[i].NewDigest))
				Expect(item.OldDigest).To(Equal(scenario3.expected.Items[i].OldDigest))
				Expect(item.Path).To(Equal(scenario3.expected.Items[i].Path))
			}
		})

		It("should determine when to remove plugins", func() {
			scenario4 := testLeftJoinData{
				leftConfig:  removeLogging,
				rightConfig: addlogging,
				expected: &controlv1.PatchList{
					Items: []*controlv1.PatchSpec{
						{
							Op:        controlv1.PatchOp_Remove,
							Package:   addlogging.Items[1].GetPackage(),
							OldDigest: addlogging.Items[1].GetDigest(),
							NewDigest: "",
							Path:      addlogging.Items[1].GetPath(),
						},
					},
				},
			}
			res := patch.LeftJoinOn(scenario4.leftConfig, scenario4.rightConfig)
			By("checking that the output from the two manifests is well formed")
			Expect(res.Validate()).To(Succeed())
			Expect(len(res.Items)).To(Equal(len(scenario4.expected.Items)))
			for i, item := range res.Items {
				By("checking that the generated operation is remove")
				Expect(item.Op).To(Equal(controlv1.PatchOp_Remove))
				By("checking that the metadata is generated correctly")
				Expect(item.Package).To(Equal(scenario4.expected.Items[i].Package))
				Expect(item.NewDigest).To(Equal(scenario4.expected.Items[i].NewDigest))
				Expect(item.OldDigest).To(Equal(scenario4.expected.Items[i].OldDigest))
				Expect(item.Path).To(Equal(scenario4.expected.Items[i].Path))
			}
		})

		It("should determine when to replace plugins with new contents", func() {
			scenario5 := testLeftJoinData{
				leftConfig:  metricsRenamedPackage,
				rightConfig: metrics,
				expected: &controlv1.PatchList{
					Items: []*controlv1.PatchSpec{
						{
							Op:        controlv1.PatchOp_Create,
							Package:   metricsRenamedPackage.Items[0].GetPackage(),
							OldDigest: "",
							NewDigest: metricsRenamedPackage.Items[0].GetDigest(),
							Path:      metricsRenamedPackage.Items[0].GetPath(),
						},
						{
							Op:        controlv1.PatchOp_Remove,
							Package:   metrics.Items[0].GetPackage(),
							OldDigest: metrics.Items[0].GetDigest(),
							NewDigest: "",
							Path:      metrics.Items[0].GetPath(),
						},
					},
				},
			}
			res := patch.LeftJoinOn(scenario5.leftConfig, scenario5.rightConfig)
			By("checking that the output from the two manifests is well formed")
			Expect(res.Validate()).To(Succeed())
			Expect(len(res.Items)).To(Equal(len(scenario5.expected.Items)))
			for i, item := range res.Items {
				By("checking that the generated operation is create or remove")
				Expect(item.Op).To(Equal(scenario5.expected.Items[i].Op))
				By("checking that the metadata is generated correctly")
				Expect(item.Package).To(Equal(scenario5.expected.Items[i].Package))
				Expect(item.NewDigest).To(Equal(scenario5.expected.Items[i].NewDigest))
				Expect(item.OldDigest).To(Equal(scenario5.expected.Items[i].OldDigest))
				Expect(item.Path).To(Equal(scenario5.expected.Items[i].Path))
			}
		})
	})
})

var _ = Describe("Filesystem Discovery", Ordered, func() {
	var fsys afero.Afero
	tmpDir := "/tmp"
	It("should discover plugins from the filesystem", func() {
		fsys = afero.Afero{Fs: afero.NewMemMapFs()}

		Expect(fsys.WriteFile(filepath.Join(tmpDir, patch.PluginsDir, "plugin_test1"), testBinaries["test1"]["v1"], 0755)).To(Succeed())
		Expect(fsys.WriteFile(filepath.Join(tmpDir, patch.PluginsDir, "plugin_test2"), testBinaries["test2"]["v1"], 0755)).To(Succeed())

		mv1, err := patch.GetFilesystemPlugins(plugins.DiscoveryConfig{
			Dir:    filepath.Join(tmpDir, patch.PluginsDir),
			Fs:     fsys,
			Logger: testlog.Log,
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(mv1.Items).To(HaveLen(2))
		Expect(mv1.Items[0].Metadata.Package).To(Equal(test1Package))
		Expect(mv1.Items[1].Metadata.Package).To(Equal(test2Package))
		Expect(mv1.Items[0].Metadata.Digest).To(Equal(v1Manifest.Items[0].Metadata.Digest))
		Expect(mv1.Items[1].Metadata.Digest).To(Equal(v1Manifest.Items[1].Metadata.Digest))
		Expect(mv1.Items[0].Metadata.Path).To(Equal("plugin_test1"))
		Expect(mv1.Items[1].Metadata.Path).To(Equal("plugin_test2"))
	})
	When("a plugin has invalid contents", func() {
		It("should log an error and skip it", func() {
			af := afero.Afero{
				Fs: afero.NewMemMapFs(),
			}

			Expect(af.WriteFile(filepath.Join(tmpDir, "plugin_test1"), testBinaries["test1"]["v1"], 0755)).To(Succeed())
			Expect(af.WriteFile(filepath.Join(tmpDir, "plugin_test2"), []byte("invalid"), 0755)).To(Succeed())

			mv1, err := patch.GetFilesystemPlugins(plugins.DiscoveryConfig{
				Dir:    tmpDir,
				Fs:     af,
				Logger: testlog.Log,
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(mv1.Items).To(HaveLen(1))
			Expect(mv1.Items[0].Metadata.Package).To(Equal(test1Package))
		})
	})
	When("a plugin cannot be opened for reading", func() {
		It("should log an error and skip it", func() {
			af := afero.Afero{
				Fs: memfs.NewModeAwareMemFs(),
			}
			af.MkdirAll(tmpDir, 0755)

			Expect(af.WriteFile(filepath.Join(tmpDir, "plugin_test1"), testBinaries["test1"]["v1"], 0755)).To(Succeed())
			Expect(af.WriteFile(filepath.Join(tmpDir, "plugin_test2"), testBinaries["test2"]["v1"], 0200)).To(Succeed())

			mv1, err := patch.GetFilesystemPlugins(plugins.DiscoveryConfig{
				Dir:    tmpDir,
				Fs:     af,
				Logger: testlog.Log,
			})

			Expect(err).NotTo(HaveOccurred())
			Expect(mv1.Items).To(HaveLen(1))

			Expect(mv1.Items[0].Metadata.Package).To(Equal(test1Package))
		})
	})
})
