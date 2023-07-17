// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllCustomer(params *GetAllCustomerParams, opts ...ClientOption) (*GetAllCustomerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAllCustomer gets list of customer

Return all list of customer
*/
func (a *Client) GetAllCustomer(params *GetAllCustomerParams, opts ...ClientOption) (*GetAllCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllCustomer",
		Method:             "GET",
		PathPattern:        "/api/v1/customers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllCustomerReader{formats: a.formats},
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
	success, ok := result.(*GetAllCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAllCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
