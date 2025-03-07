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

// checks if the FirewallRuleView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRuleView{}

// FirewallRuleView struct for FirewallRuleView
type FirewallRuleView struct {
	ActionDetails map[string]interface{} `json:"action_details,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Wireshark-like filter expression
	FilterExpr *string `json:"filter_expr,omitempty"`
	Action *string `json:"action,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	Note *string `json:"note,omitempty"`
}

// NewFirewallRuleView instantiates a new FirewallRuleView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRuleView() *FirewallRuleView {
	this := FirewallRuleView{}
	return &this
}

// NewFirewallRuleViewWithDefaults instantiates a new FirewallRuleView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleViewWithDefaults() *FirewallRuleView {
	this := FirewallRuleView{}
	return &this
}

// GetActionDetails returns the ActionDetails field value if set, zero value otherwise.
func (o *FirewallRuleView) GetActionDetails() map[string]interface{} {
	if o == nil || IsNil(o.ActionDetails) {
		var ret map[string]interface{}
		return ret
	}
	return o.ActionDetails
}

// GetActionDetailsOk returns a tuple with the ActionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetActionDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ActionDetails) {
		return map[string]interface{}{}, false
	}
	return o.ActionDetails, true
}

// HasActionDetails returns a boolean if a field has been set.
func (o *FirewallRuleView) HasActionDetails() bool {
	if o != nil && !IsNil(o.ActionDetails) {
		return true
	}

	return false
}

// SetActionDetails gets a reference to the given map[string]interface{} and assigns it to the ActionDetails field.
func (o *FirewallRuleView) SetActionDetails(v map[string]interface{}) {
	o.ActionDetails = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirewallRuleView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirewallRuleView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FirewallRuleView) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirewallRuleView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirewallRuleView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirewallRuleView) SetName(v string) {
	o.Name = &v
}

// GetFilterExpr returns the FilterExpr field value if set, zero value otherwise.
func (o *FirewallRuleView) GetFilterExpr() string {
	if o == nil || IsNil(o.FilterExpr) {
		var ret string
		return ret
	}
	return *o.FilterExpr
}

// GetFilterExprOk returns a tuple with the FilterExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetFilterExprOk() (*string, bool) {
	if o == nil || IsNil(o.FilterExpr) {
		return nil, false
	}
	return o.FilterExpr, true
}

// HasFilterExpr returns a boolean if a field has been set.
func (o *FirewallRuleView) HasFilterExpr() bool {
	if o != nil && !IsNil(o.FilterExpr) {
		return true
	}

	return false
}

// SetFilterExpr gets a reference to the given string and assigns it to the FilterExpr field.
func (o *FirewallRuleView) SetFilterExpr(v string) {
	o.FilterExpr = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *FirewallRuleView) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *FirewallRuleView) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *FirewallRuleView) SetAction(v string) {
	o.Action = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *FirewallRuleView) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *FirewallRuleView) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *FirewallRuleView) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *FirewallRuleView) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRuleView) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *FirewallRuleView) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *FirewallRuleView) SetNote(v string) {
	o.Note = &v
}

func (o FirewallRuleView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRuleView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionDetails) {
		toSerialize["action_details"] = o.ActionDetails
	}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FilterExpr) {
		toSerialize["filter_expr"] = o.FilterExpr
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

type NullableFirewallRuleView struct {
	value *FirewallRuleView
	isSet bool
}

func (v NullableFirewallRuleView) Get() *FirewallRuleView {
	return v.value
}

func (v *NullableFirewallRuleView) Set(val *FirewallRuleView) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRuleView) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRuleView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRuleView(val *FirewallRuleView) *NullableFirewallRuleView {
	return &NullableFirewallRuleView{value: val, isSet: true}
}

func (v NullableFirewallRuleView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRuleView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


