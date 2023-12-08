package patch_test

import (
	"bytes"
	"encoding/hex"
	"os"
	"sync"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	controlv1 "github.com/open-panoptes/opni/pkg/apis/control/v1"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"github.com/open-panoptes/opni/pkg/test/testutil"
	"github.com/open-panoptes/opni/pkg/update/patch"
	"github.com/spf13/afero"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/sync/errgroup"
)

func TestPatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Patch Suite")
}

var (
	test1Path    = "github.com/open-panoptes/opni/pkg/test/testdata/patch/test1"
	test1Package = "urn:opni:plugin:binary:github.com/open-panoptes/opni/pkg/test/testdata/patch/test1"
	test2Path    = "github.com/open-panoptes/opni/pkg/test/testdata/patch/test2"
	test2Package = "urn:opni:plugin:binary:github.com/open-panoptes/opni/pkg/test/testdata/patch/test2"

	testPackages = map[string]string{
		"test1": test1Package,
		"test2": test2Package,
	}

	testBinaries = map[string]map[string][]byte{
		"test1": {},
		"test2": {},
	}

	testBinaryDigests = map[string]map[string]string{
		"test1": {},
		"test2": {},
	}

	test1v1tov2Patch = new(bytes.Buffer)
	test2v1tov2Patch = new(bytes.Buffer)

	testPatches = map[string]*bytes.Buffer{
		"test1": test1v1tov2Patch,
		"test2": test2v1tov2Patch,
	}
)

var osfs = afero.Afero{Fs: afero.NewOsFs()}

func b2sum(fs afero.Afero, filename string) string {
	contents, err := fs.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	sum := blake2b.Sum256(contents)
	return hex.EncodeToString(sum[:])
}

var v1Manifest *controlv1.PluginArchive
var v2Manifest *controlv1.PluginArchive

var ctrl *gomock.Controller

var _ = BeforeSuite(func() {
	ctrl = gomock.NewController(GinkgoT())

	var eg errgroup.Group
	var mu sync.Mutex

	eg.Go(func() error {
		test1v1BinaryPath, err := gexec.Build(test1Path, "-tags=v1")
		if err != nil {
			return err
		}
		mu.Lock()
		testBinaries["test1"]["v1"] = testutil.Must(os.ReadFile(test1v1BinaryPath))
		testBinaryDigests["test1"]["v1"] = b2sum(osfs, test1v1BinaryPath)
		mu.Unlock()
		return err
	})
	eg.Go(func() error {
		test1v2BinaryPath, err := gexec.Build(test1Path, "-tags=v2")
		if err != nil {
			return err
		}
		mu.Lock()
		testBinaries["test1"]["v2"] = testutil.Must(os.ReadFile(test1v2BinaryPath))
		testBinaryDigests["test1"]["v2"] = b2sum(osfs, test1v2BinaryPath)
		mu.Unlock()
		return err
	})
	eg.Go(func() error {
		test2v1BinaryPath, err := gexec.Build(test2Path, "-tags=v1")
		if err != nil {
			return err
		}
		mu.Lock()
		testBinaries["test2"]["v1"] = testutil.Must(os.ReadFile(test2v1BinaryPath))
		testBinaryDigests["test2"]["v1"] = b2sum(osfs, test2v1BinaryPath)
		mu.Unlock()
		return err
	})
	eg.Go(func() error {
		test2v2BinaryPath, err := gexec.Build(test2Path, "-tags=v2")
		if err != nil {
			return err
		}
		mu.Lock()
		testBinaries["test2"]["v2"] = testutil.Must(os.ReadFile(test2v2BinaryPath))
		testBinaryDigests["test2"]["v2"] = b2sum(osfs, test2v2BinaryPath)
		mu.Unlock()
		return err
	})
	Expect(eg.Wait()).To(Succeed())

	patcher := patch.BsdiffPatcher{}
	eg = errgroup.Group{}

	eg.Go(func() error {
		return patcher.GeneratePatch(
			bytes.NewReader(testBinaries["test1"]["v1"]),
			bytes.NewReader(testBinaries["test1"]["v2"]),
			test1v1tov2Patch,
		)
	})

	eg.Go(func() error {
		return patcher.GeneratePatch(
			bytes.NewReader(testBinaries["test2"]["v1"]),
			bytes.NewReader(testBinaries["test2"]["v2"]),
			test2v1tov2Patch,
		)
	})

	Expect(eg.Wait()).To(Succeed())

	v1Manifest = &controlv1.PluginArchive{
		Items: []*controlv1.PluginArchiveEntry{
			{
				Metadata: &controlv1.UpdateManifestEntry{
					Package: test1Package,
					Path:    "test1",
					Digest:  testBinaryDigests["test1"]["v1"],
				},
				Data: testBinaries["test1"]["v1"],
			},
			{
				Metadata: &controlv1.UpdateManifestEntry{
					Package: test2Package,
					Path:    "test2",
					Digest:  testBinaryDigests["test2"]["v1"],
				},
				Data: testBinaries["test2"]["v1"],
			},
		},
	}
	v2Manifest = &controlv1.PluginArchive{
		Items: []*controlv1.PluginArchiveEntry{
			{
				Metadata: &controlv1.UpdateManifestEntry{
					Package: test1Package,
					Path:    "test1",
					Digest:  testBinaryDigests["test1"]["v2"],
				},
				Data: testBinaries["test1"]["v2"],
			},
			{
				Metadata: &controlv1.UpdateManifestEntry{
					Package: test2Package,
					Path:    "test2",
					Digest:  testBinaryDigests["test2"]["v2"],
				},
				Data: testBinaries["test2"]["v2"],
			},
		},
	}

	DeferCleanup(func() {
		gexec.CleanupBuildArtifacts()
	})
})
