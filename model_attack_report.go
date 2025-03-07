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

// checks if the AttackReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReport{}

// AttackReport struct for AttackReport
type AttackReport struct {
	Statistics *AttackReportStatistics `json:"statistics,omitempty"`
	Charts *AttackReportCharts `json:"charts,omitempty"`
}

// NewAttackReport instantiates a new AttackReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReport() *AttackReport {
	this := AttackReport{}
	return &this
}

// NewAttackReportWithDefaults instantiates a new AttackReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportWithDefaults() *AttackReport {
	this := AttackReport{}
	return &this
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *AttackReport) GetStatistics() AttackReportStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret AttackReportStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReport) GetStatisticsOk() (*AttackReportStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *AttackReport) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given AttackReportStatistics and assigns it to the Statistics field.
func (o *AttackReport) SetStatistics(v AttackReportStatistics) {
	o.Statistics = &v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *AttackReport) GetCharts() AttackReportCharts {
	if o == nil || IsNil(o.Charts) {
		var ret AttackReportCharts
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReport) GetChartsOk() (*AttackReportCharts, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *AttackReport) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given AttackReportCharts and assigns it to the Charts field.
func (o *AttackReport) SetCharts(v AttackReportCharts) {
	o.Charts = &v
}

func (o AttackReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	return toSerialize, nil
}

type NullableAttackReport struct {
	value *AttackReport
	isSet bool
}

func (v NullableAttackReport) Get() *AttackReport {
	return v.value
}

func (v *NullableAttackReport) Set(val *AttackReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReport(val *AttackReport) *NullableAttackReport {
	return &NullableAttackReport{value: val, isSet: true}
}

func (v NullableAttackReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


