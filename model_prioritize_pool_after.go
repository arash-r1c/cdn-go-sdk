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

// checks if the PrioritizePoolAfter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrioritizePoolAfter{}

// PrioritizePoolAfter struct for PrioritizePoolAfter
type PrioritizePoolAfter struct {
	// ID of the pool you want to move
	PoolId string `json:"pool_id"`
	// ID of the pool you want to be prior to the selected pool
	AfterPoolId string `json:"after_pool_id"`
}

// NewPrioritizePoolAfter instantiates a new PrioritizePoolAfter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrioritizePoolAfter(poolId string, afterPoolId string) *PrioritizePoolAfter {
	this := PrioritizePoolAfter{}
	this.PoolId = poolId
	this.AfterPoolId = afterPoolId
	return &this
}

// NewPrioritizePoolAfterWithDefaults instantiates a new PrioritizePoolAfter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrioritizePoolAfterWithDefaults() *PrioritizePoolAfter {
	this := PrioritizePoolAfter{}
	return &this
}

// GetPoolId returns the PoolId field value
func (o *PrioritizePoolAfter) GetPoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *PrioritizePoolAfter) GetPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *PrioritizePoolAfter) SetPoolId(v string) {
	o.PoolId = v
}

// GetAfterPoolId returns the AfterPoolId field value
func (o *PrioritizePoolAfter) GetAfterPoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfterPoolId
}

// GetAfterPoolIdOk returns a tuple with the AfterPoolId field value
// and a boolean to check if the value has been set.
func (o *PrioritizePoolAfter) GetAfterPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfterPoolId, true
}

// SetAfterPoolId sets field value
func (o *PrioritizePoolAfter) SetAfterPoolId(v string) {
	o.AfterPoolId = v
}

func (o PrioritizePoolAfter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrioritizePoolAfter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pool_id"] = o.PoolId
	toSerialize["after_pool_id"] = o.AfterPoolId
	return toSerialize, nil
}

type NullablePrioritizePoolAfter struct {
	value *PrioritizePoolAfter
	isSet bool
}

func (v NullablePrioritizePoolAfter) Get() *PrioritizePoolAfter {
	return v.value
}

func (v *NullablePrioritizePoolAfter) Set(val *PrioritizePoolAfter) {
	v.value = val
	v.isSet = true
}

func (v NullablePrioritizePoolAfter) IsSet() bool {
	return v.isSet
}

func (v *NullablePrioritizePoolAfter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrioritizePoolAfter(val *PrioritizePoolAfter) *NullablePrioritizePoolAfter {
	return &NullablePrioritizePoolAfter{value: val, isSet: true}
}

func (v NullablePrioritizePoolAfter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrioritizePoolAfter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


