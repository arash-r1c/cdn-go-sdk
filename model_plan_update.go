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

// checks if the PlanUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanUpdate{}

// PlanUpdate struct for PlanUpdate
type PlanUpdate struct {
	// - `0` - Traffic - `1` - Basic - `2` - Growth - `3` - Professional - `4` - Enterprise 
	PlanLevel int32 `json:"plan_level"`
}

// NewPlanUpdate instantiates a new PlanUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanUpdate(planLevel int32) *PlanUpdate {
	this := PlanUpdate{}
	this.PlanLevel = planLevel
	return &this
}

// NewPlanUpdateWithDefaults instantiates a new PlanUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanUpdateWithDefaults() *PlanUpdate {
	this := PlanUpdate{}
	return &this
}

// GetPlanLevel returns the PlanLevel field value
func (o *PlanUpdate) GetPlanLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlanLevel
}

// GetPlanLevelOk returns a tuple with the PlanLevel field value
// and a boolean to check if the value has been set.
func (o *PlanUpdate) GetPlanLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanLevel, true
}

// SetPlanLevel sets field value
func (o *PlanUpdate) SetPlanLevel(v int32) {
	o.PlanLevel = v
}

func (o PlanUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan_level"] = o.PlanLevel
	return toSerialize, nil
}

type NullablePlanUpdate struct {
	value *PlanUpdate
	isSet bool
}

func (v NullablePlanUpdate) Get() *PlanUpdate {
	return v.value
}

func (v *NullablePlanUpdate) Set(val *PlanUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanUpdate(val *PlanUpdate) *NullablePlanUpdate {
	return &NullablePlanUpdate{value: val, isSet: true}
}

func (v NullablePlanUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


