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

// checks if the DdosRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdosRuleResponse{}

// DdosRuleResponse struct for DdosRuleResponse
type DdosRuleResponse struct {
	Data *DdosRule `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewDdosRuleResponse instantiates a new DdosRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosRuleResponse() *DdosRuleResponse {
	this := DdosRuleResponse{}
	return &this
}

// NewDdosRuleResponseWithDefaults instantiates a new DdosRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosRuleResponseWithDefaults() *DdosRuleResponse {
	this := DdosRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DdosRuleResponse) GetData() DdosRule {
	if o == nil || IsNil(o.Data) {
		var ret DdosRule
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRuleResponse) GetDataOk() (*DdosRule, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DdosRuleResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DdosRule and assigns it to the Data field.
func (o *DdosRuleResponse) SetData(v DdosRule) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdosRuleResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdosRuleResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DdosRuleResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *DdosRuleResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *DdosRuleResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *DdosRuleResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o DdosRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdosRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableDdosRuleResponse struct {
	value *DdosRuleResponse
	isSet bool
}

func (v NullableDdosRuleResponse) Get() *DdosRuleResponse {
	return v.value
}

func (v *NullableDdosRuleResponse) Set(val *DdosRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDdosRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDdosRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdosRuleResponse(val *DdosRuleResponse) *NullableDdosRuleResponse {
	return &NullableDdosRuleResponse{value: val, isSet: true}
}

func (v NullableDdosRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdosRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


