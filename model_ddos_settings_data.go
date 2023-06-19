/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the DdosSettingsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdosSettingsData{}

// DdosSettingsData struct for DdosSettingsData
type DdosSettingsData struct {
	Data *DdosSettings `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewDdosSettingsData instantiates a new DdosSettingsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosSettingsData() *DdosSettingsData {
	this := DdosSettingsData{}
	return &this
}

// NewDdosSettingsDataWithDefaults instantiates a new DdosSettingsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosSettingsDataWithDefaults() *DdosSettingsData {
	this := DdosSettingsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DdosSettingsData) GetData() DdosSettings {
	if o == nil || IsNil(o.Data) {
		var ret DdosSettings
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettingsData) GetDataOk() (*DdosSettings, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DdosSettingsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DdosSettings and assigns it to the Data field.
func (o *DdosSettingsData) SetData(v DdosSettings) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosSettingsData) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosSettingsData) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DdosSettingsData) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *DdosSettingsData) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *DdosSettingsData) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *DdosSettingsData) UnsetMessage() {
	o.Message.Unset()
}

func (o DdosSettingsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdosSettingsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableDdosSettingsData struct {
	value *DdosSettingsData
	isSet bool
}

func (v NullableDdosSettingsData) Get() *DdosSettingsData {
	return v.value
}

func (v *NullableDdosSettingsData) Set(val *DdosSettingsData) {
	v.value = val
	v.isSet = true
}

func (v NullableDdosSettingsData) IsSet() bool {
	return v.isSet
}

func (v *NullableDdosSettingsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdosSettingsData(val *DdosSettingsData) *NullableDdosSettingsData {
	return &NullableDdosSettingsData{value: val, isSet: true}
}

func (v NullableDdosSettingsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdosSettingsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


