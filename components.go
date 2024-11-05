package types

import "time"

type ComponentType string

func (t ComponentType) String() string {
	return string(t)
}

const (
	ComponentTypeVirtualMachine          ComponentType = "vm"
	ComponentTypeDatabase                ComponentType = "db"
	ComponentTypeApplicationLoadBalancer ComponentType = "alb"
	ComponentTypeNetworkLoadBalancer     ComponentType = "nlb"
	ComponentTypeStorage                 ComponentType = "storage"
	ComponentTypeNetwork                 ComponentType = "network"
	ComponentTypeSecurityGroup           ComponentType = "security-group"
)

var (
	ComponentTypes = []ComponentType{
		ComponentTypeVirtualMachine,
		ComponentTypeDatabase,
		ComponentTypeApplicationLoadBalancer,
		ComponentTypeNetworkLoadBalancer,
		ComponentTypeStorage,
		ComponentTypeNetwork,
		ComponentTypeSecurityGroup,
	}
)

type ComponentProvider string

func (t ComponentProvider) String() string {
	return string(t)
}

const (
	ComponentProviderAWS       ComponentProvider = "aws"
	ComponentProviderGCP       ComponentProvider = "gcp"
	ComponentProviderAzure     ComponentProvider = "azure"
	ComponentProviderOnPremise ComponentProvider = "onpremise"
)

var (
	ComponentProviders = []ComponentProvider{
		ComponentProviderAWS,
		ComponentProviderGCP,
		ComponentProviderAzure,
		ComponentProviderOnPremise,
	}
)

type Component struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        ComponentType     `json:"type"`
	Provider    ComponentProvider `json:"provider"`
	Status      Status            `json:"status"`
	Created     time.Time         `json:"created"`
	Updated     time.Time         `json:"updated"`
	Tags        []string          `json:"tags,omitempty"`
}

type ComponentCreate struct {
	Name        string            `json:"name" validate:"required,min=1,max=32"`
	Description string            `json:"description" validate:"required,min=1,max=255"`
	Type        ComponentType     `json:"type" validate:"required"`
	Provider    ComponentProvider `json:"provider" validate:"required"`
}
