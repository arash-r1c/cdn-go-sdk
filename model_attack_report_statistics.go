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

// checks if the AttackReportStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportStatistics{}

// AttackReportStatistics struct for AttackReportStatistics
type AttackReportStatistics struct {
	Attacks *DnsRequestReportStatistics `json:"attacks,omitempty"`
}

// NewAttackReportStatistics instantiates a new AttackReportStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportStatistics() *AttackReportStatistics {
	this := AttackReportStatistics{}
	return &this
}

// NewAttackReportStatisticsWithDefaults instantiates a new AttackReportStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportStatisticsWithDefaults() *AttackReportStatistics {
	this := AttackReportStatistics{}
	return &this
}

// GetAttacks returns the Attacks field value if set, zero value otherwise.
func (o *AttackReportStatistics) GetAttacks() DnsRequestReportStatistics {
	if o == nil || IsNil(o.Attacks) {
		var ret DnsRequestReportStatistics
		return ret
	}
	return *o.Attacks
}

// GetAttacksOk returns a tuple with the Attacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportStatistics) GetAttacksOk() (*DnsRequestReportStatistics, bool) {
	if o == nil || IsNil(o.Attacks) {
		return nil, false
	}
	return o.Attacks, true
}

// HasAttacks returns a boolean if a field has been set.
func (o *AttackReportStatistics) HasAttacks() bool {
	if o != nil && !IsNil(o.Attacks) {
		return true
	}

	return false
}

// SetAttacks gets a reference to the given DnsRequestReportStatistics and assigns it to the Attacks field.
func (o *AttackReportStatistics) SetAttacks(v DnsRequestReportStatistics) {
	o.Attacks = &v
}

func (o AttackReportStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attacks) {
		toSerialize["attacks"] = o.Attacks
	}
	return toSerialize, nil
}

type NullableAttackReportStatistics struct {
	value *AttackReportStatistics
	isSet bool
}

func (v NullableAttackReportStatistics) Get() *AttackReportStatistics {
	return v.value
}

func (v *NullableAttackReportStatistics) Set(val *AttackReportStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportStatistics(val *AttackReportStatistics) *NullableAttackReportStatistics {
	return &NullableAttackReportStatistics{value: val, isSet: true}
}

func (v NullableAttackReportStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


