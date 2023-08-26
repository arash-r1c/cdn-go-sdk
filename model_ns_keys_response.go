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

// checks if the NsKeysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsKeysResponse{}

// NsKeysResponse struct for NsKeysResponse
type NsKeysResponse struct {
	Data *map[string]interface{} `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewNsKeysResponse instantiates a new NsKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsKeysResponse() *NsKeysResponse {
	this := NsKeysResponse{}
	return &this
}

// NewNsKeysResponseWithDefaults instantiates a new NsKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsKeysResponseWithDefaults() *NsKeysResponse {
	this := NsKeysResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NsKeysResponse) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsKeysResponse) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NsKeysResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *NsKeysResponse) SetData(v map[string]interface{}) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NsKeysResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NsKeysResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *NsKeysResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *NsKeysResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *NsKeysResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *NsKeysResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o NsKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsKeysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableNsKeysResponse struct {
	value *NsKeysResponse
	isSet bool
}

func (v NullableNsKeysResponse) Get() *NsKeysResponse {
	return v.value
}

func (v *NullableNsKeysResponse) Set(val *NsKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNsKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNsKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsKeysResponse(val *NsKeysResponse) *NullableNsKeysResponse {
	return &NullableNsKeysResponse{value: val, isSet: true}
}

func (v NullableNsKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


