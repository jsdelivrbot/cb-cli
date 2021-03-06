// Code generated by go-swagger; DO NOT EDIT.

package v1accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1accountpreferences API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1accountpreferences API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAccountPreferencesEndpoint retrieves account preferences for admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) GetAccountPreferencesEndpoint(params *GetAccountPreferencesEndpointParams) (*GetAccountPreferencesEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountPreferencesEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccountPreferencesEndpoint",
		Method:             "GET",
		PathPattern:        "/v1/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccountPreferencesEndpointReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountPreferencesEndpointOK), nil

}

/*
IsPlatformSelectionDisabled is platform selection disabled

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) IsPlatformSelectionDisabled(params *IsPlatformSelectionDisabledParams) (*IsPlatformSelectionDisabledOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsPlatformSelectionDisabledParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "isPlatformSelectionDisabled",
		Method:             "GET",
		PathPattern:        "/v1/accountpreferences/isplatformselectiondisabled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsPlatformSelectionDisabledReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IsPlatformSelectionDisabledOK), nil

}

/*
PlatformEnablement is platform selection enabled

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) PlatformEnablement(params *PlatformEnablementParams) (*PlatformEnablementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPlatformEnablementParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "platformEnablement",
		Method:             "GET",
		PathPattern:        "/v1/accountpreferences/platformenabled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PlatformEnablementReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PlatformEnablementOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
