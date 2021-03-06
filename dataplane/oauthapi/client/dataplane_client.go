// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/oidc"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/roles"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/status"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/strategies"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/strategytypes"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/users"
)

// Default dataplane HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:10080"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new dataplane HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Dataplane {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new dataplane HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Dataplane {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new dataplane client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Dataplane {
	cli := new(Dataplane)
	cli.Transport = transport

	cli.Oidc = oidc.New(transport, formats)

	cli.Roles = roles.New(transport, formats)

	cli.Status = status.New(transport, formats)

	cli.Strategies = strategies.New(transport, formats)

	cli.Strategytypes = strategytypes.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Dataplane is a client for dataplane
type Dataplane struct {
	Oidc *oidc.Client

	Roles *roles.Client

	Status *status.Client

	Strategies *strategies.Client

	Strategytypes *strategytypes.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Dataplane) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Oidc.SetTransport(transport)

	c.Roles.SetTransport(transport)

	c.Status.SetTransport(transport)

	c.Strategies.SetTransport(transport)

	c.Strategytypes.SetTransport(transport)

	c.Users.SetTransport(transport)

}
