package provider

import (
	"github.com/mesg-foundation/core/protobuf/core"
)

// Provider is a struct that provides all methods required by any command.
type Provider struct {
	*CoreProvider
	*ServiceProvider
}

// New creates Provider based on given CoreClient.
func New(c core.CoreClient) *Provider {
	return &Provider{
		CoreProvider:    NewCoreProvider(c),
		ServiceProvider: NewServiceProvider(c),
	}
}
