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

// checks if the DynamicFieldValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicFieldValue{}

// DynamicFieldValue struct for DynamicFieldValue
type DynamicFieldValue struct {
	Value *DynamicFieldType `json:"value,omitempty"`
	Desc *string `json:"desc,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewDynamicFieldValue instantiates a new DynamicFieldValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicFieldValue() *DynamicFieldValue {
	this := DynamicFieldValue{}
	return &this
}

// NewDynamicFieldValueWithDefaults instantiates a new DynamicFieldValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicFieldValueWithDefaults() *DynamicFieldValue {
	this := DynamicFieldValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DynamicFieldValue) GetValue() DynamicFieldType {
	if o == nil || IsNil(o.Value) {
		var ret DynamicFieldType
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicFieldValue) GetValueOk() (*DynamicFieldType, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DynamicFieldValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given DynamicFieldType and assigns it to the Value field.
func (o *DynamicFieldValue) SetValue(v DynamicFieldType) {
	o.Value = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *DynamicFieldValue) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicFieldValue) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *DynamicFieldValue) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *DynamicFieldValue) SetDesc(v string) {
	o.Desc = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DynamicFieldValue) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicFieldValue) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DynamicFieldValue) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DynamicFieldValue) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o DynamicFieldValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicFieldValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	// skip: created_at is readOnly
	return toSerialize, nil
}

type NullableDynamicFieldValue struct {
	value *DynamicFieldValue
	isSet bool
}

func (v NullableDynamicFieldValue) Get() *DynamicFieldValue {
	return v.value
}

func (v *NullableDynamicFieldValue) Set(val *DynamicFieldValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicFieldValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicFieldValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicFieldValue(val *DynamicFieldValue) *NullableDynamicFieldValue {
	return &NullableDynamicFieldValue{value: val, isSet: true}
}

func (v NullableDynamicFieldValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicFieldValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


