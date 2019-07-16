package provisioner

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	bpdep "github.com/bosh-dep-forks/bosh-provisioner/deployment"
	bpeventlog "github.com/bosh-dep-forks/bosh-provisioner/eventlog"
	bpinstance "github.com/bosh-dep-forks/bosh-provisioner/instance"
	bpvm "github.com/bosh-dep-forks/bosh-provisioner/vm"
)

type SingleVMProvisionerFactory struct {
	deploymentReaderFactory     bpdep.ReaderFactory
	deploymentProvisionerConfig DeploymentProvisionerConfig

	vmProvisioner       bpvm.Provisioner
	releaseCompiler     ReleaseCompiler
	instanceProvisioner bpinstance.Provisioner

	eventLog bpeventlog.Log
	logger   boshlog.Logger
}

func NewSingleVMProvisionerFactory(
	deploymentReaderFactory bpdep.ReaderFactory,
	deploymentProvisionerConfig DeploymentProvisionerConfig,
	vmProvisioner bpvm.Provisioner,
	releaseCompiler ReleaseCompiler,
	instanceProvisioner bpinstance.Provisioner,
	eventLog bpeventlog.Log,
	logger boshlog.Logger,
) SingleVMProvisionerFactory {
	return SingleVMProvisionerFactory{
		deploymentReaderFactory:     deploymentReaderFactory,
		deploymentProvisionerConfig: deploymentProvisionerConfig,

		vmProvisioner:       vmProvisioner,
		releaseCompiler:     releaseCompiler,
		instanceProvisioner: instanceProvisioner,

		eventLog: eventLog,
		logger:   logger,
	}
}

func (f SingleVMProvisionerFactory) NewSingleVMProvisioner() DeploymentProvisioner {
	var prov DeploymentProvisioner

	if len(f.deploymentProvisionerConfig.ManifestPath) > 0 {
		prov = NewSingleConfiguredVMProvisioner(
			f.deploymentProvisionerConfig.ManifestPath,
			f.deploymentReaderFactory,
			f.vmProvisioner,
			f.releaseCompiler,
			f.instanceProvisioner,
			f.eventLog,
			f.logger,
		)
	} else {
		prov = NewSingleNonConfiguredVMProvisioner(
			f.vmProvisioner,
			f.eventLog,
			f.logger,
		)
	}

	return prov
}
