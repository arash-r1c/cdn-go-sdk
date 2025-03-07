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

// checks if the FeatureSets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureSets{}

// FeatureSets struct for FeatureSets
type FeatureSets struct {
	Currency *Currency `json:"currency,omitempty"`
	Plans []PlanInfo `json:"plans,omitempty"`
	FeatureSets []FeatureSet `json:"feature_sets,omitempty"`
}

// NewFeatureSets instantiates a new FeatureSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureSets() *FeatureSets {
	this := FeatureSets{}
	return &this
}

// NewFeatureSetsWithDefaults instantiates a new FeatureSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureSetsWithDefaults() *FeatureSets {
	this := FeatureSets{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *FeatureSets) GetCurrency() Currency {
	if o == nil || IsNil(o.Currency) {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSets) GetCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *FeatureSets) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *FeatureSets) SetCurrency(v Currency) {
	o.Currency = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *FeatureSets) GetPlans() []PlanInfo {
	if o == nil || IsNil(o.Plans) {
		var ret []PlanInfo
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSets) GetPlansOk() ([]PlanInfo, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *FeatureSets) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []PlanInfo and assigns it to the Plans field.
func (o *FeatureSets) SetPlans(v []PlanInfo) {
	o.Plans = v
}

// GetFeatureSets returns the FeatureSets field value if set, zero value otherwise.
func (o *FeatureSets) GetFeatureSets() []FeatureSet {
	if o == nil || IsNil(o.FeatureSets) {
		var ret []FeatureSet
		return ret
	}
	return o.FeatureSets
}

// GetFeatureSetsOk returns a tuple with the FeatureSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSets) GetFeatureSetsOk() ([]FeatureSet, bool) {
	if o == nil || IsNil(o.FeatureSets) {
		return nil, false
	}
	return o.FeatureSets, true
}

// HasFeatureSets returns a boolean if a field has been set.
func (o *FeatureSets) HasFeatureSets() bool {
	if o != nil && !IsNil(o.FeatureSets) {
		return true
	}

	return false
}

// SetFeatureSets gets a reference to the given []FeatureSet and assigns it to the FeatureSets field.
func (o *FeatureSets) SetFeatureSets(v []FeatureSet) {
	o.FeatureSets = v
}

func (o FeatureSets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureSets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !IsNil(o.FeatureSets) {
		toSerialize["feature_sets"] = o.FeatureSets
	}
	return toSerialize, nil
}

type NullableFeatureSets struct {
	value *FeatureSets
	isSet bool
}

func (v NullableFeatureSets) Get() *FeatureSets {
	return v.value
}

func (v *NullableFeatureSets) Set(val *FeatureSets) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureSets) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureSets(val *FeatureSets) *NullableFeatureSets {
	return &NullableFeatureSets{value: val, isSet: true}
}

func (v NullableFeatureSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


