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

// checks if the HealthCheckReportSummaryDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckReportSummaryDetail{}

// HealthCheckReportSummaryDetail struct for HealthCheckReportSummaryDetail
type HealthCheckReportSummaryDetail struct {
	Date *string `json:"date,omitempty"`
	Status *bool `json:"status,omitempty"`
}

// NewHealthCheckReportSummaryDetail instantiates a new HealthCheckReportSummaryDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckReportSummaryDetail() *HealthCheckReportSummaryDetail {
	this := HealthCheckReportSummaryDetail{}
	return &this
}

// NewHealthCheckReportSummaryDetailWithDefaults instantiates a new HealthCheckReportSummaryDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckReportSummaryDetailWithDefaults() *HealthCheckReportSummaryDetail {
	this := HealthCheckReportSummaryDetail{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *HealthCheckReportSummaryDetail) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummaryDetail) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *HealthCheckReportSummaryDetail) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *HealthCheckReportSummaryDetail) SetDate(v string) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheckReportSummaryDetail) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummaryDetail) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheckReportSummaryDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HealthCheckReportSummaryDetail) SetStatus(v bool) {
	o.Status = &v
}

func (o HealthCheckReportSummaryDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckReportSummaryDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHealthCheckReportSummaryDetail struct {
	value *HealthCheckReportSummaryDetail
	isSet bool
}

func (v NullableHealthCheckReportSummaryDetail) Get() *HealthCheckReportSummaryDetail {
	return v.value
}

func (v *NullableHealthCheckReportSummaryDetail) Set(val *HealthCheckReportSummaryDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckReportSummaryDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckReportSummaryDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckReportSummaryDetail(val *HealthCheckReportSummaryDetail) *NullableHealthCheckReportSummaryDetail {
	return &NullableHealthCheckReportSummaryDetail{value: val, isSet: true}
}

func (v NullableHealthCheckReportSummaryDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckReportSummaryDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


