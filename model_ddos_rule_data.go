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

// checks if the DdosRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdosRuleData{}

// DdosRuleData struct for DdosRuleData
type DdosRuleData struct {
	Data *DdosRule `json:"data,omitempty"`
}

// NewDdosRuleData instantiates a new DdosRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosRuleData() *DdosRuleData {
	this := DdosRuleData{}
	return &this
}

// NewDdosRuleDataWithDefaults instantiates a new DdosRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosRuleDataWithDefaults() *DdosRuleData {
	this := DdosRuleData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DdosRuleData) GetData() DdosRule {
	if o == nil || IsNil(o.Data) {
		var ret DdosRule
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRuleData) GetDataOk() (*DdosRule, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DdosRuleData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DdosRule and assigns it to the Data field.
func (o *DdosRuleData) SetData(v DdosRule) {
	o.Data = &v
}

func (o DdosRuleData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdosRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDdosRuleData struct {
	value *DdosRuleData
	isSet bool
}

func (v NullableDdosRuleData) Get() *DdosRuleData {
	return v.value
}

func (v *NullableDdosRuleData) Set(val *DdosRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableDdosRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableDdosRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdosRuleData(val *DdosRuleData) *NullableDdosRuleData {
	return &NullableDdosRuleData{value: val, isSet: true}
}

func (v NullableDdosRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdosRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


