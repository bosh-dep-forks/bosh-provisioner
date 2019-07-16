package compiledpackagesrepo

import (
	bprel "github.com/bosh-dep-forks/bosh-provisioner/release"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
)

type CompiledPackageRecord struct {
	BlobID string
	SHA1   boshcrypto.MultipleDigest
}

// CompiledPackagesRepository maintains list of compiled packages as blobs
// todo account for stemcell
type CompiledPackagesRepository interface {
	Find(bprel.Package) (CompiledPackageRecord, bool, error)
	Save(bprel.Package, CompiledPackageRecord) error
}
