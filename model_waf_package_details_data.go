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

// checks if the WafPackageDetailsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPackageDetailsData{}

// WafPackageDetailsData struct for WafPackageDetailsData
type WafPackageDetailsData struct {
	Data *WafPackageDetails `json:"data,omitempty"`
}

// NewWafPackageDetailsData instantiates a new WafPackageDetailsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPackageDetailsData() *WafPackageDetailsData {
	this := WafPackageDetailsData{}
	return &this
}

// NewWafPackageDetailsDataWithDefaults instantiates a new WafPackageDetailsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPackageDetailsDataWithDefaults() *WafPackageDetailsData {
	this := WafPackageDetailsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafPackageDetailsData) GetData() WafPackageDetails {
	if o == nil || IsNil(o.Data) {
		var ret WafPackageDetails
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetailsData) GetDataOk() (*WafPackageDetails, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafPackageDetailsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WafPackageDetails and assigns it to the Data field.
func (o *WafPackageDetailsData) SetData(v WafPackageDetails) {
	o.Data = &v
}

func (o WafPackageDetailsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPackageDetailsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableWafPackageDetailsData struct {
	value *WafPackageDetailsData
	isSet bool
}

func (v NullableWafPackageDetailsData) Get() *WafPackageDetailsData {
	return v.value
}

func (v *NullableWafPackageDetailsData) Set(val *WafPackageDetailsData) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPackageDetailsData) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPackageDetailsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPackageDetailsData(val *WafPackageDetailsData) *NullableWafPackageDetailsData {
	return &NullableWafPackageDetailsData{value: val, isSet: true}
}

func (v NullableWafPackageDetailsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPackageDetailsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


