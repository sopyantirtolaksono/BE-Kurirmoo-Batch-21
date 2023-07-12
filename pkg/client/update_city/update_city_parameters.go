// Code generated by go-swagger; DO NOT EDIT.

package update_city

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUpdateCityParams creates a new UpdateCityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCityParams() *UpdateCityParams {
	return &UpdateCityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCityParamsWithTimeout creates a new UpdateCityParams object
// with the ability to set a timeout on a request.
func NewUpdateCityParamsWithTimeout(timeout time.Duration) *UpdateCityParams {
	return &UpdateCityParams{
		timeout: timeout,
	}
}

// NewUpdateCityParamsWithContext creates a new UpdateCityParams object
// with the ability to set a context for a request.
func NewUpdateCityParamsWithContext(ctx context.Context) *UpdateCityParams {
	return &UpdateCityParams{
		Context: ctx,
	}
}

// NewUpdateCityParamsWithHTTPClient creates a new UpdateCityParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCityParamsWithHTTPClient(client *http.Client) *UpdateCityParams {
	return &UpdateCityParams{
		HTTPClient: client,
	}
}

/*
UpdateCityParams contains all the parameters to send to the API endpoint

	for the update city operation.

	Typically these are written to a http.Request.
*/
type UpdateCityParams struct {

	// Code.
	Code string

	// ID.
	ID int64

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update city params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCityParams) WithDefaults() *UpdateCityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update city params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update city params
func (o *UpdateCityParams) WithTimeout(timeout time.Duration) *UpdateCityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update city params
func (o *UpdateCityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update city params
func (o *UpdateCityParams) WithContext(ctx context.Context) *UpdateCityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update city params
func (o *UpdateCityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update city params
func (o *UpdateCityParams) WithHTTPClient(client *http.Client) *UpdateCityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update city params
func (o *UpdateCityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the update city params
func (o *UpdateCityParams) WithCode(code string) *UpdateCityParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the update city params
func (o *UpdateCityParams) SetCode(code string) {
	o.Code = code
}

// WithID adds the id to the update city params
func (o *UpdateCityParams) WithID(id int64) *UpdateCityParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update city params
func (o *UpdateCityParams) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the update city params
func (o *UpdateCityParams) WithName(name string) *UpdateCityParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update city params
func (o *UpdateCityParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param code
	frCode := o.Code
	fCode := frCode
	if fCode != "" {
		if err := r.SetFormParam("code", fCode); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
