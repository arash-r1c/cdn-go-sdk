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

// checks if the CNAMERecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CNAMERecordValue{}

// CNAMERecordValue struct for CNAMERecordValue
type CNAMERecordValue struct {
	// A fully-qualified domain name (FQDN)
	Host string `json:"host"`
	HostHeader NullableString `json:"host_header"`
	Port NullableInt32 `json:"port,omitempty"`
}

// NewCNAMERecordValue instantiates a new CNAMERecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCNAMERecordValue(host string, hostHeader NullableString) *CNAMERecordValue {
	this := CNAMERecordValue{}
	this.Host = host
	this.HostHeader = hostHeader
	return &this
}

// NewCNAMERecordValueWithDefaults instantiates a new CNAMERecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCNAMERecordValueWithDefaults() *CNAMERecordValue {
	this := CNAMERecordValue{}
	return &this
}

// GetHost returns the Host field value
func (o *CNAMERecordValue) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CNAMERecordValue) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CNAMERecordValue) SetHost(v string) {
	o.Host = v
}

// GetHostHeader returns the HostHeader field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CNAMERecordValue) GetHostHeader() string {
	if o == nil || o.HostHeader.Get() == nil {
		var ret string
		return ret
	}

	return *o.HostHeader.Get()
}

// GetHostHeaderOk returns a tuple with the HostHeader field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CNAMERecordValue) GetHostHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostHeader.Get(), o.HostHeader.IsSet()
}

// SetHostHeader sets field value
func (o *CNAMERecordValue) SetHostHeader(v string) {
	o.HostHeader.Set(&v)
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CNAMERecordValue) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CNAMERecordValue) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *CNAMERecordValue) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *CNAMERecordValue) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *CNAMERecordValue) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *CNAMERecordValue) UnsetPort() {
	o.Port.Unset()
}

func (o CNAMERecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CNAMERecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["host_header"] = o.HostHeader.Get()
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	return toSerialize, nil
}

type NullableCNAMERecordValue struct {
	value *CNAMERecordValue
	isSet bool
}

func (v NullableCNAMERecordValue) Get() *CNAMERecordValue {
	return v.value
}

func (v *NullableCNAMERecordValue) Set(val *CNAMERecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCNAMERecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCNAMERecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCNAMERecordValue(val *CNAMERecordValue) *NullableCNAMERecordValue {
	return &NullableCNAMERecordValue{value: val, isSet: true}
}

func (v NullableCNAMERecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCNAMERecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


