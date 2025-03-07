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

// checks if the HealthCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckResponse{}

// HealthCheckResponse struct for HealthCheckResponse
type HealthCheckResponse struct {
	Data *HealthCheckView `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewHealthCheckResponse instantiates a new HealthCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckResponse() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// NewHealthCheckResponseWithDefaults instantiates a new HealthCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckResponseWithDefaults() *HealthCheckResponse {
	this := HealthCheckResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HealthCheckResponse) GetData() HealthCheckView {
	if o == nil || IsNil(o.Data) {
		var ret HealthCheckView
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResponse) GetDataOk() (*HealthCheckView, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given HealthCheckView and assigns it to the Data field.
func (o *HealthCheckResponse) SetData(v HealthCheckView) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthCheckResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthCheckResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *HealthCheckResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *HealthCheckResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *HealthCheckResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *HealthCheckResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o HealthCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableHealthCheckResponse struct {
	value *HealthCheckResponse
	isSet bool
}

func (v NullableHealthCheckResponse) Get() *HealthCheckResponse {
	return v.value
}

func (v *NullableHealthCheckResponse) Set(val *HealthCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckResponse(val *HealthCheckResponse) *NullableHealthCheckResponse {
	return &NullableHealthCheckResponse{value: val, isSet: true}
}

func (v NullableHealthCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


