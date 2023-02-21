/*
WhyLabs Songbird

WhyLabs API that enables end-to-end AI observability

API version: 0.1
Contact: support@whylabs.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RemoveMembershipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveMembershipRequest{}

// RemoveMembershipRequest struct for RemoveMembershipRequest
type RemoveMembershipRequest struct {
	OrgId string `json:"orgId"`
	Email string `json:"email"`
}

// NewRemoveMembershipRequest instantiates a new RemoveMembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveMembershipRequest(orgId string, email string) *RemoveMembershipRequest {
	this := RemoveMembershipRequest{}
	this.OrgId = orgId
	this.Email = email
	return &this
}

// NewRemoveMembershipRequestWithDefaults instantiates a new RemoveMembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveMembershipRequestWithDefaults() *RemoveMembershipRequest {
	this := RemoveMembershipRequest{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *RemoveMembershipRequest) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RemoveMembershipRequest) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *RemoveMembershipRequest) SetOrgId(v string) {
	o.OrgId = v
}

// GetEmail returns the Email field value
func (o *RemoveMembershipRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RemoveMembershipRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RemoveMembershipRequest) SetEmail(v string) {
	o.Email = v
}

func (o RemoveMembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveMembershipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orgId"] = o.OrgId
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableRemoveMembershipRequest struct {
	value *RemoveMembershipRequest
	isSet bool
}

func (v NullableRemoveMembershipRequest) Get() *RemoveMembershipRequest {
	return v.value
}

func (v *NullableRemoveMembershipRequest) Set(val *RemoveMembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveMembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveMembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveMembershipRequest(val *RemoveMembershipRequest) *NullableRemoveMembershipRequest {
	return &NullableRemoveMembershipRequest{value: val, isSet: true}
}

func (v NullableRemoveMembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveMembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

