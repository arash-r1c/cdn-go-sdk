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

// checks if the ResponseTimeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeData{}

// ResponseTimeData struct for ResponseTimeData
type ResponseTimeData struct {
	Data *ResponseTime `json:"data,omitempty"`
}

// NewResponseTimeData instantiates a new ResponseTimeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeData() *ResponseTimeData {
	this := ResponseTimeData{}
	return &this
}

// NewResponseTimeDataWithDefaults instantiates a new ResponseTimeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeDataWithDefaults() *ResponseTimeData {
	this := ResponseTimeData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseTimeData) GetData() ResponseTime {
	if o == nil || IsNil(o.Data) {
		var ret ResponseTime
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeData) GetDataOk() (*ResponseTime, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseTimeData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResponseTime and assigns it to the Data field.
func (o *ResponseTimeData) SetData(v ResponseTime) {
	o.Data = &v
}

func (o ResponseTimeData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableResponseTimeData struct {
	value *ResponseTimeData
	isSet bool
}

func (v NullableResponseTimeData) Get() *ResponseTimeData {
	return v.value
}

func (v *NullableResponseTimeData) Set(val *ResponseTimeData) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeData) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeData(val *ResponseTimeData) *NullableResponseTimeData {
	return &NullableResponseTimeData{value: val, isSet: true}
}

func (v NullableResponseTimeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


