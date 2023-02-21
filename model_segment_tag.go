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

// checks if the SegmentTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentTag{}

// SegmentTag struct for SegmentTag
type SegmentTag struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewSegmentTag instantiates a new SegmentTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentTag() *SegmentTag {
	this := SegmentTag{}
	return &this
}

// NewSegmentTagWithDefaults instantiates a new SegmentTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentTagWithDefaults() *SegmentTag {
	this := SegmentTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SegmentTag) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentTag) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SegmentTag) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SegmentTag) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SegmentTag) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentTag) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SegmentTag) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SegmentTag) SetValue(v string) {
	o.Value = &v
}

func (o SegmentTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSegmentTag struct {
	value *SegmentTag
	isSet bool
}

func (v NullableSegmentTag) Get() *SegmentTag {
	return v.value
}

func (v *NullableSegmentTag) Set(val *SegmentTag) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentTag) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentTag(val *SegmentTag) *NullableSegmentTag {
	return &NullableSegmentTag{value: val, isSet: true}
}

func (v NullableSegmentTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

