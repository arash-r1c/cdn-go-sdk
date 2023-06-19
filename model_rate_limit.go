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

// checks if the RateLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimit{}

// RateLimit struct for RateLimit
type RateLimit struct {
	DdosDetection *bool `json:"ddos_detection,omitempty"`
	ExcludeSources []string `json:"exclude_sources,omitempty"`
	Rules []RateLimitRuleView `json:"rules,omitempty"`
}

// NewRateLimit instantiates a new RateLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimit() *RateLimit {
	this := RateLimit{}
	return &this
}

// NewRateLimitWithDefaults instantiates a new RateLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitWithDefaults() *RateLimit {
	this := RateLimit{}
	return &this
}

// GetDdosDetection returns the DdosDetection field value if set, zero value otherwise.
func (o *RateLimit) GetDdosDetection() bool {
	if o == nil || IsNil(o.DdosDetection) {
		var ret bool
		return ret
	}
	return *o.DdosDetection
}

// GetDdosDetectionOk returns a tuple with the DdosDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetDdosDetectionOk() (*bool, bool) {
	if o == nil || IsNil(o.DdosDetection) {
		return nil, false
	}
	return o.DdosDetection, true
}

// HasDdosDetection returns a boolean if a field has been set.
func (o *RateLimit) HasDdosDetection() bool {
	if o != nil && !IsNil(o.DdosDetection) {
		return true
	}

	return false
}

// SetDdosDetection gets a reference to the given bool and assigns it to the DdosDetection field.
func (o *RateLimit) SetDdosDetection(v bool) {
	o.DdosDetection = &v
}

// GetExcludeSources returns the ExcludeSources field value if set, zero value otherwise.
func (o *RateLimit) GetExcludeSources() []string {
	if o == nil || IsNil(o.ExcludeSources) {
		var ret []string
		return ret
	}
	return o.ExcludeSources
}

// GetExcludeSourcesOk returns a tuple with the ExcludeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetExcludeSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeSources) {
		return nil, false
	}
	return o.ExcludeSources, true
}

// HasExcludeSources returns a boolean if a field has been set.
func (o *RateLimit) HasExcludeSources() bool {
	if o != nil && !IsNil(o.ExcludeSources) {
		return true
	}

	return false
}

// SetExcludeSources gets a reference to the given []string and assigns it to the ExcludeSources field.
func (o *RateLimit) SetExcludeSources(v []string) {
	o.ExcludeSources = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *RateLimit) GetRules() []RateLimitRuleView {
	if o == nil || IsNil(o.Rules) {
		var ret []RateLimitRuleView
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetRulesOk() ([]RateLimitRuleView, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *RateLimit) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []RateLimitRuleView and assigns it to the Rules field.
func (o *RateLimit) SetRules(v []RateLimitRuleView) {
	o.Rules = v
}

func (o RateLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DdosDetection) {
		toSerialize["ddos_detection"] = o.DdosDetection
	}
	if !IsNil(o.ExcludeSources) {
		toSerialize["exclude_sources"] = o.ExcludeSources
	}
	// skip: rules is readOnly
	return toSerialize, nil
}

type NullableRateLimit struct {
	value *RateLimit
	isSet bool
}

func (v NullableRateLimit) Get() *RateLimit {
	return v.value
}

func (v *NullableRateLimit) Set(val *RateLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimit(val *RateLimit) *NullableRateLimit {
	return &NullableRateLimit{value: val, isSet: true}
}

func (v NullableRateLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


