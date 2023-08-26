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

// checks if the LoadBalancersRegionsIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancersRegionsIndex200Response{}

// LoadBalancersRegionsIndex200Response struct for LoadBalancersRegionsIndex200Response
type LoadBalancersRegionsIndex200Response struct {
	Data []LoadBalancerRegion `json:"data,omitempty"`
}

// NewLoadBalancersRegionsIndex200Response instantiates a new LoadBalancersRegionsIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancersRegionsIndex200Response() *LoadBalancersRegionsIndex200Response {
	this := LoadBalancersRegionsIndex200Response{}
	return &this
}

// NewLoadBalancersRegionsIndex200ResponseWithDefaults instantiates a new LoadBalancersRegionsIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancersRegionsIndex200ResponseWithDefaults() *LoadBalancersRegionsIndex200Response {
	this := LoadBalancersRegionsIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancersRegionsIndex200Response) GetData() []LoadBalancerRegion {
	if o == nil || IsNil(o.Data) {
		var ret []LoadBalancerRegion
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancersRegionsIndex200Response) GetDataOk() ([]LoadBalancerRegion, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancersRegionsIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LoadBalancerRegion and assigns it to the Data field.
func (o *LoadBalancersRegionsIndex200Response) SetData(v []LoadBalancerRegion) {
	o.Data = v
}

func (o LoadBalancersRegionsIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancersRegionsIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancersRegionsIndex200Response struct {
	value *LoadBalancersRegionsIndex200Response
	isSet bool
}

func (v NullableLoadBalancersRegionsIndex200Response) Get() *LoadBalancersRegionsIndex200Response {
	return v.value
}

func (v *NullableLoadBalancersRegionsIndex200Response) Set(val *LoadBalancersRegionsIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancersRegionsIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancersRegionsIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancersRegionsIndex200Response(val *LoadBalancersRegionsIndex200Response) *NullableLoadBalancersRegionsIndex200Response {
	return &NullableLoadBalancersRegionsIndex200Response{value: val, isSet: true}
}

func (v NullableLoadBalancersRegionsIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancersRegionsIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


