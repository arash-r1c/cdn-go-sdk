/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.115.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the AccelerationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccelerationUpdate{}

// AccelerationUpdate struct for AccelerationUpdate
type AccelerationUpdate struct {
	Status *string `json:"status,omitempty"`
	Extensions []string `json:"extensions,omitempty"`
}

// NewAccelerationUpdate instantiates a new AccelerationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccelerationUpdate() *AccelerationUpdate {
	this := AccelerationUpdate{}
	return &this
}

// NewAccelerationUpdateWithDefaults instantiates a new AccelerationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccelerationUpdateWithDefaults() *AccelerationUpdate {
	this := AccelerationUpdate{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccelerationUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccelerationUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccelerationUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccelerationUpdate) SetStatus(v string) {
	o.Status = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *AccelerationUpdate) GetExtensions() []string {
	if o == nil || IsNil(o.Extensions) {
		var ret []string
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccelerationUpdate) GetExtensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *AccelerationUpdate) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []string and assigns it to the Extensions field.
func (o *AccelerationUpdate) SetExtensions(v []string) {
	o.Extensions = v
}

func (o AccelerationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccelerationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	return toSerialize, nil
}

type NullableAccelerationUpdate struct {
	value *AccelerationUpdate
	isSet bool
}

func (v NullableAccelerationUpdate) Get() *AccelerationUpdate {
	return v.value
}

func (v *NullableAccelerationUpdate) Set(val *AccelerationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccelerationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccelerationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccelerationUpdate(val *AccelerationUpdate) *NullableAccelerationUpdate {
	return &NullableAccelerationUpdate{value: val, isSet: true}
}

func (v NullableAccelerationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccelerationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


