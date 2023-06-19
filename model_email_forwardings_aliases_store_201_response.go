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

// checks if the EmailForwardingsAliasesStore201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingsAliasesStore201Response{}

// EmailForwardingsAliasesStore201Response struct for EmailForwardingsAliasesStore201Response
type EmailForwardingsAliasesStore201Response struct {
	Data *EmailForwardingAlias `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewEmailForwardingsAliasesStore201Response instantiates a new EmailForwardingsAliasesStore201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingsAliasesStore201Response() *EmailForwardingsAliasesStore201Response {
	this := EmailForwardingsAliasesStore201Response{}
	return &this
}

// NewEmailForwardingsAliasesStore201ResponseWithDefaults instantiates a new EmailForwardingsAliasesStore201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingsAliasesStore201ResponseWithDefaults() *EmailForwardingsAliasesStore201Response {
	this := EmailForwardingsAliasesStore201Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailForwardingsAliasesStore201Response) GetData() EmailForwardingAlias {
	if o == nil || IsNil(o.Data) {
		var ret EmailForwardingAlias
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingsAliasesStore201Response) GetDataOk() (*EmailForwardingAlias, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailForwardingsAliasesStore201Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailForwardingAlias and assigns it to the Data field.
func (o *EmailForwardingsAliasesStore201Response) SetData(v EmailForwardingAlias) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailForwardingsAliasesStore201Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailForwardingsAliasesStore201Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *EmailForwardingsAliasesStore201Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *EmailForwardingsAliasesStore201Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *EmailForwardingsAliasesStore201Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *EmailForwardingsAliasesStore201Response) UnsetMessage() {
	o.Message.Unset()
}

func (o EmailForwardingsAliasesStore201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingsAliasesStore201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableEmailForwardingsAliasesStore201Response struct {
	value *EmailForwardingsAliasesStore201Response
	isSet bool
}

func (v NullableEmailForwardingsAliasesStore201Response) Get() *EmailForwardingsAliasesStore201Response {
	return v.value
}

func (v *NullableEmailForwardingsAliasesStore201Response) Set(val *EmailForwardingsAliasesStore201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingsAliasesStore201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingsAliasesStore201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingsAliasesStore201Response(val *EmailForwardingsAliasesStore201Response) *NullableEmailForwardingsAliasesStore201Response {
	return &NullableEmailForwardingsAliasesStore201Response{value: val, isSet: true}
}

func (v NullableEmailForwardingsAliasesStore201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingsAliasesStore201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


