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

// checks if the AttackReportMapData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportMapData{}

// AttackReportMapData struct for AttackReportMapData
type AttackReportMapData struct {
	Data *AttackReportMap `json:"data,omitempty"`
}

// NewAttackReportMapData instantiates a new AttackReportMapData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportMapData() *AttackReportMapData {
	this := AttackReportMapData{}
	return &this
}

// NewAttackReportMapDataWithDefaults instantiates a new AttackReportMapData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportMapDataWithDefaults() *AttackReportMapData {
	this := AttackReportMapData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttackReportMapData) GetData() AttackReportMap {
	if o == nil || IsNil(o.Data) {
		var ret AttackReportMap
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapData) GetDataOk() (*AttackReportMap, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttackReportMapData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttackReportMap and assigns it to the Data field.
func (o *AttackReportMapData) SetData(v AttackReportMap) {
	o.Data = &v
}

func (o AttackReportMapData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportMapData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAttackReportMapData struct {
	value *AttackReportMapData
	isSet bool
}

func (v NullableAttackReportMapData) Get() *AttackReportMapData {
	return v.value
}

func (v *NullableAttackReportMapData) Set(val *AttackReportMapData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportMapData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportMapData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportMapData(val *AttackReportMapData) *NullableAttackReportMapData {
	return &NullableAttackReportMapData{value: val, isSet: true}
}

func (v NullableAttackReportMapData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportMapData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


