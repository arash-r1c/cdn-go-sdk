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

// checks if the ErrorLogChartChartsStatusCodeSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogChartChartsStatusCodeSeriesInner{}

// ErrorLogChartChartsStatusCodeSeriesInner struct for ErrorLogChartChartsStatusCodeSeriesInner
type ErrorLogChartChartsStatusCodeSeriesInner struct {
	// The error message
	Name *string `json:"name,omitempty"`
	Data []int64 `json:"data,omitempty"`
}

// NewErrorLogChartChartsStatusCodeSeriesInner instantiates a new ErrorLogChartChartsStatusCodeSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogChartChartsStatusCodeSeriesInner() *ErrorLogChartChartsStatusCodeSeriesInner {
	this := ErrorLogChartChartsStatusCodeSeriesInner{}
	return &this
}

// NewErrorLogChartChartsStatusCodeSeriesInnerWithDefaults instantiates a new ErrorLogChartChartsStatusCodeSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogChartChartsStatusCodeSeriesInnerWithDefaults() *ErrorLogChartChartsStatusCodeSeriesInner {
	this := ErrorLogChartChartsStatusCodeSeriesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) GetData() []int64 {
	if o == nil || IsNil(o.Data) {
		var ret []int64
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) GetDataOk() ([]int64, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int64 and assigns it to the Data field.
func (o *ErrorLogChartChartsStatusCodeSeriesInner) SetData(v []int64) {
	o.Data = v
}

func (o ErrorLogChartChartsStatusCodeSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogChartChartsStatusCodeSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableErrorLogChartChartsStatusCodeSeriesInner struct {
	value *ErrorLogChartChartsStatusCodeSeriesInner
	isSet bool
}

func (v NullableErrorLogChartChartsStatusCodeSeriesInner) Get() *ErrorLogChartChartsStatusCodeSeriesInner {
	return v.value
}

func (v *NullableErrorLogChartChartsStatusCodeSeriesInner) Set(val *ErrorLogChartChartsStatusCodeSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogChartChartsStatusCodeSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogChartChartsStatusCodeSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogChartChartsStatusCodeSeriesInner(val *ErrorLogChartChartsStatusCodeSeriesInner) *NullableErrorLogChartChartsStatusCodeSeriesInner {
	return &NullableErrorLogChartChartsStatusCodeSeriesInner{value: val, isSet: true}
}

func (v NullableErrorLogChartChartsStatusCodeSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogChartChartsStatusCodeSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


