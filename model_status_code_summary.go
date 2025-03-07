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

// checks if the StatusCodeSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeSummary{}

// StatusCodeSummary struct for StatusCodeSummary
type StatusCodeSummary struct {
	// Deprecated
	Statistics map[string]interface{} `json:"statistics,omitempty"`
	Charts *StatusCodeSummaryCharts `json:"charts,omitempty"`
}

// NewStatusCodeSummary instantiates a new StatusCodeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeSummary() *StatusCodeSummary {
	this := StatusCodeSummary{}
	return &this
}

// NewStatusCodeSummaryWithDefaults instantiates a new StatusCodeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeSummaryWithDefaults() *StatusCodeSummary {
	this := StatusCodeSummary{}
	return &this
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
// Deprecated
func (o *StatusCodeSummary) GetStatistics() map[string]interface{} {
	if o == nil || IsNil(o.Statistics) {
		var ret map[string]interface{}
		return ret
	}
	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *StatusCodeSummary) GetStatisticsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Statistics) {
		return map[string]interface{}{}, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *StatusCodeSummary) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given map[string]interface{} and assigns it to the Statistics field.
// Deprecated
func (o *StatusCodeSummary) SetStatistics(v map[string]interface{}) {
	o.Statistics = v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *StatusCodeSummary) GetCharts() StatusCodeSummaryCharts {
	if o == nil || IsNil(o.Charts) {
		var ret StatusCodeSummaryCharts
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeSummary) GetChartsOk() (*StatusCodeSummaryCharts, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *StatusCodeSummary) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given StatusCodeSummaryCharts and assigns it to the Charts field.
func (o *StatusCodeSummary) SetCharts(v StatusCodeSummaryCharts) {
	o.Charts = &v
}

func (o StatusCodeSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	return toSerialize, nil
}

type NullableStatusCodeSummary struct {
	value *StatusCodeSummary
	isSet bool
}

func (v NullableStatusCodeSummary) Get() *StatusCodeSummary {
	return v.value
}

func (v *NullableStatusCodeSummary) Set(val *StatusCodeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeSummary(val *StatusCodeSummary) *NullableStatusCodeSummary {
	return &NullableStatusCodeSummary{value: val, isSet: true}
}

func (v NullableStatusCodeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


