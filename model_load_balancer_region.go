/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the LoadBalancerRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerRegion{}

// LoadBalancerRegion struct for LoadBalancerRegion
type LoadBalancerRegion struct {
	Id *string `json:"id,omitempty"`
	Region *string `json:"region,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewLoadBalancerRegion instantiates a new LoadBalancerRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerRegion() *LoadBalancerRegion {
	this := LoadBalancerRegion{}
	return &this
}

// NewLoadBalancerRegionWithDefaults instantiates a new LoadBalancerRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerRegionWithDefaults() *LoadBalancerRegion {
	this := LoadBalancerRegion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancerRegion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerRegion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadBalancerRegion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadBalancerRegion) SetId(v string) {
	o.Id = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *LoadBalancerRegion) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerRegion) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *LoadBalancerRegion) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *LoadBalancerRegion) SetRegion(v string) {
	o.Region = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerRegion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerRegion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerRegion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerRegion) SetName(v string) {
	o.Name = &v
}

func (o LoadBalancerRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	// skip: name is readOnly
	return toSerialize, nil
}

type NullableLoadBalancerRegion struct {
	value *LoadBalancerRegion
	isSet bool
}

func (v NullableLoadBalancerRegion) Get() *LoadBalancerRegion {
	return v.value
}

func (v *NullableLoadBalancerRegion) Set(val *LoadBalancerRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerRegion(val *LoadBalancerRegion) *NullableLoadBalancerRegion {
	return &NullableLoadBalancerRegion{value: val, isSet: true}
}

func (v NullableLoadBalancerRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


