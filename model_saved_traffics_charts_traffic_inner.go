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

// checks if the SavedTrafficsChartsTrafficInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrafficsChartsTrafficInner{}

// SavedTrafficsChartsTrafficInner struct for SavedTrafficsChartsTrafficInner
type SavedTrafficsChartsTrafficInner struct {
	Name *string `json:"name,omitempty"`
	Y *int32 `json:"y,omitempty"`
}

// NewSavedTrafficsChartsTrafficInner instantiates a new SavedTrafficsChartsTrafficInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrafficsChartsTrafficInner() *SavedTrafficsChartsTrafficInner {
	this := SavedTrafficsChartsTrafficInner{}
	return &this
}

// NewSavedTrafficsChartsTrafficInnerWithDefaults instantiates a new SavedTrafficsChartsTrafficInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrafficsChartsTrafficInnerWithDefaults() *SavedTrafficsChartsTrafficInner {
	this := SavedTrafficsChartsTrafficInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SavedTrafficsChartsTrafficInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsChartsTrafficInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SavedTrafficsChartsTrafficInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SavedTrafficsChartsTrafficInner) SetName(v string) {
	o.Name = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SavedTrafficsChartsTrafficInner) GetY() int32 {
	if o == nil || IsNil(o.Y) {
		var ret int32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsChartsTrafficInner) GetYOk() (*int32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SavedTrafficsChartsTrafficInner) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given int32 and assigns it to the Y field.
func (o *SavedTrafficsChartsTrafficInner) SetY(v int32) {
	o.Y = &v
}

func (o SavedTrafficsChartsTrafficInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrafficsChartsTrafficInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableSavedTrafficsChartsTrafficInner struct {
	value *SavedTrafficsChartsTrafficInner
	isSet bool
}

func (v NullableSavedTrafficsChartsTrafficInner) Get() *SavedTrafficsChartsTrafficInner {
	return v.value
}

func (v *NullableSavedTrafficsChartsTrafficInner) Set(val *SavedTrafficsChartsTrafficInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrafficsChartsTrafficInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrafficsChartsTrafficInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrafficsChartsTrafficInner(val *SavedTrafficsChartsTrafficInner) *NullableSavedTrafficsChartsTrafficInner {
	return &NullableSavedTrafficsChartsTrafficInner{value: val, isSet: true}
}

func (v NullableSavedTrafficsChartsTrafficInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrafficsChartsTrafficInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


