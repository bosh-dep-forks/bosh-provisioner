package packagesrepo

import (
	bprel "github.com/bosh-dep-forks/bosh-provisioner/release"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
)

type PackageRecord struct {
	BlobID string
	SHA1   boshcrypto.MultipleDigest
}

// PackagesRepository maintains list of package source code as blobs.
type PackagesRepository interface {
	Find(bprel.Package) (PackageRecord, bool, error)
	Save(bprel.Package, PackageRecord) error
}
