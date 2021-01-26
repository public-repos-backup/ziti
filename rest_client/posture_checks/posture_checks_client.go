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

package posture_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new posture checks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for posture checks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePostureCheck(params *CreatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePostureCheckCreated, error)

	CreatePostureResponse(params *CreatePostureResponseParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePostureResponseCreated, error)

	CreatePostureResponseBulk(params *CreatePostureResponseBulkParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePostureResponseBulkOK, error)

	DeletePostureCheck(params *DeletePostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePostureCheckOK, error)

	DetailPostureCheck(params *DetailPostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*DetailPostureCheckOK, error)

	DetailPostureCheckType(params *DetailPostureCheckTypeParams, authInfo runtime.ClientAuthInfoWriter) (*DetailPostureCheckTypeOK, error)

	ListPostureCheckTypes(params *ListPostureCheckTypesParams, authInfo runtime.ClientAuthInfoWriter) (*ListPostureCheckTypesOK, error)

	ListPostureChecks(params *ListPostureChecksParams, authInfo runtime.ClientAuthInfoWriter) (*ListPostureChecksOK, error)

	PatchPostureCheck(params *PatchPostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPostureCheckOK, error)

	UpdatePostureCheck(params *UpdatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePostureCheckOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePostureCheck creates a posture checks

  Creates a Posture Checks
*/
func (a *Client) CreatePostureCheck(params *CreatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePostureCheckCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostureCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPostureCheck",
		Method:             "POST",
		PathPattern:        "/posture-checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePostureCheckCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreatePostureResponse submits a posture response to a posture query

  Submits posture responses
*/
func (a *Client) CreatePostureResponse(params *CreatePostureResponseParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePostureResponseCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostureResponseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPostureResponse",
		Method:             "POST",
		PathPattern:        "/posture-response",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePostureResponseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePostureResponseCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPostureResponse: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreatePostureResponseBulk submits multiple posture responses

  Submits posture responses
*/
func (a *Client) CreatePostureResponseBulk(params *CreatePostureResponseBulkParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePostureResponseBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePostureResponseBulkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPostureResponseBulk",
		Method:             "POST",
		PathPattern:        "/posture-response-bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePostureResponseBulkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePostureResponseBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPostureResponseBulk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePostureCheck deletes an posture checks

  Deletes and Posture Checks by id
*/
func (a *Client) DeletePostureCheck(params *DeletePostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePostureCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePostureCheck",
		Method:             "DELETE",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DetailPostureCheck retrieves a single posture checks

  Retrieves a single Posture Checks by id
*/
func (a *Client) DetailPostureCheck(params *DetailPostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*DetailPostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailPostureCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detailPostureCheck",
		Method:             "GET",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailPostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetailPostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailPostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DetailPostureCheckType retrieves a single posture check type

  Retrieves a single posture check type by id
*/
func (a *Client) DetailPostureCheckType(params *DetailPostureCheckTypeParams, authInfo runtime.ClientAuthInfoWriter) (*DetailPostureCheckTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetailPostureCheckTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detailPostureCheckType",
		Method:             "GET",
		PathPattern:        "/posture-check-types/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetailPostureCheckTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DetailPostureCheckTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for detailPostureCheckType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPostureCheckTypes lists a subset of posture check types

  Retrieves a list of posture check types

*/
func (a *Client) ListPostureCheckTypes(params *ListPostureCheckTypesParams, authInfo runtime.ClientAuthInfoWriter) (*ListPostureCheckTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPostureCheckTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPostureCheckTypes",
		Method:             "GET",
		PathPattern:        "/posture-check-types",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPostureCheckTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPostureCheckTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPostureCheckTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPostureChecks lists a subset of posture checks

  Retrieves a list of posture checks

*/
func (a *Client) ListPostureChecks(params *ListPostureChecksParams, authInfo runtime.ClientAuthInfoWriter) (*ListPostureChecksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPostureChecksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPostureChecks",
		Method:             "GET",
		PathPattern:        "/posture-checks",
		ProducesMediaTypes: []string{"application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPostureChecksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPostureChecksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPostureChecks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchPostureCheck updates the supplied fields on a posture checks

  Update only the supplied fields on a Posture Checks by id
*/
func (a *Client) PatchPostureCheck(params *PatchPostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPostureCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchPostureCheck",
		Method:             "PATCH",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchPostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchPostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePostureCheck updates all fields on a posture checks

  Update all fields on a Posture Checks by id
*/
func (a *Client) UpdatePostureCheck(params *UpdatePostureCheckParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePostureCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePostureCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePostureCheck",
		Method:             "PUT",
		PathPattern:        "/posture-checks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePostureCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePostureCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePostureCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
