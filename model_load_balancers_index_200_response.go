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

// checks if the LoadBalancersIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancersIndex200Response{}

// LoadBalancersIndex200Response struct for LoadBalancersIndex200Response
type LoadBalancersIndex200Response struct {
	Data []LoadBalancer `json:"data,omitempty"`
}

// NewLoadBalancersIndex200Response instantiates a new LoadBalancersIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancersIndex200Response() *LoadBalancersIndex200Response {
	this := LoadBalancersIndex200Response{}
	return &this
}

// NewLoadBalancersIndex200ResponseWithDefaults instantiates a new LoadBalancersIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancersIndex200ResponseWithDefaults() *LoadBalancersIndex200Response {
	this := LoadBalancersIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancersIndex200Response) GetData() []LoadBalancer {
	if o == nil || IsNil(o.Data) {
		var ret []LoadBalancer
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancersIndex200Response) GetDataOk() ([]LoadBalancer, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancersIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LoadBalancer and assigns it to the Data field.
func (o *LoadBalancersIndex200Response) SetData(v []LoadBalancer) {
	o.Data = v
}

func (o LoadBalancersIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancersIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancersIndex200Response struct {
	value *LoadBalancersIndex200Response
	isSet bool
}

func (v NullableLoadBalancersIndex200Response) Get() *LoadBalancersIndex200Response {
	return v.value
}

func (v *NullableLoadBalancersIndex200Response) Set(val *LoadBalancersIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancersIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancersIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancersIndex200Response(val *LoadBalancersIndex200Response) *NullableLoadBalancersIndex200Response {
	return &NullableLoadBalancersIndex200Response{value: val, isSet: true}
}

func (v NullableLoadBalancersIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancersIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


