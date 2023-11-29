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

// checks if the DnsRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRecordResponse{}

// DnsRecordResponse struct for DnsRecordResponse
type DnsRecordResponse struct {
	Data *DnsRecordGeneric `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewDnsRecordResponse instantiates a new DnsRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRecordResponse() *DnsRecordResponse {
	this := DnsRecordResponse{}
	return &this
}

// NewDnsRecordResponseWithDefaults instantiates a new DnsRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRecordResponseWithDefaults() *DnsRecordResponse {
	this := DnsRecordResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DnsRecordResponse) GetData() DnsRecordGeneric {
	if o == nil || IsNil(o.Data) {
		var ret DnsRecordGeneric
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordResponse) GetDataOk() (*DnsRecordGeneric, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DnsRecordResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DnsRecordGeneric and assigns it to the Data field.
func (o *DnsRecordResponse) SetData(v DnsRecordGeneric) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DnsRecordResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsRecordResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DnsRecordResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *DnsRecordResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *DnsRecordResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *DnsRecordResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o DnsRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableDnsRecordResponse struct {
	value *DnsRecordResponse
	isSet bool
}

func (v NullableDnsRecordResponse) Get() *DnsRecordResponse {
	return v.value
}

func (v *NullableDnsRecordResponse) Set(val *DnsRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecordResponse(val *DnsRecordResponse) *NullableDnsRecordResponse {
	return &NullableDnsRecordResponse{value: val, isSet: true}
}

func (v NullableDnsRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


