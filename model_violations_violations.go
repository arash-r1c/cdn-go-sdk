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

// checks if the ViolationsViolations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViolationsViolations{}

// ViolationsViolations struct for ViolationsViolations
type ViolationsViolations struct {
	Paygo []string `json:"paygo,omitempty"`
	Basic []string `json:"basic,omitempty"`
	Growth []string `json:"growth,omitempty"`
	Professional []string `json:"professional,omitempty"`
	Enterprise []string `json:"enterprise,omitempty"`
}

// NewViolationsViolations instantiates a new ViolationsViolations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolationsViolations() *ViolationsViolations {
	this := ViolationsViolations{}
	return &this
}

// NewViolationsViolationsWithDefaults instantiates a new ViolationsViolations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationsViolationsWithDefaults() *ViolationsViolations {
	this := ViolationsViolations{}
	return &this
}

// GetPaygo returns the Paygo field value if set, zero value otherwise.
func (o *ViolationsViolations) GetPaygo() []string {
	if o == nil || IsNil(o.Paygo) {
		var ret []string
		return ret
	}
	return o.Paygo
}

// GetPaygoOk returns a tuple with the Paygo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationsViolations) GetPaygoOk() ([]string, bool) {
	if o == nil || IsNil(o.Paygo) {
		return nil, false
	}
	return o.Paygo, true
}

// HasPaygo returns a boolean if a field has been set.
func (o *ViolationsViolations) HasPaygo() bool {
	if o != nil && !IsNil(o.Paygo) {
		return true
	}

	return false
}

// SetPaygo gets a reference to the given []string and assigns it to the Paygo field.
func (o *ViolationsViolations) SetPaygo(v []string) {
	o.Paygo = v
}

// GetBasic returns the Basic field value if set, zero value otherwise.
func (o *ViolationsViolations) GetBasic() []string {
	if o == nil || IsNil(o.Basic) {
		var ret []string
		return ret
	}
	return o.Basic
}

// GetBasicOk returns a tuple with the Basic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationsViolations) GetBasicOk() ([]string, bool) {
	if o == nil || IsNil(o.Basic) {
		return nil, false
	}
	return o.Basic, true
}

// HasBasic returns a boolean if a field has been set.
func (o *ViolationsViolations) HasBasic() bool {
	if o != nil && !IsNil(o.Basic) {
		return true
	}

	return false
}

// SetBasic gets a reference to the given []string and assigns it to the Basic field.
func (o *ViolationsViolations) SetBasic(v []string) {
	o.Basic = v
}

// GetGrowth returns the Growth field value if set, zero value otherwise.
func (o *ViolationsViolations) GetGrowth() []string {
	if o == nil || IsNil(o.Growth) {
		var ret []string
		return ret
	}
	return o.Growth
}

// GetGrowthOk returns a tuple with the Growth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationsViolations) GetGrowthOk() ([]string, bool) {
	if o == nil || IsNil(o.Growth) {
		return nil, false
	}
	return o.Growth, true
}

// HasGrowth returns a boolean if a field has been set.
func (o *ViolationsViolations) HasGrowth() bool {
	if o != nil && !IsNil(o.Growth) {
		return true
	}

	return false
}

// SetGrowth gets a reference to the given []string and assigns it to the Growth field.
func (o *ViolationsViolations) SetGrowth(v []string) {
	o.Growth = v
}

// GetProfessional returns the Professional field value if set, zero value otherwise.
func (o *ViolationsViolations) GetProfessional() []string {
	if o == nil || IsNil(o.Professional) {
		var ret []string
		return ret
	}
	return o.Professional
}

// GetProfessionalOk returns a tuple with the Professional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationsViolations) GetProfessionalOk() ([]string, bool) {
	if o == nil || IsNil(o.Professional) {
		return nil, false
	}
	return o.Professional, true
}

// HasProfessional returns a boolean if a field has been set.
func (o *ViolationsViolations) HasProfessional() bool {
	if o != nil && !IsNil(o.Professional) {
		return true
	}

	return false
}

// SetProfessional gets a reference to the given []string and assigns it to the Professional field.
func (o *ViolationsViolations) SetProfessional(v []string) {
	o.Professional = v
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise.
func (o *ViolationsViolations) GetEnterprise() []string {
	if o == nil || IsNil(o.Enterprise) {
		var ret []string
		return ret
	}
	return o.Enterprise
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViolationsViolations) GetEnterpriseOk() ([]string, bool) {
	if o == nil || IsNil(o.Enterprise) {
		return nil, false
	}
	return o.Enterprise, true
}

// HasEnterprise returns a boolean if a field has been set.
func (o *ViolationsViolations) HasEnterprise() bool {
	if o != nil && !IsNil(o.Enterprise) {
		return true
	}

	return false
}

// SetEnterprise gets a reference to the given []string and assigns it to the Enterprise field.
func (o *ViolationsViolations) SetEnterprise(v []string) {
	o.Enterprise = v
}

func (o ViolationsViolations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViolationsViolations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paygo) {
		toSerialize["paygo"] = o.Paygo
	}
	if !IsNil(o.Basic) {
		toSerialize["basic"] = o.Basic
	}
	if !IsNil(o.Growth) {
		toSerialize["growth"] = o.Growth
	}
	if !IsNil(o.Professional) {
		toSerialize["professional"] = o.Professional
	}
	if !IsNil(o.Enterprise) {
		toSerialize["enterprise"] = o.Enterprise
	}
	return toSerialize, nil
}

type NullableViolationsViolations struct {
	value *ViolationsViolations
	isSet bool
}

func (v NullableViolationsViolations) Get() *ViolationsViolations {
	return v.value
}

func (v *NullableViolationsViolations) Set(val *ViolationsViolations) {
	v.value = val
	v.isSet = true
}

func (v NullableViolationsViolations) IsSet() bool {
	return v.isSet
}

func (v *NullableViolationsViolations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolationsViolations(val *ViolationsViolations) *NullableViolationsViolations {
	return &NullableViolationsViolations{value: val, isSet: true}
}

func (v NullableViolationsViolations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolationsViolations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


