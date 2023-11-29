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

// checks if the DomainWafPackageStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainWafPackageStore{}

// DomainWafPackageStore struct for DomainWafPackageStore
type DomainWafPackageStore struct {
	Id string `json:"id"`
}

// NewDomainWafPackageStore instantiates a new DomainWafPackageStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainWafPackageStore(id string) *DomainWafPackageStore {
	this := DomainWafPackageStore{}
	this.Id = id
	return &this
}

// NewDomainWafPackageStoreWithDefaults instantiates a new DomainWafPackageStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWafPackageStoreWithDefaults() *DomainWafPackageStore {
	this := DomainWafPackageStore{}
	return &this
}

// GetId returns the Id field value
func (o *DomainWafPackageStore) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DomainWafPackageStore) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DomainWafPackageStore) SetId(v string) {
	o.Id = v
}

func (o DomainWafPackageStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainWafPackageStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableDomainWafPackageStore struct {
	value *DomainWafPackageStore
	isSet bool
}

func (v NullableDomainWafPackageStore) Get() *DomainWafPackageStore {
	return v.value
}

func (v *NullableDomainWafPackageStore) Set(val *DomainWafPackageStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainWafPackageStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainWafPackageStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainWafPackageStore(val *DomainWafPackageStore) *NullableDomainWafPackageStore {
	return &NullableDomainWafPackageStore{value: val, isSet: true}
}

func (v NullableDomainWafPackageStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainWafPackageStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


