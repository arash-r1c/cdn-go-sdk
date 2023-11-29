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

// checks if the WafSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafSettings{}

// WafSettings struct for WafSettings
type WafSettings struct {
	IsEnabled *bool `json:"is_enabled,omitempty"`
	Mode *string `json:"mode,omitempty"`
	// Pacakges and their configurations that are used to configure WAF.
	Packages []DomainWafPackage `json:"packages,omitempty"`
}

// NewWafSettings instantiates a new WafSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafSettings() *WafSettings {
	this := WafSettings{}
	return &this
}

// NewWafSettingsWithDefaults instantiates a new WafSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafSettingsWithDefaults() *WafSettings {
	this := WafSettings{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *WafSettings) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSettings) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *WafSettings) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *WafSettings) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WafSettings) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSettings) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WafSettings) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *WafSettings) SetMode(v string) {
	o.Mode = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *WafSettings) GetPackages() []DomainWafPackage {
	if o == nil || IsNil(o.Packages) {
		var ret []DomainWafPackage
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafSettings) GetPackagesOk() ([]DomainWafPackage, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *WafSettings) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []DomainWafPackage and assigns it to the Packages field.
func (o *WafSettings) SetPackages(v []DomainWafPackage) {
	o.Packages = v
}

func (o WafSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: is_enabled is readOnly
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	// skip: packages is readOnly
	return toSerialize, nil
}

type NullableWafSettings struct {
	value *WafSettings
	isSet bool
}

func (v NullableWafSettings) Get() *WafSettings {
	return v.value
}

func (v *NullableWafSettings) Set(val *WafSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWafSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWafSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafSettings(val *WafSettings) *NullableWafSettings {
	return &NullableWafSettings{value: val, isSet: true}
}

func (v NullableWafSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


