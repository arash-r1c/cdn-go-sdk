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

// checks if the DomainWafPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainWafPackage{}

// DomainWafPackage struct for DomainWafPackage
type DomainWafPackage struct {
	// parameters of the package
	Params map[string]interface{} `json:"params,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Provider *WafPackageProvider `json:"provider,omitempty"`
	// JSON-schema of parameters of the package
	ParamsSchema map[string]interface{} `json:"params_schema,omitempty"`
	// It will be filled by default disabled rules when it's not provided
	DisabledRules []string `json:"disabled_rules,omitempty"`
	// It will be filled by default disabled rulesets when it's not provided
	DisabledRulesets []string `json:"disabled_rulesets,omitempty"`
}

// NewDomainWafPackage instantiates a new DomainWafPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainWafPackage() *DomainWafPackage {
	this := DomainWafPackage{}
	return &this
}

// NewDomainWafPackageWithDefaults instantiates a new DomainWafPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWafPackageWithDefaults() *DomainWafPackage {
	this := DomainWafPackage{}
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	return &this
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *DomainWafPackage) GetParams() map[string]interface{} {
	if o == nil || IsNil(o.Params) {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *DomainWafPackage) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *DomainWafPackage) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *DomainWafPackage) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *DomainWafPackage) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *DomainWafPackage) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DomainWafPackage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DomainWafPackage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DomainWafPackage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DomainWafPackage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DomainWafPackage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DomainWafPackage) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *DomainWafPackage) GetProvider() WafPackageProvider {
	if o == nil || IsNil(o.Provider) {
		var ret WafPackageProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetProviderOk() (*WafPackageProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *DomainWafPackage) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given WafPackageProvider and assigns it to the Provider field.
func (o *DomainWafPackage) SetProvider(v WafPackageProvider) {
	o.Provider = &v
}

// GetParamsSchema returns the ParamsSchema field value if set, zero value otherwise.
func (o *DomainWafPackage) GetParamsSchema() map[string]interface{} {
	if o == nil || IsNil(o.ParamsSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParamsSchema
}

// GetParamsSchemaOk returns a tuple with the ParamsSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetParamsSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParamsSchema) {
		return map[string]interface{}{}, false
	}
	return o.ParamsSchema, true
}

// HasParamsSchema returns a boolean if a field has been set.
func (o *DomainWafPackage) HasParamsSchema() bool {
	if o != nil && !IsNil(o.ParamsSchema) {
		return true
	}

	return false
}

// SetParamsSchema gets a reference to the given map[string]interface{} and assigns it to the ParamsSchema field.
func (o *DomainWafPackage) SetParamsSchema(v map[string]interface{}) {
	o.ParamsSchema = v
}

// GetDisabledRules returns the DisabledRules field value if set, zero value otherwise.
func (o *DomainWafPackage) GetDisabledRules() []string {
	if o == nil || IsNil(o.DisabledRules) {
		var ret []string
		return ret
	}
	return o.DisabledRules
}

// GetDisabledRulesOk returns a tuple with the DisabledRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetDisabledRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledRules) {
		return nil, false
	}
	return o.DisabledRules, true
}

// HasDisabledRules returns a boolean if a field has been set.
func (o *DomainWafPackage) HasDisabledRules() bool {
	if o != nil && !IsNil(o.DisabledRules) {
		return true
	}

	return false
}

// SetDisabledRules gets a reference to the given []string and assigns it to the DisabledRules field.
func (o *DomainWafPackage) SetDisabledRules(v []string) {
	o.DisabledRules = v
}

// GetDisabledRulesets returns the DisabledRulesets field value if set, zero value otherwise.
func (o *DomainWafPackage) GetDisabledRulesets() []string {
	if o == nil || IsNil(o.DisabledRulesets) {
		var ret []string
		return ret
	}
	return o.DisabledRulesets
}

// GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackage) GetDisabledRulesetsOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledRulesets) {
		return nil, false
	}
	return o.DisabledRulesets, true
}

// HasDisabledRulesets returns a boolean if a field has been set.
func (o *DomainWafPackage) HasDisabledRulesets() bool {
	if o != nil && !IsNil(o.DisabledRulesets) {
		return true
	}

	return false
}

// SetDisabledRulesets gets a reference to the given []string and assigns it to the DisabledRulesets field.
func (o *DomainWafPackage) SetDisabledRulesets(v []string) {
	o.DisabledRulesets = v
}

func (o DomainWafPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainWafPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	// skip: id is readOnly
	// skip: name is readOnly
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.ParamsSchema) {
		toSerialize["params_schema"] = o.ParamsSchema
	}
	if !IsNil(o.DisabledRules) {
		toSerialize["disabled_rules"] = o.DisabledRules
	}
	if !IsNil(o.DisabledRulesets) {
		toSerialize["disabled_rulesets"] = o.DisabledRulesets
	}
	return toSerialize, nil
}

type NullableDomainWafPackage struct {
	value *DomainWafPackage
	isSet bool
}

func (v NullableDomainWafPackage) Get() *DomainWafPackage {
	return v.value
}

func (v *NullableDomainWafPackage) Set(val *DomainWafPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainWafPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainWafPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainWafPackage(val *DomainWafPackage) *NullableDomainWafPackage {
	return &NullableDomainWafPackage{value: val, isSet: true}
}

func (v NullableDomainWafPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainWafPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


