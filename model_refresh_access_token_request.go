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

// checks if the RefreshAccessTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshAccessTokenRequest{}

// RefreshAccessTokenRequest struct for RefreshAccessTokenRequest
type RefreshAccessTokenRequest struct {
	WorkspaceId string `json:"workspaceId"`
}

// NewRefreshAccessTokenRequest instantiates a new RefreshAccessTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshAccessTokenRequest(workspaceId string) *RefreshAccessTokenRequest {
	this := RefreshAccessTokenRequest{}
	this.WorkspaceId = workspaceId
	return &this
}

// NewRefreshAccessTokenRequestWithDefaults instantiates a new RefreshAccessTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshAccessTokenRequestWithDefaults() *RefreshAccessTokenRequest {
	this := RefreshAccessTokenRequest{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *RefreshAccessTokenRequest) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *RefreshAccessTokenRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *RefreshAccessTokenRequest) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o RefreshAccessTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshAccessTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspaceId"] = o.WorkspaceId
	return toSerialize, nil
}

type NullableRefreshAccessTokenRequest struct {
	value *RefreshAccessTokenRequest
	isSet bool
}

func (v NullableRefreshAccessTokenRequest) Get() *RefreshAccessTokenRequest {
	return v.value
}

func (v *NullableRefreshAccessTokenRequest) Set(val *RefreshAccessTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshAccessTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshAccessTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshAccessTokenRequest(val *RefreshAccessTokenRequest) *NullableRefreshAccessTokenRequest {
	return &NullableRefreshAccessTokenRequest{value: val, isSet: true}
}

func (v NullableRefreshAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshAccessTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

