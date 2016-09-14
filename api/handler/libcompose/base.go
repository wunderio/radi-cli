package libcompose

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

/**
 * Some usefull Base classes used by other libcompose operations
 * and configurations
 */

// A base libcompose operation with configurations for staying attached
type BaseLibcomposeStayAttachedOperation struct {
	configurations *operation.Configurations
}

// Provide static configurations for the operation
func (base *BaseLibcomposeStayAttachedOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		newConfigurations := &operation.Configurations{}

		newConfigurations.Add(operation.Configuration(&LibcomposeAttachFollowConfiguration{}))

		base.configurations = newConfigurations
	}
	return base.configurations
}

// A base libcompose operation with configurations for project-name, and yml files
type BaseLibcomposeOrchestrateNameFilesOperation struct {
	configurations *operation.Configurations
}

// Provide static configurations for the operation
func (base *BaseLibcomposeOrchestrateNameFilesOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		newConfigurations := &operation.Configurations{}

		newConfigurations.Add(operation.Configuration(&LibcomposeProjectnameConfiguration{}))
		newConfigurations.Add(operation.Configuration(&LibcomposeComposefilesConfiguration{}))
		newConfigurations.Add(operation.Configuration(&LibcomposeContextConfiguration{}))

		newConfigurations.Add(operation.Configuration(&LibcomposeOutputConfiguration{}))
		newConfigurations.Add(operation.Configuration(&LibcomposeErrorConfiguration{}))

		base.configurations = newConfigurations
	}
	return base.configurations
}

// Base Up operation
type BaseLibcomposeOrchestrateUpOperation struct {
	configurations *operation.Configurations
}

// Provide static configurations for the operation
func (base *BaseLibcomposeOrchestrateUpOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		newConfigurations := &operation.Configurations{}

		newConfigurations.Add(operation.Configuration(&LibcomposeOptionsUpConfiguration{}))

		base.configurations = newConfigurations
	}
	return base.configurations
}

// Base Down operation
type BaseLibcomposeOrchestrateDownOperation struct {
	configurations *operation.Configurations
}

// Provide static configurations for the operation
func (base *BaseLibcomposeOrchestrateDownOperation) Configurations() *operation.Configurations {
	if base.configurations == nil {
		newConfigurations := &operation.Configurations{}

		newConfigurations.Add(operation.Configuration(&LibcomposeOptionsDownConfiguration{}))

		base.configurations = newConfigurations
	}
	return base.configurations
}
