package packagescompiler

import (
	boshblob "github.com/cloudfoundry/bosh-utils/blobstore"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	bpagclient "github.com/bosh-dep-forks/bosh-provisioner/agent/client"
	bpeventlog "github.com/bosh-dep-forks/bosh-provisioner/eventlog"
	bpcpkgsrepo "github.com/bosh-dep-forks/bosh-provisioner/packagescompiler/compiledpackagesrepo"
	bppkgsrepo "github.com/bosh-dep-forks/bosh-provisioner/packagescompiler/packagesrepo"
)

type ConcretePackagesCompilerFactory struct {
	packagesRepo         bppkgsrepo.PackagesRepository
	compiledPackagesRepo bpcpkgsrepo.CompiledPackagesRepository
	blobstore            boshblob.DigestBlobstore

	eventLog bpeventlog.Log
	logger   boshlog.Logger
}

func NewConcretePackagesCompilerFactory(
	packagesRepo bppkgsrepo.PackagesRepository,
	compiledPackagesRepo bpcpkgsrepo.CompiledPackagesRepository,
	blobstore boshblob.DigestBlobstore,
	eventLog bpeventlog.Log,
	logger boshlog.Logger,
) ConcretePackagesCompilerFactory {
	return ConcretePackagesCompilerFactory{
		packagesRepo:         packagesRepo,
		compiledPackagesRepo: compiledPackagesRepo,
		blobstore:            blobstore,

		eventLog: eventLog,
		logger:   logger,
	}
}

func (f ConcretePackagesCompilerFactory) NewCompiler(agentClient bpagclient.Client) PackagesCompiler {
	return NewConcretePackagesCompiler(
		agentClient,
		f.packagesRepo,
		f.compiledPackagesRepo,
		f.blobstore,
		f.eventLog,
		f.logger,
	)
}
