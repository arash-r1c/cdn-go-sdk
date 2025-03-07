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

// checks if the NsKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsKeys{}

// NsKeys struct for NsKeys
type NsKeys struct {
	// Desired NS values for the domain
	NsKeys []string `json:"ns_keys,omitempty"`
}

// NewNsKeys instantiates a new NsKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsKeys() *NsKeys {
	this := NsKeys{}
	return &this
}

// NewNsKeysWithDefaults instantiates a new NsKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsKeysWithDefaults() *NsKeys {
	this := NsKeys{}
	return &this
}

// GetNsKeys returns the NsKeys field value if set, zero value otherwise.
func (o *NsKeys) GetNsKeys() []string {
	if o == nil || IsNil(o.NsKeys) {
		var ret []string
		return ret
	}
	return o.NsKeys
}

// GetNsKeysOk returns a tuple with the NsKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsKeys) GetNsKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.NsKeys) {
		return nil, false
	}
	return o.NsKeys, true
}

// HasNsKeys returns a boolean if a field has been set.
func (o *NsKeys) HasNsKeys() bool {
	if o != nil && !IsNil(o.NsKeys) {
		return true
	}

	return false
}

// SetNsKeys gets a reference to the given []string and assigns it to the NsKeys field.
func (o *NsKeys) SetNsKeys(v []string) {
	o.NsKeys = v
}

func (o NsKeys) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NsKeys) {
		toSerialize["ns_keys"] = o.NsKeys
	}
	return toSerialize, nil
}

type NullableNsKeys struct {
	value *NsKeys
	isSet bool
}

func (v NullableNsKeys) Get() *NsKeys {
	return v.value
}

func (v *NullableNsKeys) Set(val *NsKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableNsKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableNsKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsKeys(val *NsKeys) *NullableNsKeys {
	return &NullableNsKeys{value: val, isSet: true}
}

func (v NullableNsKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


