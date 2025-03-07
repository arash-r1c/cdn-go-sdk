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

// checks if the CacheSettingsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CacheSettingsData{}

// CacheSettingsData struct for CacheSettingsData
type CacheSettingsData struct {
	Data *CacheSettings `json:"data,omitempty"`
}

// NewCacheSettingsData instantiates a new CacheSettingsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheSettingsData() *CacheSettingsData {
	this := CacheSettingsData{}
	return &this
}

// NewCacheSettingsDataWithDefaults instantiates a new CacheSettingsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheSettingsDataWithDefaults() *CacheSettingsData {
	this := CacheSettingsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CacheSettingsData) GetData() CacheSettings {
	if o == nil || IsNil(o.Data) {
		var ret CacheSettings
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsData) GetDataOk() (*CacheSettings, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CacheSettingsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CacheSettings and assigns it to the Data field.
func (o *CacheSettingsData) SetData(v CacheSettings) {
	o.Data = &v
}

func (o CacheSettingsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CacheSettingsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCacheSettingsData struct {
	value *CacheSettingsData
	isSet bool
}

func (v NullableCacheSettingsData) Get() *CacheSettingsData {
	return v.value
}

func (v *NullableCacheSettingsData) Set(val *CacheSettingsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheSettingsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheSettingsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheSettingsData(val *CacheSettingsData) *NullableCacheSettingsData {
	return &NullableCacheSettingsData{value: val, isSet: true}
}

func (v NullableCacheSettingsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheSettingsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


