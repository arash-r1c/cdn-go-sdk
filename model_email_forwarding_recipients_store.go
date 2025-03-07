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

// checks if the EmailForwardingRecipientsStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingRecipientsStore{}

// EmailForwardingRecipientsStore struct for EmailForwardingRecipientsStore
type EmailForwardingRecipientsStore struct {
	Email string `json:"email"`
}

// NewEmailForwardingRecipientsStore instantiates a new EmailForwardingRecipientsStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingRecipientsStore(email string) *EmailForwardingRecipientsStore {
	this := EmailForwardingRecipientsStore{}
	this.Email = email
	return &this
}

// NewEmailForwardingRecipientsStoreWithDefaults instantiates a new EmailForwardingRecipientsStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingRecipientsStoreWithDefaults() *EmailForwardingRecipientsStore {
	this := EmailForwardingRecipientsStore{}
	return &this
}

// GetEmail returns the Email field value
func (o *EmailForwardingRecipientsStore) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailForwardingRecipientsStore) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailForwardingRecipientsStore) SetEmail(v string) {
	o.Email = v
}

func (o EmailForwardingRecipientsStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingRecipientsStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableEmailForwardingRecipientsStore struct {
	value *EmailForwardingRecipientsStore
	isSet bool
}

func (v NullableEmailForwardingRecipientsStore) Get() *EmailForwardingRecipientsStore {
	return v.value
}

func (v *NullableEmailForwardingRecipientsStore) Set(val *EmailForwardingRecipientsStore) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingRecipientsStore) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingRecipientsStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingRecipientsStore(val *EmailForwardingRecipientsStore) *NullableEmailForwardingRecipientsStore {
	return &NullableEmailForwardingRecipientsStore{value: val, isSet: true}
}

func (v NullableEmailForwardingRecipientsStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingRecipientsStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


