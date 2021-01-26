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

package current_identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// DetailMfaOKCode is the HTTP code returned for type DetailMfaOK
const DetailMfaOKCode int = 200

/*DetailMfaOK The details of an MFA enrollment

swagger:response detailMfaOK
*/
type DetailMfaOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailMfaEnvelope `json:"body,omitempty"`
}

// NewDetailMfaOK creates DetailMfaOK with default headers values
func NewDetailMfaOK() *DetailMfaOK {

	return &DetailMfaOK{}
}

// WithPayload adds the payload to the detail mfa o k response
func (o *DetailMfaOK) WithPayload(payload *rest_model.DetailMfaEnvelope) *DetailMfaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail mfa o k response
func (o *DetailMfaOK) SetPayload(payload *rest_model.DetailMfaEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailMfaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailMfaUnauthorizedCode is the HTTP code returned for type DetailMfaUnauthorized
const DetailMfaUnauthorizedCode int = 401

/*DetailMfaUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailMfaUnauthorized
*/
type DetailMfaUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailMfaUnauthorized creates DetailMfaUnauthorized with default headers values
func NewDetailMfaUnauthorized() *DetailMfaUnauthorized {

	return &DetailMfaUnauthorized{}
}

// WithPayload adds the payload to the detail mfa unauthorized response
func (o *DetailMfaUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailMfaUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail mfa unauthorized response
func (o *DetailMfaUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailMfaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailMfaNotFoundCode is the HTTP code returned for type DetailMfaNotFound
const DetailMfaNotFoundCode int = 404

/*DetailMfaNotFound The requested resource does not exist

swagger:response detailMfaNotFound
*/
type DetailMfaNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailMfaNotFound creates DetailMfaNotFound with default headers values
func NewDetailMfaNotFound() *DetailMfaNotFound {

	return &DetailMfaNotFound{}
}

// WithPayload adds the payload to the detail mfa not found response
func (o *DetailMfaNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailMfaNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail mfa not found response
func (o *DetailMfaNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailMfaNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
