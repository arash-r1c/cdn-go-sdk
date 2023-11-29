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

// checks if the WafPresetsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPresetsData{}

// WafPresetsData struct for WafPresetsData
type WafPresetsData struct {
	Data *WafPresets `json:"data,omitempty"`
}

// NewWafPresetsData instantiates a new WafPresetsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPresetsData() *WafPresetsData {
	this := WafPresetsData{}
	return &this
}

// NewWafPresetsDataWithDefaults instantiates a new WafPresetsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPresetsDataWithDefaults() *WafPresetsData {
	this := WafPresetsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafPresetsData) GetData() WafPresets {
	if o == nil || IsNil(o.Data) {
		var ret WafPresets
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPresetsData) GetDataOk() (*WafPresets, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafPresetsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WafPresets and assigns it to the Data field.
func (o *WafPresetsData) SetData(v WafPresets) {
	o.Data = &v
}

func (o WafPresetsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPresetsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableWafPresetsData struct {
	value *WafPresetsData
	isSet bool
}

func (v NullableWafPresetsData) Get() *WafPresetsData {
	return v.value
}

func (v *NullableWafPresetsData) Set(val *WafPresetsData) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPresetsData) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPresetsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPresetsData(val *WafPresetsData) *NullableWafPresetsData {
	return &NullableWafPresetsData{value: val, isSet: true}
}

func (v NullableWafPresetsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPresetsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


