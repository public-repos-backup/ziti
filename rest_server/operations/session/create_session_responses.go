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

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// CreateSessionCreatedCode is the HTTP code returned for type CreateSessionCreated
const CreateSessionCreatedCode int = 201

/*CreateSessionCreated The create request was successful and the resource has been added at the following location

swagger:response createSessionCreated
*/
type CreateSessionCreated struct {

	/*
	  In: Body
	*/
	Payload *rest_model.SessionCreateEnvelope `json:"body,omitempty"`
}

// NewCreateSessionCreated creates CreateSessionCreated with default headers values
func NewCreateSessionCreated() *CreateSessionCreated {

	return &CreateSessionCreated{}
}

// WithPayload adds the payload to the create session created response
func (o *CreateSessionCreated) WithPayload(payload *rest_model.SessionCreateEnvelope) *CreateSessionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create session created response
func (o *CreateSessionCreated) SetPayload(payload *rest_model.SessionCreateEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSessionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSessionBadRequestCode is the HTTP code returned for type CreateSessionBadRequest
const CreateSessionBadRequestCode int = 400

/*CreateSessionBadRequest The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information

swagger:response createSessionBadRequest
*/
type CreateSessionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateSessionBadRequest creates CreateSessionBadRequest with default headers values
func NewCreateSessionBadRequest() *CreateSessionBadRequest {

	return &CreateSessionBadRequest{}
}

// WithPayload adds the payload to the create session bad request response
func (o *CreateSessionBadRequest) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateSessionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create session bad request response
func (o *CreateSessionBadRequest) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSessionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSessionUnauthorizedCode is the HTTP code returned for type CreateSessionUnauthorized
const CreateSessionUnauthorizedCode int = 401

/*CreateSessionUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response createSessionUnauthorized
*/
type CreateSessionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewCreateSessionUnauthorized creates CreateSessionUnauthorized with default headers values
func NewCreateSessionUnauthorized() *CreateSessionUnauthorized {

	return &CreateSessionUnauthorized{}
}

// WithPayload adds the payload to the create session unauthorized response
func (o *CreateSessionUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *CreateSessionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create session unauthorized response
func (o *CreateSessionUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSessionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
