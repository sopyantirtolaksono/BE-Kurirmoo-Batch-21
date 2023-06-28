// Code generated by go-swagger; DO NOT EDIT.

package cities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetAllCitiesParams creates a new GetAllCitiesParams object
// with the default values initialized.
func NewGetAllCitiesParams() GetAllCitiesParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(10)
	)

	return GetAllCitiesParams{
		Limit: &limitDefault,
	}
}

// GetAllCitiesParams contains all the bound params for the get all cities operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAllCities
type GetAllCitiesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*filter with name
	  In: query
	*/
	CityName *string
	/*filter for limit
	  Maximum: 100
	  In: query
	  Default: 10
	*/
	Limit *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAllCitiesParams() beforehand.
func (o *GetAllCitiesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCityName, qhkCityName, _ := qs.GetOK("city_name")
	if err := o.bindCityName(qCityName, qhkCityName, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCityName binds and validates parameter CityName from query.
func (o *GetAllCitiesParams) bindCityName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CityName = &raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetAllCitiesParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAllCitiesParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetAllCitiesParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MaximumInt("limit", "query", *o.Limit, 100, false); err != nil {
		return err
	}

	return nil
}
