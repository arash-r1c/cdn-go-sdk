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

// checks if the LoadBalancerOriginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerOriginResponse{}

// LoadBalancerOriginResponse struct for LoadBalancerOriginResponse
type LoadBalancerOriginResponse struct {
	Data *LoadBalancerOrigin `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewLoadBalancerOriginResponse instantiates a new LoadBalancerOriginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerOriginResponse() *LoadBalancerOriginResponse {
	this := LoadBalancerOriginResponse{}
	return &this
}

// NewLoadBalancerOriginResponseWithDefaults instantiates a new LoadBalancerOriginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerOriginResponseWithDefaults() *LoadBalancerOriginResponse {
	this := LoadBalancerOriginResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancerOriginResponse) GetData() LoadBalancerOrigin {
	if o == nil || IsNil(o.Data) {
		var ret LoadBalancerOrigin
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginResponse) GetDataOk() (*LoadBalancerOrigin, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancerOriginResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LoadBalancerOrigin and assigns it to the Data field.
func (o *LoadBalancerOriginResponse) SetData(v LoadBalancerOrigin) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadBalancerOriginResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerOriginResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *LoadBalancerOriginResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *LoadBalancerOriginResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *LoadBalancerOriginResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *LoadBalancerOriginResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o LoadBalancerOriginResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerOriginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableLoadBalancerOriginResponse struct {
	value *LoadBalancerOriginResponse
	isSet bool
}

func (v NullableLoadBalancerOriginResponse) Get() *LoadBalancerOriginResponse {
	return v.value
}

func (v *NullableLoadBalancerOriginResponse) Set(val *LoadBalancerOriginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerOriginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerOriginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerOriginResponse(val *LoadBalancerOriginResponse) *NullableLoadBalancerOriginResponse {
	return &NullableLoadBalancerOriginResponse{value: val, isSet: true}
}

func (v NullableLoadBalancerOriginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerOriginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


