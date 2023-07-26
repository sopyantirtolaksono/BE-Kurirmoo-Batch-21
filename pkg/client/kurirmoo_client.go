// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"kurirmoo/pkg/client/cities"
	"kurirmoo/pkg/client/city_by_name"
	"kurirmoo/pkg/client/detail_data_multiplier"
	"kurirmoo/pkg/client/health"
	"kurirmoo/pkg/client/login"
	"kurirmoo/pkg/client/trucks"
	"kurirmoo/pkg/client/update_driver"
)

// Default kurirmoo HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new kurirmoo HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Kurirmoo {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new kurirmoo HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Kurirmoo {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new kurirmoo client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Kurirmoo {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Kurirmoo)
	cli.Transport = transport
	cli.Cities = cities.New(transport, formats)
	cli.CityByName = city_by_name.New(transport, formats)
	cli.DetailDataMultiplier = detail_data_multiplier.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.Login = login.New(transport, formats)
	cli.Trucks = trucks.New(transport, formats)
	cli.UpdateDriver = update_driver.New(transport, formats)
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

// Kurirmoo is a client for kurirmoo
type Kurirmoo struct {
	Cities cities.ClientService

	CityByName city_by_name.ClientService

	DetailDataMultiplier detail_data_multiplier.ClientService

	Health health.ClientService

	Login login.ClientService

	Trucks trucks.ClientService

	UpdateDriver update_driver.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Kurirmoo) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Cities.SetTransport(transport)
	c.CityByName.SetTransport(transport)
	c.DetailDataMultiplier.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.Login.SetTransport(transport)
	c.Trucks.SetTransport(transport)
	c.UpdateDriver.SetTransport(transport)
}
