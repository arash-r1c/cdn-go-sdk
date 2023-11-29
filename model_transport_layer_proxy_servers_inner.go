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

// checks if the TransportLayerProxyServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyServersInner{}

// TransportLayerProxyServersInner struct for TransportLayerProxyServersInner
type TransportLayerProxyServersInner struct {
	Address *string `json:"address,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Check *string `json:"check,omitempty"`
	Fall *int32 `json:"fall,omitempty"`
	Inter *int32 `json:"inter,omitempty"`
	Rise *int32 `json:"rise,omitempty"`
}

// NewTransportLayerProxyServersInner instantiates a new TransportLayerProxyServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyServersInner() *TransportLayerProxyServersInner {
	this := TransportLayerProxyServersInner{}
	return &this
}

// NewTransportLayerProxyServersInnerWithDefaults instantiates a new TransportLayerProxyServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyServersInnerWithDefaults() *TransportLayerProxyServersInner {
	this := TransportLayerProxyServersInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TransportLayerProxyServersInner) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *TransportLayerProxyServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *TransportLayerProxyServersInner) SetWeight(v int32) {
	o.Weight = &v
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetCheck() string {
	if o == nil || IsNil(o.Check) {
		var ret string
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetCheckOk() (*string, bool) {
	if o == nil || IsNil(o.Check) {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasCheck() bool {
	if o != nil && !IsNil(o.Check) {
		return true
	}

	return false
}

// SetCheck gets a reference to the given string and assigns it to the Check field.
func (o *TransportLayerProxyServersInner) SetCheck(v string) {
	o.Check = &v
}

// GetFall returns the Fall field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetFall() int32 {
	if o == nil || IsNil(o.Fall) {
		var ret int32
		return ret
	}
	return *o.Fall
}

// GetFallOk returns a tuple with the Fall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetFallOk() (*int32, bool) {
	if o == nil || IsNil(o.Fall) {
		return nil, false
	}
	return o.Fall, true
}

// HasFall returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasFall() bool {
	if o != nil && !IsNil(o.Fall) {
		return true
	}

	return false
}

// SetFall gets a reference to the given int32 and assigns it to the Fall field.
func (o *TransportLayerProxyServersInner) SetFall(v int32) {
	o.Fall = &v
}

// GetInter returns the Inter field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetInter() int32 {
	if o == nil || IsNil(o.Inter) {
		var ret int32
		return ret
	}
	return *o.Inter
}

// GetInterOk returns a tuple with the Inter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetInterOk() (*int32, bool) {
	if o == nil || IsNil(o.Inter) {
		return nil, false
	}
	return o.Inter, true
}

// HasInter returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasInter() bool {
	if o != nil && !IsNil(o.Inter) {
		return true
	}

	return false
}

// SetInter gets a reference to the given int32 and assigns it to the Inter field.
func (o *TransportLayerProxyServersInner) SetInter(v int32) {
	o.Inter = &v
}

// GetRise returns the Rise field value if set, zero value otherwise.
func (o *TransportLayerProxyServersInner) GetRise() int32 {
	if o == nil || IsNil(o.Rise) {
		var ret int32
		return ret
	}
	return *o.Rise
}

// GetRiseOk returns a tuple with the Rise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServersInner) GetRiseOk() (*int32, bool) {
	if o == nil || IsNil(o.Rise) {
		return nil, false
	}
	return o.Rise, true
}

// HasRise returns a boolean if a field has been set.
func (o *TransportLayerProxyServersInner) HasRise() bool {
	if o != nil && !IsNil(o.Rise) {
		return true
	}

	return false
}

// SetRise gets a reference to the given int32 and assigns it to the Rise field.
func (o *TransportLayerProxyServersInner) SetRise(v int32) {
	o.Rise = &v
}

func (o TransportLayerProxyServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Check) {
		toSerialize["check"] = o.Check
	}
	if !IsNil(o.Fall) {
		toSerialize["fall"] = o.Fall
	}
	if !IsNil(o.Inter) {
		toSerialize["inter"] = o.Inter
	}
	if !IsNil(o.Rise) {
		toSerialize["rise"] = o.Rise
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyServersInner struct {
	value *TransportLayerProxyServersInner
	isSet bool
}

func (v NullableTransportLayerProxyServersInner) Get() *TransportLayerProxyServersInner {
	return v.value
}

func (v *NullableTransportLayerProxyServersInner) Set(val *TransportLayerProxyServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyServersInner(val *TransportLayerProxyServersInner) *NullableTransportLayerProxyServersInner {
	return &NullableTransportLayerProxyServersInner{value: val, isSet: true}
}

func (v NullableTransportLayerProxyServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


