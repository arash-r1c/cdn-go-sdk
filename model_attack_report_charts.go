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

// checks if the AttackReportCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportCharts{}

// AttackReportCharts struct for AttackReportCharts
type AttackReportCharts struct {
	Attacks *AttackReportChartsAttacks `json:"attacks,omitempty"`
}

// NewAttackReportCharts instantiates a new AttackReportCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportCharts() *AttackReportCharts {
	this := AttackReportCharts{}
	return &this
}

// NewAttackReportChartsWithDefaults instantiates a new AttackReportCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportChartsWithDefaults() *AttackReportCharts {
	this := AttackReportCharts{}
	return &this
}

// GetAttacks returns the Attacks field value if set, zero value otherwise.
func (o *AttackReportCharts) GetAttacks() AttackReportChartsAttacks {
	if o == nil || IsNil(o.Attacks) {
		var ret AttackReportChartsAttacks
		return ret
	}
	return *o.Attacks
}

// GetAttacksOk returns a tuple with the Attacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportCharts) GetAttacksOk() (*AttackReportChartsAttacks, bool) {
	if o == nil || IsNil(o.Attacks) {
		return nil, false
	}
	return o.Attacks, true
}

// HasAttacks returns a boolean if a field has been set.
func (o *AttackReportCharts) HasAttacks() bool {
	if o != nil && !IsNil(o.Attacks) {
		return true
	}

	return false
}

// SetAttacks gets a reference to the given AttackReportChartsAttacks and assigns it to the Attacks field.
func (o *AttackReportCharts) SetAttacks(v AttackReportChartsAttacks) {
	o.Attacks = &v
}

func (o AttackReportCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attacks) {
		toSerialize["attacks"] = o.Attacks
	}
	return toSerialize, nil
}

type NullableAttackReportCharts struct {
	value *AttackReportCharts
	isSet bool
}

func (v NullableAttackReportCharts) Get() *AttackReportCharts {
	return v.value
}

func (v *NullableAttackReportCharts) Set(val *AttackReportCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportCharts(val *AttackReportCharts) *NullableAttackReportCharts {
	return &NullableAttackReportCharts{value: val, isSet: true}
}

func (v NullableAttackReportCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


