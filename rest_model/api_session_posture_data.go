// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APISessionPostureData api session posture data
//
// swagger:model apiSessionPostureData
type APISessionPostureData struct {

	// endpoint state
	EndpointState *PostureDataEndpointState `json:"endpointState,omitempty"`

	// mfa
	// Required: true
	Mfa *PostureDataMfa `json:"mfa"`

	// sdk version
	SdkVersion string `json:"sdkVersion,omitempty"`
}

// Validate validates this api session posture data
func (m *APISessionPostureData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMfa(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISessionPostureData) validateEndpointState(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointState) { // not required
		return nil
	}

	if m.EndpointState != nil {
		if err := m.EndpointState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpointState")
			}
			return err
		}
	}

	return nil
}

func (m *APISessionPostureData) validateMfa(formats strfmt.Registry) error {

	if err := validate.Required("mfa", "body", m.Mfa); err != nil {
		return err
	}

	if m.Mfa != nil {
		if err := m.Mfa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mfa")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this api session posture data based on the context it is used
func (m *APISessionPostureData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpointState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMfa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APISessionPostureData) contextValidateEndpointState(ctx context.Context, formats strfmt.Registry) error {

	if m.EndpointState != nil {
		if err := m.EndpointState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpointState")
			}
			return err
		}
	}

	return nil
}

func (m *APISessionPostureData) contextValidateMfa(ctx context.Context, formats strfmt.Registry) error {

	if m.Mfa != nil {
		if err := m.Mfa.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mfa")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APISessionPostureData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APISessionPostureData) UnmarshalBinary(b []byte) error {
	var res APISessionPostureData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
