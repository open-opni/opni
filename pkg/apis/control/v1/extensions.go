package v1

import (
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"slices"

	"github.com/open-panoptes/opni/pkg/urn"
	"golang.org/x/crypto/blake2b"
)

func (m *UpdateManifest) DigestSet() map[string]struct{} {
	hm := map[string]struct{}{}
	for _, v := range m.Items {
		hm[v.GetDigest()] = struct{}{}
	}
	return hm
}

func (m *UpdateManifest) DigestMap() map[string]string {
	hm := map[string]string{}
	for _, v := range m.Items {
		hm[v.GetId()] = v.GetDigest()
	}
	return hm
}

func (m *UpdateManifestEntry) GetId() string {
	return m.GetPackage()
}

func (m *UpdateManifestEntry) DigestBytes() []byte {
	decoded, _ := hex.DecodeString(m.GetDigest())
	return decoded
}

func (m *UpdateManifestEntry) DigestHash() hash.Hash {
	h, _ := blake2b.New256(nil)
	return h
}

func (m *UpdateManifest) Sort() {
	slices.SortFunc(m.Items, func(a, b *UpdateManifestEntry) int {
		return strings.Compare(a.GetPackage(), b.GetPackage())
	})
}

func (a *PluginArchive) Sort() {
	slices.SortFunc(a.Items, func(a, b *PluginArchiveEntry) int {
		return strings.Compare(a.GetMetadata().GetPackage(), b.GetMetadata().GetPackage())
	})
}

func (a *PluginArchive) ToManifest() *UpdateManifest {
	manifest := &UpdateManifest{}
	for _, entry := range a.Items {
		manifest.Items = append(manifest.Items, entry.Metadata)
	}
	return manifest
}

func (l *PatchList) Sort() {
	slices.SortFunc(l.Items, func(a, b *PatchSpec) int {
		if a.GetOp() != b.GetOp() {
			return int(a.GetOp() - b.GetOp())
		}
		if a.GetPackage() != b.GetPackage() {
			return strings.Compare(a.GetPackage(), b.GetPackage())
		}
		return strings.Compare(a.GetPath(), b.GetPath())
	})
}

func (l *PatchList) Summary() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "%d entries", len(l.Items))
	if len(l.Items) > 0 {
		fmt.Fprintf(&builder, ": ")
		updatesByOperation := map[PatchOp]int{}
		for _, entry := range l.Items {
			updatesByOperation[entry.GetOp()]++
		}
		separator := ""
		if updatesByOperation[PatchOp_Create] > 0 {
			fmt.Fprintf(&builder, "%d new", updatesByOperation[PatchOp_Create])
			separator = ", "
		}
		if updatesByOperation[PatchOp_Update] > 0 {
			fmt.Fprintf(&builder, "%s%d updated", separator, updatesByOperation[PatchOp_Update])
			separator = ", "
		}
		if updatesByOperation[PatchOp_Remove] > 0 {
			fmt.Fprintf(&builder, "%s%d removed", separator, updatesByOperation[PatchOp_Remove])
			separator = ", "
		}
		if updatesByOperation[PatchOp_Rename] > 0 {
			fmt.Fprintf(&builder, "%s%d renamed", separator, updatesByOperation[PatchOp_Rename])
		}
	}
	return builder.String()
}

const (
	updateStrategyKey = "update-strategy"
	manifestDigestKey = "manifest-digest"
	AgentBuildInfoKey = "agent-build-info"
)

func ManifestDigestKeyForType(t urn.UpdateType) string {
	return fmt.Sprintf("%s-%s", manifestDigestKey, t)
}

func UpdateStrategyKeyForType(t urn.UpdateType) string {
	return fmt.Sprintf("%s-%s", updateStrategyKey, t)
}

// Returns a hash of the manifest metadata list. This can be used to compare
// manifests between the gateway and agent.
func (m *UpdateManifest) Digest() string {
	if m == nil {
		return ""
	}
	m.Sort()
	hash, _ := blake2b.New256(nil)
	hash.Write([]byte(strconv.Itoa(len(m.GetItems()))))
	for i, entry := range m.GetItems() {
		hash.Write([]byte(strconv.Itoa(i)))

		hash.Write([]byte(strconv.Itoa(len(entry.GetPackage()))))
		hash.Write([]byte(entry.GetPackage()))

		hash.Write([]byte(strconv.Itoa(len(entry.GetPath()))))
		hash.Write([]byte(entry.GetPath()))

		hash.Write([]byte(strconv.Itoa(len(entry.GetDigest()))))
		hash.Write([]byte(entry.GetDigest()))
	}

	return hex.EncodeToString(hash.Sum(nil))
}
