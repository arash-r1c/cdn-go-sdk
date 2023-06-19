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

// checks if the FirewallIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallIndex200Response{}

// FirewallIndex200Response struct for FirewallIndex200Response
type FirewallIndex200Response struct {
	// Deprecated
	Data *Firewall `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewFirewallIndex200Response instantiates a new FirewallIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallIndex200Response() *FirewallIndex200Response {
	this := FirewallIndex200Response{}
	return &this
}

// NewFirewallIndex200ResponseWithDefaults instantiates a new FirewallIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallIndex200ResponseWithDefaults() *FirewallIndex200Response {
	this := FirewallIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
// Deprecated
func (o *FirewallIndex200Response) GetData() Firewall {
	if o == nil || IsNil(o.Data) {
		var ret Firewall
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FirewallIndex200Response) GetDataOk() (*Firewall, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FirewallIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Firewall and assigns it to the Data field.
// Deprecated
func (o *FirewallIndex200Response) SetData(v Firewall) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallIndex200Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallIndex200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *FirewallIndex200Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *FirewallIndex200Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *FirewallIndex200Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *FirewallIndex200Response) UnsetMessage() {
	o.Message.Unset()
}

func (o FirewallIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableFirewallIndex200Response struct {
	value *FirewallIndex200Response
	isSet bool
}

func (v NullableFirewallIndex200Response) Get() *FirewallIndex200Response {
	return v.value
}

func (v *NullableFirewallIndex200Response) Set(val *FirewallIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallIndex200Response(val *FirewallIndex200Response) *NullableFirewallIndex200Response {
	return &NullableFirewallIndex200Response{value: val, isSet: true}
}

func (v NullableFirewallIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


