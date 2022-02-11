// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginMount PluginMount PluginMount PluginMount PluginMount plugin mount
//
// swagger:model PluginMount
type PluginMount struct {

	// description
	// Required: true
	Description *string `json:"Description"`

	// destination
	// Required: true
	Destination *string `json:"Destination"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// options
	// Required: true
	Options []string `json:"Options"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// source
	// Required: true
	Source *string `json:"Source"`

	// type
	// Required: true
	Type *string `json:"Type"`
}

// Validate validates this plugin mount
func (m *PluginMount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginMount) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("Destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("Options", "body", m.Options); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateSettable(formats strfmt.Registry) error {

	if err := validate.Required("Settable", "body", m.Settable); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("Source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateType(formats strfmt.Registry) error {

	if err := validate.Required("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin mount based on context it is used
func (m *PluginMount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginMount) UnmarshalBinary(b []byte) error {
	var res PluginMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
