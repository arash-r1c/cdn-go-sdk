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

// checks if the AttackReportMapStatisticsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportMapStatisticsInner{}

// AttackReportMapStatisticsInner struct for AttackReportMapStatisticsInner
type AttackReportMapStatisticsInner struct {
	// The 2-letter country code
	Country *string `json:"country,omitempty"`
	// The name of the country
	Name *string `json:"name,omitempty"`
	// The 3-letter country code
	Code *string `json:"code,omitempty"`
	// The number of attacks
	Attack *int32 `json:"attack,omitempty"`
}

// NewAttackReportMapStatisticsInner instantiates a new AttackReportMapStatisticsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportMapStatisticsInner() *AttackReportMapStatisticsInner {
	this := AttackReportMapStatisticsInner{}
	return &this
}

// NewAttackReportMapStatisticsInnerWithDefaults instantiates a new AttackReportMapStatisticsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportMapStatisticsInnerWithDefaults() *AttackReportMapStatisticsInner {
	this := AttackReportMapStatisticsInner{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AttackReportMapStatisticsInner) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapStatisticsInner) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AttackReportMapStatisticsInner) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AttackReportMapStatisticsInner) SetCountry(v string) {
	o.Country = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttackReportMapStatisticsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapStatisticsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttackReportMapStatisticsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttackReportMapStatisticsInner) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AttackReportMapStatisticsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapStatisticsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AttackReportMapStatisticsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AttackReportMapStatisticsInner) SetCode(v string) {
	o.Code = &v
}

// GetAttack returns the Attack field value if set, zero value otherwise.
func (o *AttackReportMapStatisticsInner) GetAttack() int32 {
	if o == nil || IsNil(o.Attack) {
		var ret int32
		return ret
	}
	return *o.Attack
}

// GetAttackOk returns a tuple with the Attack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapStatisticsInner) GetAttackOk() (*int32, bool) {
	if o == nil || IsNil(o.Attack) {
		return nil, false
	}
	return o.Attack, true
}

// HasAttack returns a boolean if a field has been set.
func (o *AttackReportMapStatisticsInner) HasAttack() bool {
	if o != nil && !IsNil(o.Attack) {
		return true
	}

	return false
}

// SetAttack gets a reference to the given int32 and assigns it to the Attack field.
func (o *AttackReportMapStatisticsInner) SetAttack(v int32) {
	o.Attack = &v
}

func (o AttackReportMapStatisticsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportMapStatisticsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Attack) {
		toSerialize["attack"] = o.Attack
	}
	return toSerialize, nil
}

type NullableAttackReportMapStatisticsInner struct {
	value *AttackReportMapStatisticsInner
	isSet bool
}

func (v NullableAttackReportMapStatisticsInner) Get() *AttackReportMapStatisticsInner {
	return v.value
}

func (v *NullableAttackReportMapStatisticsInner) Set(val *AttackReportMapStatisticsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportMapStatisticsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportMapStatisticsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportMapStatisticsInner(val *AttackReportMapStatisticsInner) *NullableAttackReportMapStatisticsInner {
	return &NullableAttackReportMapStatisticsInner{value: val, isSet: true}
}

func (v NullableAttackReportMapStatisticsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportMapStatisticsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


