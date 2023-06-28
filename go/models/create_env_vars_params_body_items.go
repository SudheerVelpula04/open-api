// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateEnvVarsParamsBodyItems create env vars params body items
//
// swagger:model createEnvVarsParamsBodyItems
type CreateEnvVarsParamsBodyItems struct {

	// Secret values are only readable by code running on Netlify’s systems.  With secrets, only the local development context values are readable from the UI, API, and CLI. By default, environment variable values are not secret. (Enterprise plans only)
	IsSecret bool `json:"is_secret,omitempty"`

	// The existing or new name of the key, if you wish to rename it (case-sensitive)
	Key string `json:"key,omitempty"`

	// The scopes that this environment variable is set to (Pro plans and above)
	Scopes []string `json:"scopes"`

	// values
	Values []*EnvVarValue `json:"values"`
}

// Validate validates this create env vars params body items
func (m *CreateEnvVarsParamsBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createEnvVarsParamsBodyItemsScopesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["builds","functions","runtime","post-processing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createEnvVarsParamsBodyItemsScopesItemsEnum = append(createEnvVarsParamsBodyItemsScopesItemsEnum, v)
	}
}

func (m *CreateEnvVarsParamsBodyItems) validateScopesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createEnvVarsParamsBodyItemsScopesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateEnvVarsParamsBodyItems) validateScopes(formats strfmt.Registry) error {

	if swag.IsZero(m.Scopes) { // not required
		return nil
	}

	for i := 0; i < len(m.Scopes); i++ {

		// value enum
		if err := m.validateScopesItemsEnum("scopes"+"."+strconv.Itoa(i), "body", m.Scopes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateEnvVarsParamsBodyItems) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateEnvVarsParamsBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEnvVarsParamsBodyItems) UnmarshalBinary(b []byte) error {
	var res CreateEnvVarsParamsBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
