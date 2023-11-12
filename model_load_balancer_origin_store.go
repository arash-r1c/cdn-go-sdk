/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.114.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the LoadBalancerOriginStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerOriginStore{}

// LoadBalancerOriginStore struct for LoadBalancerOriginStore
type LoadBalancerOriginStore struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status bool `json:"status"`
	Address string `json:"address"`
	Port int32 `json:"port"`
	Weight int32 `json:"weight"`
	Protocol string `json:"protocol"`
	HostHeader *string `json:"host_header,omitempty"`
}

// NewLoadBalancerOriginStore instantiates a new LoadBalancerOriginStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerOriginStore(status bool, address string, port int32, weight int32, protocol string) *LoadBalancerOriginStore {
	this := LoadBalancerOriginStore{}
	this.Status = status
	this.Address = address
	this.Port = port
	this.Weight = weight
	this.Protocol = protocol
	return &this
}

// NewLoadBalancerOriginStoreWithDefaults instantiates a new LoadBalancerOriginStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerOriginStoreWithDefaults() *LoadBalancerOriginStore {
	this := LoadBalancerOriginStore{}
	var protocol string = "auto"
	this.Protocol = protocol
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancerOriginStore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadBalancerOriginStore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadBalancerOriginStore) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerOriginStore) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerOriginStore) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerOriginStore) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value
func (o *LoadBalancerOriginStore) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LoadBalancerOriginStore) SetStatus(v bool) {
	o.Status = v
}

// GetAddress returns the Address field value
func (o *LoadBalancerOriginStore) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *LoadBalancerOriginStore) SetAddress(v string) {
	o.Address = v
}

// GetPort returns the Port field value
func (o *LoadBalancerOriginStore) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *LoadBalancerOriginStore) SetPort(v int32) {
	o.Port = v
}

// GetWeight returns the Weight field value
func (o *LoadBalancerOriginStore) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *LoadBalancerOriginStore) SetWeight(v int32) {
	o.Weight = v
}

// GetProtocol returns the Protocol field value
func (o *LoadBalancerOriginStore) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *LoadBalancerOriginStore) SetProtocol(v string) {
	o.Protocol = v
}

// GetHostHeader returns the HostHeader field value if set, zero value otherwise.
func (o *LoadBalancerOriginStore) GetHostHeader() string {
	if o == nil || IsNil(o.HostHeader) {
		var ret string
		return ret
	}
	return *o.HostHeader
}

// GetHostHeaderOk returns a tuple with the HostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginStore) GetHostHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HostHeader) {
		return nil, false
	}
	return o.HostHeader, true
}

// HasHostHeader returns a boolean if a field has been set.
func (o *LoadBalancerOriginStore) HasHostHeader() bool {
	if o != nil && !IsNil(o.HostHeader) {
		return true
	}

	return false
}

// SetHostHeader gets a reference to the given string and assigns it to the HostHeader field.
func (o *LoadBalancerOriginStore) SetHostHeader(v string) {
	o.HostHeader = &v
}

func (o LoadBalancerOriginStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerOriginStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["status"] = o.Status
	toSerialize["address"] = o.Address
	toSerialize["port"] = o.Port
	toSerialize["weight"] = o.Weight
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.HostHeader) {
		toSerialize["host_header"] = o.HostHeader
	}
	return toSerialize, nil
}

type NullableLoadBalancerOriginStore struct {
	value *LoadBalancerOriginStore
	isSet bool
}

func (v NullableLoadBalancerOriginStore) Get() *LoadBalancerOriginStore {
	return v.value
}

func (v *NullableLoadBalancerOriginStore) Set(val *LoadBalancerOriginStore) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerOriginStore) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerOriginStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerOriginStore(val *LoadBalancerOriginStore) *NullableLoadBalancerOriginStore {
	return &NullableLoadBalancerOriginStore{value: val, isSet: true}
}

func (v NullableLoadBalancerOriginStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerOriginStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


