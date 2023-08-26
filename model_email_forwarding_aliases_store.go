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

// checks if the EmailForwardingAliasesStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingAliasesStore{}

// EmailForwardingAliasesStore struct for EmailForwardingAliasesStore
type EmailForwardingAliasesStore struct {
	LocalPart string `json:"local_part"`
	Recipients []string `json:"recipients"`
}

// NewEmailForwardingAliasesStore instantiates a new EmailForwardingAliasesStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingAliasesStore(localPart string, recipients []string) *EmailForwardingAliasesStore {
	this := EmailForwardingAliasesStore{}
	this.LocalPart = localPart
	this.Recipients = recipients
	return &this
}

// NewEmailForwardingAliasesStoreWithDefaults instantiates a new EmailForwardingAliasesStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingAliasesStoreWithDefaults() *EmailForwardingAliasesStore {
	this := EmailForwardingAliasesStore{}
	return &this
}

// GetLocalPart returns the LocalPart field value
func (o *EmailForwardingAliasesStore) GetLocalPart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalPart
}

// GetLocalPartOk returns a tuple with the LocalPart field value
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesStore) GetLocalPartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalPart, true
}

// SetLocalPart sets field value
func (o *EmailForwardingAliasesStore) SetLocalPart(v string) {
	o.LocalPart = v
}

// GetRecipients returns the Recipients field value
func (o *EmailForwardingAliasesStore) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesStore) GetRecipientsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *EmailForwardingAliasesStore) SetRecipients(v []string) {
	o.Recipients = v
}

func (o EmailForwardingAliasesStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingAliasesStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["local_part"] = o.LocalPart
	toSerialize["recipients"] = o.Recipients
	return toSerialize, nil
}

type NullableEmailForwardingAliasesStore struct {
	value *EmailForwardingAliasesStore
	isSet bool
}

func (v NullableEmailForwardingAliasesStore) Get() *EmailForwardingAliasesStore {
	return v.value
}

func (v *NullableEmailForwardingAliasesStore) Set(val *EmailForwardingAliasesStore) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingAliasesStore) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingAliasesStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingAliasesStore(val *EmailForwardingAliasesStore) *NullableEmailForwardingAliasesStore {
	return &NullableEmailForwardingAliasesStore{value: val, isSet: true}
}

func (v NullableEmailForwardingAliasesStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingAliasesStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


