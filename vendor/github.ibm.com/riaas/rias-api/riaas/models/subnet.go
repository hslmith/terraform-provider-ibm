// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subnet Subnet
// swagger:model subnet
type Subnet struct {

	// The number of available IPv4 addresses in this subnet
	AvailableIPV4AddressCount int64 `json:"available_ipv4_address_count,omitempty"`

	// The date and time that the subnet was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this subnet
	Crn string `json:"crn,omitempty"`

	// generation
	Generation Generation `json:"generation,omitempty"`

	// The URL for this subnet
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this subnet
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The IP version(s) supported by this subnet
	// Enum: [ipv4 ipv6 both]
	IPVersion string `json:"ip_version,omitempty"`

	// The IPv4 range of the subnet, expressed in CIDR format
	IPV4CidrBlock string `json:"ipv4_cidr_block,omitempty"`

	// The IPv6 range of the subnet, expressed in CIDR format
	IPV6CidrBlock string `json:"ipv6_cidr_block,omitempty"`

	// The user-defined name for this subnet
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// network acl
	NetworkACL *ResourceReference `json:"network_acl,omitempty"`

	// public gateway
	PublicGateway *ResourceReference `json:"public_gateway,omitempty"`

	// resource group
	ResourceGroup *ResourceReference `json:"resource_group,omitempty"`

	// The status of the subnet
	// Enum: [pending available]
	Status string `json:"status,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// The total number of IPv4 addresses in this subnet
	TotalIPV4AddressCount int64 `json:"total_ipv4_address_count,omitempty"`

	// vpc
	Vpc *ResourceReference `json:"vpc,omitempty"`

	// zone
	Zone *SubnetZone `json:"zone,omitempty"`
}

// Validate validates this subnet
func (m *Subnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subnet) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateGeneration(formats strfmt.Registry) error {

	if swag.IsZero(m.Generation) { // not required
		return nil
	}

	if err := m.Generation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("generation")
		}
		return err
	}

	return nil
}

func (m *Subnet) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var subnetTypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subnetTypeIPVersionPropEnum = append(subnetTypeIPVersionPropEnum, v)
	}
}

const (

	// SubnetIPVersionIPV4 captures enum value "ipv4"
	SubnetIPVersionIPV4 string = "ipv4"

	// SubnetIPVersionIPV6 captures enum value "ipv6"
	SubnetIPVersionIPV6 string = "ipv6"

	// SubnetIPVersionBoth captures enum value "both"
	SubnetIPVersionBoth string = "both"
)

// prop value enum
func (m *Subnet) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subnetTypeIPVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subnet) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPVersionEnum("ip_version", "body", m.IPVersion); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateNetworkACL(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkACL) { // not required
		return nil
	}

	if m.NetworkACL != nil {
		if err := m.NetworkACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_acl")
			}
			return err
		}
	}

	return nil
}

func (m *Subnet) validatePublicGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicGateway) { // not required
		return nil
	}

	if m.PublicGateway != nil {
		if err := m.PublicGateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_gateway")
			}
			return err
		}
	}

	return nil
}

func (m *Subnet) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

var subnetTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","available"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subnetTypeStatusPropEnum = append(subnetTypeStatusPropEnum, v)
	}
}

const (

	// SubnetStatusPending captures enum value "pending"
	SubnetStatusPending string = "pending"

	// SubnetStatusAvailable captures enum value "available"
	SubnetStatusAvailable string = "available"
)

// prop value enum
func (m *Subnet) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subnetTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subnet) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Subnet) validateVpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

func (m *Subnet) validateZone(formats strfmt.Registry) error {

	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subnet) UnmarshalBinary(b []byte) error {
	var res Subnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}