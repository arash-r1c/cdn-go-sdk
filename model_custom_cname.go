/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the CustomCname type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomCname{}

// CustomCname struct for CustomCname
type CustomCname struct {
	Address string `json:"address"`
}

// NewCustomCname instantiates a new CustomCname object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCname(address string) *CustomCname {
	this := CustomCname{}
	this.Address = address
	return &this
}

// NewCustomCnameWithDefaults instantiates a new CustomCname object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCnameWithDefaults() *CustomCname {
	this := CustomCname{}
	return &this
}

// GetAddress returns the Address field value
func (o *CustomCname) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CustomCname) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CustomCname) SetAddress(v string) {
	o.Address = v
}

func (o CustomCname) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomCname) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableCustomCname struct {
	value *CustomCname
	isSet bool
}

func (v NullableCustomCname) Get() *CustomCname {
	return v.value
}

func (v *NullableCustomCname) Set(val *CustomCname) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCname) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCname) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCname(val *CustomCname) *NullableCustomCname {
	return &NullableCustomCname{value: val, isSet: true}
}

func (v NullableCustomCname) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCname) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


