// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3 workspace id connectors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 workspace id connectors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRecommendationForWorkspace creates a recommendation that advises cloud resources for the given blueprint

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) CreateRecommendationForWorkspace(params *CreateRecommendationForWorkspaceParams) (*CreateRecommendationForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRecommendationForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRecommendationForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/recommendation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRecommendationForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRecommendationForWorkspaceOK), nil

}

/*
GetAccessConfigsForWorkspace retrives access configs with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetAccessConfigsForWorkspace(params *GetAccessConfigsForWorkspaceParams) (*GetAccessConfigsForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccessConfigsForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccessConfigsForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/accessconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAccessConfigsForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccessConfigsForWorkspaceOK), nil

}

/*
GetDisktypesForWorkspace retrives available disk types

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetDisktypesForWorkspace(params *GetDisktypesForWorkspaceParams) (*GetDisktypesForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDisktypesForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDisktypesForWorkspace",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/connectors/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDisktypesForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDisktypesForWorkspaceOK), nil

}

/*
GetEncryptionKeysForWorkspace retrives encryption keys with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetEncryptionKeysForWorkspace(params *GetEncryptionKeysForWorkspaceParams) (*GetEncryptionKeysForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEncryptionKeysForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEncryptionKeysForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/encryptionkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEncryptionKeysForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEncryptionKeysForWorkspaceOK), nil

}

/*
GetGatewaysCredentialIDForWorkspace retrives gateways with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetGatewaysCredentialIDForWorkspace(params *GetGatewaysCredentialIDForWorkspaceParams) (*GetGatewaysCredentialIDForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGatewaysCredentialIDForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGatewaysCredentialIdForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGatewaysCredentialIDForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGatewaysCredentialIDForWorkspaceOK), nil

}

/*
GetIPPoolsCredentialIDForWorkspace retrives ip pools with properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetIPPoolsCredentialIDForWorkspace(params *GetIPPoolsCredentialIDForWorkspaceParams) (*GetIPPoolsCredentialIDForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIPPoolsCredentialIDForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIpPoolsCredentialIdForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/ippools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetIPPoolsCredentialIDForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIPPoolsCredentialIDForWorkspaceOK), nil

}

/*
GetPlatformNetworksForWorkspace retrives network properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformNetworksForWorkspace(params *GetPlatformNetworksForWorkspaceParams) (*GetPlatformNetworksForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformNetworksForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformNetworksForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformNetworksForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformNetworksForWorkspaceOK), nil

}

/*
GetPlatformSShKeysForWorkspace retrives sshkeys properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformSShKeysForWorkspace(params *GetPlatformSShKeysForWorkspaceParams) (*GetPlatformSShKeysForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformSShKeysForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformSShKeysForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformSShKeysForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSShKeysForWorkspaceOK), nil

}

/*
GetPlatformSecurityGroupsForWorkspace retrives securitygroups properties

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetPlatformSecurityGroupsForWorkspace(params *GetPlatformSecurityGroupsForWorkspaceParams) (*GetPlatformSecurityGroupsForWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlatformSecurityGroupsForWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlatformSecurityGroupsForWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/securitygroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPlatformSecurityGroupsForWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlatformSecurityGroupsForWorkspaceOK), nil

}

/*
GetRegionsByCredentialAndWorkspace retrives regions by type

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetRegionsByCredentialAndWorkspace(params *GetRegionsByCredentialAndWorkspaceParams) (*GetRegionsByCredentialAndWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsByCredentialAndWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegionsByCredentialAndWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegionsByCredentialAndWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegionsByCredentialAndWorkspaceOK), nil

}

/*
GetVMTypesByCredentialAndWorkspace retrives vmtype properties by credential

Each cloud provider has it's own specific resources like instance types and disk types. These endpoints are collecting them.
*/
func (a *Client) GetVMTypesByCredentialAndWorkspace(params *GetVMTypesByCredentialAndWorkspaceParams) (*GetVMTypesByCredentialAndWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMTypesByCredentialAndWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVmTypesByCredentialAndWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/connectors/vmtypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVMTypesByCredentialAndWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVMTypesByCredentialAndWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
