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

// checks if the ErrorLogsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogsData{}

// ErrorLogsData struct for ErrorLogsData
type ErrorLogsData struct {
	Data []ErrorLog `json:"data,omitempty"`
}

// NewErrorLogsData instantiates a new ErrorLogsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogsData() *ErrorLogsData {
	this := ErrorLogsData{}
	return &this
}

// NewErrorLogsDataWithDefaults instantiates a new ErrorLogsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogsDataWithDefaults() *ErrorLogsData {
	this := ErrorLogsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ErrorLogsData) GetData() []ErrorLog {
	if o == nil || IsNil(o.Data) {
		var ret []ErrorLog
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogsData) GetDataOk() ([]ErrorLog, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ErrorLogsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ErrorLog and assigns it to the Data field.
func (o *ErrorLogsData) SetData(v []ErrorLog) {
	o.Data = v
}

func (o ErrorLogsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableErrorLogsData struct {
	value *ErrorLogsData
	isSet bool
}

func (v NullableErrorLogsData) Get() *ErrorLogsData {
	return v.value
}

func (v *NullableErrorLogsData) Set(val *ErrorLogsData) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogsData) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogsData(val *ErrorLogsData) *NullableErrorLogsData {
	return &NullableErrorLogsData{value: val, isSet: true}
}

func (v NullableErrorLogsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


