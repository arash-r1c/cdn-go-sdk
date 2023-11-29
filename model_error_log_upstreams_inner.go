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

// checks if the ErrorLogUpstreamsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogUpstreamsInner{}

// ErrorLogUpstreamsInner struct for ErrorLogUpstreamsInner
type ErrorLogUpstreamsInner struct {
	// The upstream's address
	Address *string `json:"address,omitempty"`
	// Error count per upstream
	Count *int64 `json:"count,omitempty"`
}

// NewErrorLogUpstreamsInner instantiates a new ErrorLogUpstreamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogUpstreamsInner() *ErrorLogUpstreamsInner {
	this := ErrorLogUpstreamsInner{}
	return &this
}

// NewErrorLogUpstreamsInnerWithDefaults instantiates a new ErrorLogUpstreamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogUpstreamsInnerWithDefaults() *ErrorLogUpstreamsInner {
	this := ErrorLogUpstreamsInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ErrorLogUpstreamsInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogUpstreamsInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ErrorLogUpstreamsInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ErrorLogUpstreamsInner) SetAddress(v string) {
	o.Address = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ErrorLogUpstreamsInner) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogUpstreamsInner) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ErrorLogUpstreamsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ErrorLogUpstreamsInner) SetCount(v int64) {
	o.Count = &v
}

func (o ErrorLogUpstreamsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogUpstreamsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableErrorLogUpstreamsInner struct {
	value *ErrorLogUpstreamsInner
	isSet bool
}

func (v NullableErrorLogUpstreamsInner) Get() *ErrorLogUpstreamsInner {
	return v.value
}

func (v *NullableErrorLogUpstreamsInner) Set(val *ErrorLogUpstreamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogUpstreamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogUpstreamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogUpstreamsInner(val *ErrorLogUpstreamsInner) *NullableErrorLogUpstreamsInner {
	return &NullableErrorLogUpstreamsInner{value: val, isSet: true}
}

func (v NullableErrorLogUpstreamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogUpstreamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


