/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.106.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the NsDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsDomain{}

// NsDomain struct for NsDomain
type NsDomain struct {
	// Current NS records of the domain
	NsDomain []string `json:"ns_domain,omitempty"`
	// Desired NS values for the domain
	NsKeys []string `json:"ns_keys,omitempty"`
}

// NewNsDomain instantiates a new NsDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsDomain() *NsDomain {
	this := NsDomain{}
	return &this
}

// NewNsDomainWithDefaults instantiates a new NsDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsDomainWithDefaults() *NsDomain {
	this := NsDomain{}
	return &this
}

// GetNsDomain returns the NsDomain field value if set, zero value otherwise.
func (o *NsDomain) GetNsDomain() []string {
	if o == nil || IsNil(o.NsDomain) {
		var ret []string
		return ret
	}
	return o.NsDomain
}

// GetNsDomainOk returns a tuple with the NsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsDomain) GetNsDomainOk() ([]string, bool) {
	if o == nil || IsNil(o.NsDomain) {
		return nil, false
	}
	return o.NsDomain, true
}

// HasNsDomain returns a boolean if a field has been set.
func (o *NsDomain) HasNsDomain() bool {
	if o != nil && !IsNil(o.NsDomain) {
		return true
	}

	return false
}

// SetNsDomain gets a reference to the given []string and assigns it to the NsDomain field.
func (o *NsDomain) SetNsDomain(v []string) {
	o.NsDomain = v
}

// GetNsKeys returns the NsKeys field value if set, zero value otherwise.
func (o *NsDomain) GetNsKeys() []string {
	if o == nil || IsNil(o.NsKeys) {
		var ret []string
		return ret
	}
	return o.NsKeys
}

// GetNsKeysOk returns a tuple with the NsKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsDomain) GetNsKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.NsKeys) {
		return nil, false
	}
	return o.NsKeys, true
}

// HasNsKeys returns a boolean if a field has been set.
func (o *NsDomain) HasNsKeys() bool {
	if o != nil && !IsNil(o.NsKeys) {
		return true
	}

	return false
}

// SetNsKeys gets a reference to the given []string and assigns it to the NsKeys field.
func (o *NsDomain) SetNsKeys(v []string) {
	o.NsKeys = v
}

func (o NsDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NsDomain) {
		toSerialize["ns_domain"] = o.NsDomain
	}
	if !IsNil(o.NsKeys) {
		toSerialize["ns_keys"] = o.NsKeys
	}
	return toSerialize, nil
}

type NullableNsDomain struct {
	value *NsDomain
	isSet bool
}

func (v NullableNsDomain) Get() *NsDomain {
	return v.value
}

func (v *NullableNsDomain) Set(val *NsDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableNsDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableNsDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsDomain(val *NsDomain) *NullableNsDomain {
	return &NullableNsDomain{value: val, isSet: true}
}

func (v NullableNsDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


