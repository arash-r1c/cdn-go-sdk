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

// checks if the AttackReportUri type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportUri{}

// AttackReportUri struct for AttackReportUri
type AttackReportUri struct {
	Uri *string `json:"uri,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewAttackReportUri instantiates a new AttackReportUri object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportUri() *AttackReportUri {
	this := AttackReportUri{}
	return &this
}

// NewAttackReportUriWithDefaults instantiates a new AttackReportUri object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportUriWithDefaults() *AttackReportUri {
	this := AttackReportUri{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AttackReportUri) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportUri) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AttackReportUri) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AttackReportUri) SetUri(v string) {
	o.Uri = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AttackReportUri) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportUri) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AttackReportUri) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AttackReportUri) SetCount(v int32) {
	o.Count = &v
}

func (o AttackReportUri) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportUri) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableAttackReportUri struct {
	value *AttackReportUri
	isSet bool
}

func (v NullableAttackReportUri) Get() *AttackReportUri {
	return v.value
}

func (v *NullableAttackReportUri) Set(val *AttackReportUri) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportUri) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportUri) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportUri(val *AttackReportUri) *NullableAttackReportUri {
	return &NullableAttackReportUri{value: val, isSet: true}
}

func (v NullableAttackReportUri) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportUri) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


