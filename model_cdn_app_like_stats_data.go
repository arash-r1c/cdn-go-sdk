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

// checks if the CdnAppLikeStatsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdnAppLikeStatsData{}

// CdnAppLikeStatsData struct for CdnAppLikeStatsData
type CdnAppLikeStatsData struct {
	Data *CdnAppLikeStats `json:"data,omitempty"`
}

// NewCdnAppLikeStatsData instantiates a new CdnAppLikeStatsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnAppLikeStatsData() *CdnAppLikeStatsData {
	this := CdnAppLikeStatsData{}
	return &this
}

// NewCdnAppLikeStatsDataWithDefaults instantiates a new CdnAppLikeStatsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnAppLikeStatsDataWithDefaults() *CdnAppLikeStatsData {
	this := CdnAppLikeStatsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CdnAppLikeStatsData) GetData() CdnAppLikeStats {
	if o == nil || IsNil(o.Data) {
		var ret CdnAppLikeStats
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnAppLikeStatsData) GetDataOk() (*CdnAppLikeStats, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CdnAppLikeStatsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CdnAppLikeStats and assigns it to the Data field.
func (o *CdnAppLikeStatsData) SetData(v CdnAppLikeStats) {
	o.Data = &v
}

func (o CdnAppLikeStatsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdnAppLikeStatsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCdnAppLikeStatsData struct {
	value *CdnAppLikeStatsData
	isSet bool
}

func (v NullableCdnAppLikeStatsData) Get() *CdnAppLikeStatsData {
	return v.value
}

func (v *NullableCdnAppLikeStatsData) Set(val *CdnAppLikeStatsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnAppLikeStatsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnAppLikeStatsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnAppLikeStatsData(val *CdnAppLikeStatsData) *NullableCdnAppLikeStatsData {
	return &NullableCdnAppLikeStatsData{value: val, isSet: true}
}

func (v NullableCdnAppLikeStatsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnAppLikeStatsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


