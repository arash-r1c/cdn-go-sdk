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

// checks if the TcpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TcpConfig{}

// TcpConfig struct for TcpConfig
type TcpConfig struct {
	Port int32 `json:"port"`
	// In milliseconds
	Timeout int32 `json:"timeout"`
}

// NewTcpConfig instantiates a new TcpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTcpConfig(port int32, timeout int32) *TcpConfig {
	this := TcpConfig{}
	this.Port = port
	this.Timeout = timeout
	return &this
}

// NewTcpConfigWithDefaults instantiates a new TcpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTcpConfigWithDefaults() *TcpConfig {
	this := TcpConfig{}
	return &this
}

// GetPort returns the Port field value
func (o *TcpConfig) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *TcpConfig) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *TcpConfig) SetPort(v int32) {
	o.Port = v
}

// GetTimeout returns the Timeout field value
func (o *TcpConfig) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *TcpConfig) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *TcpConfig) SetTimeout(v int32) {
	o.Timeout = v
}

func (o TcpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TcpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	toSerialize["timeout"] = o.Timeout
	return toSerialize, nil
}

type NullableTcpConfig struct {
	value *TcpConfig
	isSet bool
}

func (v NullableTcpConfig) Get() *TcpConfig {
	return v.value
}

func (v *NullableTcpConfig) Set(val *TcpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTcpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTcpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTcpConfig(val *TcpConfig) *NullableTcpConfig {
	return &NullableTcpConfig{value: val, isSet: true}
}

func (v NullableTcpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTcpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


