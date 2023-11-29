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

// checks if the ANAMERecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ANAMERecordValue{}

// ANAMERecordValue struct for ANAMERecordValue
type ANAMERecordValue struct {
	// A fully-qualified domain name (FQDN)
	Location string `json:"location"`
	HostHeader NullableString `json:"host_header"`
	Port NullableInt32 `json:"port,omitempty"`
}

// NewANAMERecordValue instantiates a new ANAMERecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewANAMERecordValue(location string, hostHeader NullableString) *ANAMERecordValue {
	this := ANAMERecordValue{}
	this.Location = location
	this.HostHeader = hostHeader
	return &this
}

// NewANAMERecordValueWithDefaults instantiates a new ANAMERecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewANAMERecordValueWithDefaults() *ANAMERecordValue {
	this := ANAMERecordValue{}
	return &this
}

// GetLocation returns the Location field value
func (o *ANAMERecordValue) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ANAMERecordValue) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ANAMERecordValue) SetLocation(v string) {
	o.Location = v
}

// GetHostHeader returns the HostHeader field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ANAMERecordValue) GetHostHeader() string {
	if o == nil || o.HostHeader.Get() == nil {
		var ret string
		return ret
	}

	return *o.HostHeader.Get()
}

// GetHostHeaderOk returns a tuple with the HostHeader field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ANAMERecordValue) GetHostHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostHeader.Get(), o.HostHeader.IsSet()
}

// SetHostHeader sets field value
func (o *ANAMERecordValue) SetHostHeader(v string) {
	o.HostHeader.Set(&v)
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ANAMERecordValue) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ANAMERecordValue) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *ANAMERecordValue) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *ANAMERecordValue) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *ANAMERecordValue) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *ANAMERecordValue) UnsetPort() {
	o.Port.Unset()
}

func (o ANAMERecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ANAMERecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["host_header"] = o.HostHeader.Get()
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	return toSerialize, nil
}

type NullableANAMERecordValue struct {
	value *ANAMERecordValue
	isSet bool
}

func (v NullableANAMERecordValue) Get() *ANAMERecordValue {
	return v.value
}

func (v *NullableANAMERecordValue) Set(val *ANAMERecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableANAMERecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableANAMERecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableANAMERecordValue(val *ANAMERecordValue) *NullableANAMERecordValue {
	return &NullableANAMERecordValue{value: val, isSet: true}
}

func (v NullableANAMERecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableANAMERecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


