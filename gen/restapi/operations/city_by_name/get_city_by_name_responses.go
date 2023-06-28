// Code generated by go-swagger; DO NOT EDIT.

package city_by_name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"kurirmoo/gen/models"
)

// GetCityByNameOKCode is the HTTP code returned for type GetCityByNameOK
const GetCityByNameOKCode int = 200

/*
GetCityByNameOK A JSON array of city's name

swagger:response getCityByNameOK
*/
type GetCityByNameOK struct {

	/*
	  In: Body
	*/
	Payload *GetCityByNameOKBody `json:"body,omitempty"`
}

// NewGetCityByNameOK creates GetCityByNameOK with default headers values
func NewGetCityByNameOK() *GetCityByNameOK {

	return &GetCityByNameOK{}
}

// WithPayload adds the payload to the get city by name o k response
func (o *GetCityByNameOK) WithPayload(payload *GetCityByNameOKBody) *GetCityByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get city by name o k response
func (o *GetCityByNameOK) SetPayload(payload *GetCityByNameOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCityByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCityByNameBadRequestCode is the HTTP code returned for type GetCityByNameBadRequest
const GetCityByNameBadRequestCode int = 400

/*
GetCityByNameBadRequest Bad Request

swagger:response getCityByNameBadRequest
*/
type GetCityByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCityByNameBadRequest creates GetCityByNameBadRequest with default headers values
func NewGetCityByNameBadRequest() *GetCityByNameBadRequest {

	return &GetCityByNameBadRequest{}
}

// WithPayload adds the payload to the get city by name bad request response
func (o *GetCityByNameBadRequest) WithPayload(payload *models.Error) *GetCityByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get city by name bad request response
func (o *GetCityByNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCityByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetCityByNameDefault error

swagger:response getCityByNameDefault
*/
type GetCityByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCityByNameDefault creates GetCityByNameDefault with default headers values
func NewGetCityByNameDefault(code int) *GetCityByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCityByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get city by name default response
func (o *GetCityByNameDefault) WithStatusCode(code int) *GetCityByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get city by name default response
func (o *GetCityByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get city by name default response
func (o *GetCityByNameDefault) WithPayload(payload *models.Error) *GetCityByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get city by name default response
func (o *GetCityByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCityByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
