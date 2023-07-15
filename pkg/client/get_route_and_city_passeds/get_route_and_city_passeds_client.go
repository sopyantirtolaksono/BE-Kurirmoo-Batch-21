// Code generated by go-swagger; DO NOT EDIT.

package get_route_and_city_passeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new get route and city passeds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for get route and city passeds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetRouteAndCityPasseds(params *GetRouteAndCityPassedsParams, opts ...ClientOption) (*GetRouteAndCityPassedsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetRouteAndCityPasseds gets route and city passeds

Return route and city passeds
*/
func (a *Client) GetRouteAndCityPasseds(params *GetRouteAndCityPassedsParams, opts ...ClientOption) (*GetRouteAndCityPassedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRouteAndCityPassedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRouteAndCityPasseds",
		Method:             "GET",
		PathPattern:        "/api/v1/routes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRouteAndCityPassedsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRouteAndCityPassedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRouteAndCityPassedsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
