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

// checks if the RateLimitData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitData{}

// RateLimitData struct for RateLimitData
type RateLimitData struct {
	Data *RateLimit `json:"data,omitempty"`
}

// NewRateLimitData instantiates a new RateLimitData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitData() *RateLimitData {
	this := RateLimitData{}
	return &this
}

// NewRateLimitDataWithDefaults instantiates a new RateLimitData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitDataWithDefaults() *RateLimitData {
	this := RateLimitData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RateLimitData) GetData() RateLimit {
	if o == nil || IsNil(o.Data) {
		var ret RateLimit
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitData) GetDataOk() (*RateLimit, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RateLimitData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RateLimit and assigns it to the Data field.
func (o *RateLimitData) SetData(v RateLimit) {
	o.Data = &v
}

func (o RateLimitData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRateLimitData struct {
	value *RateLimitData
	isSet bool
}

func (v NullableRateLimitData) Get() *RateLimitData {
	return v.value
}

func (v *NullableRateLimitData) Set(val *RateLimitData) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitData) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitData(val *RateLimitData) *NullableRateLimitData {
	return &NullableRateLimitData{value: val, isSet: true}
}

func (v NullableRateLimitData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


