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

// checks if the DeprecatedNs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedNs{}

// DeprecatedNs struct for DeprecatedNs
type DeprecatedNs struct {
	// Current NS records of the domain
	NsDomain []string `json:"ns_domain,omitempty"`
	// Desired NS values for the domain
	NsKeys []string `json:"ns_keys,omitempty"`
	NsStatus *bool `json:"ns_status,omitempty"`
}

// NewDeprecatedNs instantiates a new DeprecatedNs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedNs() *DeprecatedNs {
	this := DeprecatedNs{}
	return &this
}

// NewDeprecatedNsWithDefaults instantiates a new DeprecatedNs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedNsWithDefaults() *DeprecatedNs {
	this := DeprecatedNs{}
	return &this
}

// GetNsDomain returns the NsDomain field value if set, zero value otherwise.
func (o *DeprecatedNs) GetNsDomain() []string {
	if o == nil || IsNil(o.NsDomain) {
		var ret []string
		return ret
	}
	return o.NsDomain
}

// GetNsDomainOk returns a tuple with the NsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedNs) GetNsDomainOk() ([]string, bool) {
	if o == nil || IsNil(o.NsDomain) {
		return nil, false
	}
	return o.NsDomain, true
}

// HasNsDomain returns a boolean if a field has been set.
func (o *DeprecatedNs) HasNsDomain() bool {
	if o != nil && !IsNil(o.NsDomain) {
		return true
	}

	return false
}

// SetNsDomain gets a reference to the given []string and assigns it to the NsDomain field.
func (o *DeprecatedNs) SetNsDomain(v []string) {
	o.NsDomain = v
}

// GetNsKeys returns the NsKeys field value if set, zero value otherwise.
func (o *DeprecatedNs) GetNsKeys() []string {
	if o == nil || IsNil(o.NsKeys) {
		var ret []string
		return ret
	}
	return o.NsKeys
}

// GetNsKeysOk returns a tuple with the NsKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedNs) GetNsKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.NsKeys) {
		return nil, false
	}
	return o.NsKeys, true
}

// HasNsKeys returns a boolean if a field has been set.
func (o *DeprecatedNs) HasNsKeys() bool {
	if o != nil && !IsNil(o.NsKeys) {
		return true
	}

	return false
}

// SetNsKeys gets a reference to the given []string and assigns it to the NsKeys field.
func (o *DeprecatedNs) SetNsKeys(v []string) {
	o.NsKeys = v
}

// GetNsStatus returns the NsStatus field value if set, zero value otherwise.
func (o *DeprecatedNs) GetNsStatus() bool {
	if o == nil || IsNil(o.NsStatus) {
		var ret bool
		return ret
	}
	return *o.NsStatus
}

// GetNsStatusOk returns a tuple with the NsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedNs) GetNsStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.NsStatus) {
		return nil, false
	}
	return o.NsStatus, true
}

// HasNsStatus returns a boolean if a field has been set.
func (o *DeprecatedNs) HasNsStatus() bool {
	if o != nil && !IsNil(o.NsStatus) {
		return true
	}

	return false
}

// SetNsStatus gets a reference to the given bool and assigns it to the NsStatus field.
func (o *DeprecatedNs) SetNsStatus(v bool) {
	o.NsStatus = &v
}

func (o DeprecatedNs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedNs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NsDomain) {
		toSerialize["ns_domain"] = o.NsDomain
	}
	if !IsNil(o.NsKeys) {
		toSerialize["ns_keys"] = o.NsKeys
	}
	if !IsNil(o.NsStatus) {
		toSerialize["ns_status"] = o.NsStatus
	}
	return toSerialize, nil
}

type NullableDeprecatedNs struct {
	value *DeprecatedNs
	isSet bool
}

func (v NullableDeprecatedNs) Get() *DeprecatedNs {
	return v.value
}

func (v *NullableDeprecatedNs) Set(val *DeprecatedNs) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedNs) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedNs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedNs(val *DeprecatedNs) *NullableDeprecatedNs {
	return &NullableDeprecatedNs{value: val, isSet: true}
}

func (v NullableDeprecatedNs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedNs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


