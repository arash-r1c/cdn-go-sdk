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

// checks if the Redirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Redirect{}

// Redirect struct for Redirect
type Redirect struct {
	FRedirectToWww *string `json:"f_redirect_to_www,omitempty"`
}

// NewRedirect instantiates a new Redirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirect() *Redirect {
	this := Redirect{}
	return &this
}

// NewRedirectWithDefaults instantiates a new Redirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectWithDefaults() *Redirect {
	this := Redirect{}
	return &this
}

// GetFRedirectToWww returns the FRedirectToWww field value if set, zero value otherwise.
func (o *Redirect) GetFRedirectToWww() string {
	if o == nil || IsNil(o.FRedirectToWww) {
		var ret string
		return ret
	}
	return *o.FRedirectToWww
}

// GetFRedirectToWwwOk returns a tuple with the FRedirectToWww field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetFRedirectToWwwOk() (*string, bool) {
	if o == nil || IsNil(o.FRedirectToWww) {
		return nil, false
	}
	return o.FRedirectToWww, true
}

// HasFRedirectToWww returns a boolean if a field has been set.
func (o *Redirect) HasFRedirectToWww() bool {
	if o != nil && !IsNil(o.FRedirectToWww) {
		return true
	}

	return false
}

// SetFRedirectToWww gets a reference to the given string and assigns it to the FRedirectToWww field.
func (o *Redirect) SetFRedirectToWww(v string) {
	o.FRedirectToWww = &v
}

func (o Redirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Redirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FRedirectToWww) {
		toSerialize["f_redirect_to_www"] = o.FRedirectToWww
	}
	return toSerialize, nil
}

type NullableRedirect struct {
	value *Redirect
	isSet bool
}

func (v NullableRedirect) Get() *Redirect {
	return v.value
}

func (v *NullableRedirect) Set(val *Redirect) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirect(val *Redirect) *NullableRedirect {
	return &NullableRedirect{value: val, isSet: true}
}

func (v NullableRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


