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

// checks if the WafRuleset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafRuleset{}

// WafRuleset struct for WafRuleset
type WafRuleset struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Rules []WafRulesetRulesInner `json:"rules,omitempty"`
}

// NewWafRuleset instantiates a new WafRuleset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleset() *WafRuleset {
	this := WafRuleset{}
	return &this
}

// NewWafRulesetWithDefaults instantiates a new WafRuleset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRulesetWithDefaults() *WafRuleset {
	this := WafRuleset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafRuleset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafRuleset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafRuleset) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafRuleset) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleset) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafRuleset) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafRuleset) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *WafRuleset) GetRules() []WafRulesetRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []WafRulesetRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleset) GetRulesOk() ([]WafRulesetRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *WafRuleset) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []WafRulesetRulesInner and assigns it to the Rules field.
func (o *WafRuleset) SetRules(v []WafRulesetRulesInner) {
	o.Rules = v
}

func (o WafRuleset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafRuleset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableWafRuleset struct {
	value *WafRuleset
	isSet bool
}

func (v NullableWafRuleset) Get() *WafRuleset {
	return v.value
}

func (v *NullableWafRuleset) Set(val *WafRuleset) {
	v.value = val
	v.isSet = true
}

func (v NullableWafRuleset) IsSet() bool {
	return v.isSet
}

func (v *NullableWafRuleset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafRuleset(val *WafRuleset) *NullableWafRuleset {
	return &NullableWafRuleset{value: val, isSet: true}
}

func (v NullableWafRuleset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafRuleset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


