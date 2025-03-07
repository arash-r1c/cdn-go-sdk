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

// checks if the TransportLayerProxyServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyServer{}

// TransportLayerProxyServer struct for TransportLayerProxyServer
type TransportLayerProxyServer struct {
	Address string `json:"address"`
	Port int32 `json:"port"`
	Weight int32 `json:"weight"`
	Check string `json:"check"`
	Fall *int32 `json:"fall,omitempty"`
	Inter *int32 `json:"inter,omitempty"`
	Rise *int32 `json:"rise,omitempty"`
}

// NewTransportLayerProxyServer instantiates a new TransportLayerProxyServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyServer(address string, port int32, weight int32, check string) *TransportLayerProxyServer {
	this := TransportLayerProxyServer{}
	this.Address = address
	this.Port = port
	this.Weight = weight
	this.Check = check
	return &this
}

// NewTransportLayerProxyServerWithDefaults instantiates a new TransportLayerProxyServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyServerWithDefaults() *TransportLayerProxyServer {
	this := TransportLayerProxyServer{}
	return &this
}

// GetAddress returns the Address field value
func (o *TransportLayerProxyServer) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *TransportLayerProxyServer) SetAddress(v string) {
	o.Address = v
}

// GetPort returns the Port field value
func (o *TransportLayerProxyServer) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *TransportLayerProxyServer) SetPort(v int32) {
	o.Port = v
}

// GetWeight returns the Weight field value
func (o *TransportLayerProxyServer) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *TransportLayerProxyServer) SetWeight(v int32) {
	o.Weight = v
}

// GetCheck returns the Check field value
func (o *TransportLayerProxyServer) GetCheck() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Check
}

// GetCheckOk returns a tuple with the Check field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetCheckOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Check, true
}

// SetCheck sets field value
func (o *TransportLayerProxyServer) SetCheck(v string) {
	o.Check = v
}

// GetFall returns the Fall field value if set, zero value otherwise.
func (o *TransportLayerProxyServer) GetFall() int32 {
	if o == nil || IsNil(o.Fall) {
		var ret int32
		return ret
	}
	return *o.Fall
}

// GetFallOk returns a tuple with the Fall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetFallOk() (*int32, bool) {
	if o == nil || IsNil(o.Fall) {
		return nil, false
	}
	return o.Fall, true
}

// HasFall returns a boolean if a field has been set.
func (o *TransportLayerProxyServer) HasFall() bool {
	if o != nil && !IsNil(o.Fall) {
		return true
	}

	return false
}

// SetFall gets a reference to the given int32 and assigns it to the Fall field.
func (o *TransportLayerProxyServer) SetFall(v int32) {
	o.Fall = &v
}

// GetInter returns the Inter field value if set, zero value otherwise.
func (o *TransportLayerProxyServer) GetInter() int32 {
	if o == nil || IsNil(o.Inter) {
		var ret int32
		return ret
	}
	return *o.Inter
}

// GetInterOk returns a tuple with the Inter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetInterOk() (*int32, bool) {
	if o == nil || IsNil(o.Inter) {
		return nil, false
	}
	return o.Inter, true
}

// HasInter returns a boolean if a field has been set.
func (o *TransportLayerProxyServer) HasInter() bool {
	if o != nil && !IsNil(o.Inter) {
		return true
	}

	return false
}

// SetInter gets a reference to the given int32 and assigns it to the Inter field.
func (o *TransportLayerProxyServer) SetInter(v int32) {
	o.Inter = &v
}

// GetRise returns the Rise field value if set, zero value otherwise.
func (o *TransportLayerProxyServer) GetRise() int32 {
	if o == nil || IsNil(o.Rise) {
		var ret int32
		return ret
	}
	return *o.Rise
}

// GetRiseOk returns a tuple with the Rise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyServer) GetRiseOk() (*int32, bool) {
	if o == nil || IsNil(o.Rise) {
		return nil, false
	}
	return o.Rise, true
}

// HasRise returns a boolean if a field has been set.
func (o *TransportLayerProxyServer) HasRise() bool {
	if o != nil && !IsNil(o.Rise) {
		return true
	}

	return false
}

// SetRise gets a reference to the given int32 and assigns it to the Rise field.
func (o *TransportLayerProxyServer) SetRise(v int32) {
	o.Rise = &v
}

func (o TransportLayerProxyServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["port"] = o.Port
	toSerialize["weight"] = o.Weight
	toSerialize["check"] = o.Check
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

type NullableTransportLayerProxyServer struct {
	value *TransportLayerProxyServer
	isSet bool
}

func (v NullableTransportLayerProxyServer) Get() *TransportLayerProxyServer {
	return v.value
}

func (v *NullableTransportLayerProxyServer) Set(val *TransportLayerProxyServer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyServer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyServer(val *TransportLayerProxyServer) *NullableTransportLayerProxyServer {
	return &NullableTransportLayerProxyServer{value: val, isSet: true}
}

func (v NullableTransportLayerProxyServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


