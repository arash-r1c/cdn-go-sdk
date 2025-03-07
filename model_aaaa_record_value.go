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

// checks if the AAAARecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AAAARecordValue{}

// AAAARecordValue struct for AAAARecordValue
type AAAARecordValue struct {
	Ip string `json:"ip"`
	Port NullableInt32 `json:"port,omitempty"`
	Weight NullableInt32 `json:"weight,omitempty"`
	// This key shows itself here if the weights have been changed by Health Check.
	OriginalWeight *int32 `json:"original_weight,omitempty"`
	// ISO 3166 alpha-2 country code
	Country NullableString `json:"country,omitempty"`
}

// NewAAAARecordValue instantiates a new AAAARecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAAAARecordValue(ip string) *AAAARecordValue {
	this := AAAARecordValue{}
	this.Ip = ip
	return &this
}

// NewAAAARecordValueWithDefaults instantiates a new AAAARecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAAAARecordValueWithDefaults() *AAAARecordValue {
	this := AAAARecordValue{}
	return &this
}

// GetIp returns the Ip field value
func (o *AAAARecordValue) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *AAAARecordValue) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *AAAARecordValue) SetIp(v string) {
	o.Ip = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AAAARecordValue) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AAAARecordValue) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *AAAARecordValue) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *AAAARecordValue) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *AAAARecordValue) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *AAAARecordValue) UnsetPort() {
	o.Port.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AAAARecordValue) GetWeight() int32 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret int32
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AAAARecordValue) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *AAAARecordValue) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableInt32 and assigns it to the Weight field.
func (o *AAAARecordValue) SetWeight(v int32) {
	o.Weight.Set(&v)
}
// SetWeightNil sets the value for Weight to be an explicit nil
func (o *AAAARecordValue) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *AAAARecordValue) UnsetWeight() {
	o.Weight.Unset()
}

// GetOriginalWeight returns the OriginalWeight field value if set, zero value otherwise.
func (o *AAAARecordValue) GetOriginalWeight() int32 {
	if o == nil || IsNil(o.OriginalWeight) {
		var ret int32
		return ret
	}
	return *o.OriginalWeight
}

// GetOriginalWeightOk returns a tuple with the OriginalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AAAARecordValue) GetOriginalWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalWeight) {
		return nil, false
	}
	return o.OriginalWeight, true
}

// HasOriginalWeight returns a boolean if a field has been set.
func (o *AAAARecordValue) HasOriginalWeight() bool {
	if o != nil && !IsNil(o.OriginalWeight) {
		return true
	}

	return false
}

// SetOriginalWeight gets a reference to the given int32 and assigns it to the OriginalWeight field.
func (o *AAAARecordValue) SetOriginalWeight(v int32) {
	o.OriginalWeight = &v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AAAARecordValue) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AAAARecordValue) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *AAAARecordValue) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *AAAARecordValue) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *AAAARecordValue) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *AAAARecordValue) UnsetCountry() {
	o.Country.Unset()
}

func (o AAAARecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AAAARecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	// skip: original_weight is readOnly
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	return toSerialize, nil
}

type NullableAAAARecordValue struct {
	value *AAAARecordValue
	isSet bool
}

func (v NullableAAAARecordValue) Get() *AAAARecordValue {
	return v.value
}

func (v *NullableAAAARecordValue) Set(val *AAAARecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAAAARecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAAAARecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAAAARecordValue(val *AAAARecordValue) *NullableAAAARecordValue {
	return &NullableAAAARecordValue{value: val, isSet: true}
}

func (v NullableAAAARecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAAAARecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


