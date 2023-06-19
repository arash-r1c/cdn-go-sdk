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

// checks if the DomainStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainStore{}

// DomainStore struct for DomainStore
type DomainStore struct {
	// Name of the domain
	Domain string `json:"domain"`
	// If you want to register a subdomain, you can use cname setup by sending partial type
	DomainType *string `json:"domain_type,omitempty"`
	// - `0` - Traffic - `1` - Basic - `2` - Growth - `3` - Professional - `4` - Enterprise 
	PlanLevel *int32 `json:"plan_level,omitempty"`
}

// NewDomainStore instantiates a new DomainStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainStore(domain string) *DomainStore {
	this := DomainStore{}
	this.Domain = domain
	var domainType string = "full"
	this.DomainType = &domainType
	return &this
}

// NewDomainStoreWithDefaults instantiates a new DomainStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainStoreWithDefaults() *DomainStore {
	this := DomainStore{}
	var domainType string = "full"
	this.DomainType = &domainType
	return &this
}

// GetDomain returns the Domain field value
func (o *DomainStore) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainStore) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainStore) SetDomain(v string) {
	o.Domain = v
}

// GetDomainType returns the DomainType field value if set, zero value otherwise.
func (o *DomainStore) GetDomainType() string {
	if o == nil || IsNil(o.DomainType) {
		var ret string
		return ret
	}
	return *o.DomainType
}

// GetDomainTypeOk returns a tuple with the DomainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainStore) GetDomainTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DomainType) {
		return nil, false
	}
	return o.DomainType, true
}

// HasDomainType returns a boolean if a field has been set.
func (o *DomainStore) HasDomainType() bool {
	if o != nil && !IsNil(o.DomainType) {
		return true
	}

	return false
}

// SetDomainType gets a reference to the given string and assigns it to the DomainType field.
func (o *DomainStore) SetDomainType(v string) {
	o.DomainType = &v
}

// GetPlanLevel returns the PlanLevel field value if set, zero value otherwise.
func (o *DomainStore) GetPlanLevel() int32 {
	if o == nil || IsNil(o.PlanLevel) {
		var ret int32
		return ret
	}
	return *o.PlanLevel
}

// GetPlanLevelOk returns a tuple with the PlanLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainStore) GetPlanLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PlanLevel) {
		return nil, false
	}
	return o.PlanLevel, true
}

// HasPlanLevel returns a boolean if a field has been set.
func (o *DomainStore) HasPlanLevel() bool {
	if o != nil && !IsNil(o.PlanLevel) {
		return true
	}

	return false
}

// SetPlanLevel gets a reference to the given int32 and assigns it to the PlanLevel field.
func (o *DomainStore) SetPlanLevel(v int32) {
	o.PlanLevel = &v
}

func (o DomainStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	if !IsNil(o.DomainType) {
		toSerialize["domain_type"] = o.DomainType
	}
	if !IsNil(o.PlanLevel) {
		toSerialize["plan_level"] = o.PlanLevel
	}
	return toSerialize, nil
}

type NullableDomainStore struct {
	value *DomainStore
	isSet bool
}

func (v NullableDomainStore) Get() *DomainStore {
	return v.value
}

func (v *NullableDomainStore) Set(val *DomainStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainStore(val *DomainStore) *NullableDomainStore {
	return &NullableDomainStore{value: val, isSet: true}
}

func (v NullableDomainStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


