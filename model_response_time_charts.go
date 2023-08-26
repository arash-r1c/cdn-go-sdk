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

// checks if the ResponseTimeCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeCharts{}

// ResponseTimeCharts struct for ResponseTimeCharts
type ResponseTimeCharts struct {
	Ir *ResponseTimeChartsIr `json:"ir,omitempty"`
}

// NewResponseTimeCharts instantiates a new ResponseTimeCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeCharts() *ResponseTimeCharts {
	this := ResponseTimeCharts{}
	return &this
}

// NewResponseTimeChartsWithDefaults instantiates a new ResponseTimeCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeChartsWithDefaults() *ResponseTimeCharts {
	this := ResponseTimeCharts{}
	return &this
}

// GetIr returns the Ir field value if set, zero value otherwise.
func (o *ResponseTimeCharts) GetIr() ResponseTimeChartsIr {
	if o == nil || IsNil(o.Ir) {
		var ret ResponseTimeChartsIr
		return ret
	}
	return *o.Ir
}

// GetIrOk returns a tuple with the Ir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeCharts) GetIrOk() (*ResponseTimeChartsIr, bool) {
	if o == nil || IsNil(o.Ir) {
		return nil, false
	}
	return o.Ir, true
}

// HasIr returns a boolean if a field has been set.
func (o *ResponseTimeCharts) HasIr() bool {
	if o != nil && !IsNil(o.Ir) {
		return true
	}

	return false
}

// SetIr gets a reference to the given ResponseTimeChartsIr and assigns it to the Ir field.
func (o *ResponseTimeCharts) SetIr(v ResponseTimeChartsIr) {
	o.Ir = &v
}

func (o ResponseTimeCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ir) {
		toSerialize["ir"] = o.Ir
	}
	return toSerialize, nil
}

type NullableResponseTimeCharts struct {
	value *ResponseTimeCharts
	isSet bool
}

func (v NullableResponseTimeCharts) Get() *ResponseTimeCharts {
	return v.value
}

func (v *NullableResponseTimeCharts) Set(val *ResponseTimeCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeCharts(val *ResponseTimeCharts) *NullableResponseTimeCharts {
	return &NullableResponseTimeCharts{value: val, isSet: true}
}

func (v NullableResponseTimeCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


