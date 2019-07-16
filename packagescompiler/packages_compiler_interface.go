package packagescompiler

import (
	bprel "github.com/bosh-dep-forks/bosh-provisioner/release"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
)

type CompiledPackageRecord struct {
	SHA1   boshcrypto.MultipleDigest
	BlobID string
}

// PackagesCompiler takes each release package and compiles it.
// Compiled packages are used as:
//   (1) compile dependencies for other packages
//   (2) runtime dependencies for jobs
// todo account for stemcells
type PackagesCompiler interface {
	Compile(bprel.Release) error
	FindCompiledPackage(bprel.Package) (CompiledPackageRecord, error)
}
