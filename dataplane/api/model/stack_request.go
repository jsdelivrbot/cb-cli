// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackRequest stack request
// swagger:model StackRequest
type StackRequest struct {

	// specific version of ambari
	AmbariVersion string `json:"ambariVersion,omitempty"`

	// stack related application tags
	ApplicationTags map[string]string `json:"applicationTags,omitempty"`

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// type of cloud provider
	// Read Only: true
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// using the cluster name to create subdomain
	ClusterNameAsSubdomain *bool `json:"clusterNameAsSubdomain,omitempty"`

	// cluster request object on stack
	ClusterRequest *ClusterRequest `json:"clusterRequest,omitempty"`

	// Shared service request
	ClusterToAttach int64 `json:"clusterToAttach,omitempty"`

	// stack related credential
	Credential *CredentialRequest `json:"credential,omitempty"`

	// credential resource id for the stack
	CredentialID int64 `json:"credentialId,omitempty"`

	// credential resource name for the stack
	CredentialName string `json:"credentialName,omitempty"`

	// source credential object for cloning
	CredentialSource *CredentialSourceRequest `json:"credentialSource,omitempty"`

	// custom domain name for the nodes in the stack
	CustomDomain string `json:"customDomain,omitempty"`

	// custom hostname for nodes in the stack
	CustomHostname string `json:"customHostname,omitempty"`

	// AmbariKerberosDescriptor parameters as a json
	CustomInputs map[string]interface{} `json:"customInputs,omitempty"`

	// stack related default tags
	DefaultTags map[string]string `json:"defaultTags,omitempty"`

	// environment which the stack is assigned to
	Environment string `json:"environment,omitempty"`

	// failure policy in case of failures
	FailurePolicy *FailurePolicyRequest `json:"failurePolicy,omitempty"`

	// id of the related flex subscription
	FlexID int64 `json:"flexId,omitempty"`

	// port of the gateway secured proxy
	GatewayPort int32 `json:"gatewayPort,omitempty"`

	// specific version of HDP
	HdpVersion string `json:"hdpVersion,omitempty"`

	// using the hostgroup names to create hostnames
	HostgroupNameAsHostname *bool `json:"hostgroupNameAsHostname,omitempty"`

	// custom image catalog URL
	ImageCatalog string `json:"imageCatalog,omitempty"`

	// virtual machine image id from ImageCatalog, machines of the cluster will be started from this image
	ImageID string `json:"imageId,omitempty"`

	// collection of instance groupst
	// Required: true
	InstanceGroups []*InstanceGroups `json:"instanceGroups"`

	// name of the stack
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// stack related network
	Network *NetworkRequest `json:"network,omitempty"`

	// network resource id for the stack
	NetworkID int64 `json:"networkId,omitempty"`

	// action on failure
	// Enum: [ROLLBACK DO_NOTHING]
	OnFailureAction string `json:"onFailureAction,omitempty"`

	// the details of the container orchestrator api to use
	Orchestrator *OrchestratorRequest `json:"orchestrator,omitempty"`

	// os type of the image, this property is only considered when no specific image id is provided
	Os string `json:"os,omitempty"`

	// additional cloud specific parameters for stack
	Parameters map[string]string `json:"parameters,omitempty"`

	// cloud provider api variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// region of the stack
	Region string `json:"region,omitempty"`

	// stack related authentication
	StackAuthentication *StackAuthentication `json:"stackAuthentication,omitempty"`

	// stack related userdefined tags
	UserDefinedTags map[string]string `json:"userDefinedTags,omitempty"`
}

// Validate validates this stack request
func (m *StackRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailurePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnFailureAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrchestrator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackRequest) validateClusterRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterRequest) { // not required
		return nil
	}

	if m.ClusterRequest != nil {
		if err := m.ClusterRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterRequest")
			}
			return err
		}
	}

	return nil
}

func (m *StackRequest) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *StackRequest) validateCredentialSource(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialSource) { // not required
		return nil
	}

	if m.CredentialSource != nil {
		if err := m.CredentialSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialSource")
			}
			return err
		}
	}

	return nil
}

func (m *StackRequest) validateFailurePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.FailurePolicy) { // not required
		return nil
	}

	if m.FailurePolicy != nil {
		if err := m.FailurePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failurePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *StackRequest) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *StackRequest) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

var stackRequestTypeOnFailureActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ROLLBACK","DO_NOTHING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackRequestTypeOnFailureActionPropEnum = append(stackRequestTypeOnFailureActionPropEnum, v)
	}
}

const (

	// StackRequestOnFailureActionROLLBACK captures enum value "ROLLBACK"
	StackRequestOnFailureActionROLLBACK string = "ROLLBACK"

	// StackRequestOnFailureActionDONOTHING captures enum value "DO_NOTHING"
	StackRequestOnFailureActionDONOTHING string = "DO_NOTHING"
)

// prop value enum
func (m *StackRequest) validateOnFailureActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackRequestTypeOnFailureActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackRequest) validateOnFailureAction(formats strfmt.Registry) error {

	if swag.IsZero(m.OnFailureAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnFailureActionEnum("onFailureAction", "body", m.OnFailureAction); err != nil {
		return err
	}

	return nil
}

func (m *StackRequest) validateOrchestrator(formats strfmt.Registry) error {

	if swag.IsZero(m.Orchestrator) { // not required
		return nil
	}

	if m.Orchestrator != nil {
		if err := m.Orchestrator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orchestrator")
			}
			return err
		}
	}

	return nil
}

func (m *StackRequest) validateStackAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.StackAuthentication) { // not required
		return nil
	}

	if m.StackAuthentication != nil {
		if err := m.StackAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackAuthentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackRequest) UnmarshalBinary(b []byte) error {
	var res StackRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
