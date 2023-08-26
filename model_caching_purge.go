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

// checks if the CachingPurge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CachingPurge{}

// CachingPurge struct for CachingPurge
type CachingPurge struct {
	// tags is deprecated
	Purge string `json:"purge"`
	// URLs to be purged from cache. Required if purge value is set to individual.
	PurgeUrls []string `json:"purge_urls,omitempty"`
	// Tags to be purged from cache. Required if purge value is set to tags. Each tag must be 32 characters or less. Only ASCII characters are acceptable. 
	// Deprecated
	PurgeTags []string `json:"purge_tags,omitempty"`
}

// NewCachingPurge instantiates a new CachingPurge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCachingPurge(purge string) *CachingPurge {
	this := CachingPurge{}
	this.Purge = purge
	return &this
}

// NewCachingPurgeWithDefaults instantiates a new CachingPurge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCachingPurgeWithDefaults() *CachingPurge {
	this := CachingPurge{}
	return &this
}

// GetPurge returns the Purge field value
func (o *CachingPurge) GetPurge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purge
}

// GetPurgeOk returns a tuple with the Purge field value
// and a boolean to check if the value has been set.
func (o *CachingPurge) GetPurgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purge, true
}

// SetPurge sets field value
func (o *CachingPurge) SetPurge(v string) {
	o.Purge = v
}

// GetPurgeUrls returns the PurgeUrls field value if set, zero value otherwise.
func (o *CachingPurge) GetPurgeUrls() []string {
	if o == nil || IsNil(o.PurgeUrls) {
		var ret []string
		return ret
	}
	return o.PurgeUrls
}

// GetPurgeUrlsOk returns a tuple with the PurgeUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CachingPurge) GetPurgeUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.PurgeUrls) {
		return nil, false
	}
	return o.PurgeUrls, true
}

// HasPurgeUrls returns a boolean if a field has been set.
func (o *CachingPurge) HasPurgeUrls() bool {
	if o != nil && !IsNil(o.PurgeUrls) {
		return true
	}

	return false
}

// SetPurgeUrls gets a reference to the given []string and assigns it to the PurgeUrls field.
func (o *CachingPurge) SetPurgeUrls(v []string) {
	o.PurgeUrls = v
}

// GetPurgeTags returns the PurgeTags field value if set, zero value otherwise.
// Deprecated
func (o *CachingPurge) GetPurgeTags() []string {
	if o == nil || IsNil(o.PurgeTags) {
		var ret []string
		return ret
	}
	return o.PurgeTags
}

// GetPurgeTagsOk returns a tuple with the PurgeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CachingPurge) GetPurgeTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.PurgeTags) {
		return nil, false
	}
	return o.PurgeTags, true
}

// HasPurgeTags returns a boolean if a field has been set.
func (o *CachingPurge) HasPurgeTags() bool {
	if o != nil && !IsNil(o.PurgeTags) {
		return true
	}

	return false
}

// SetPurgeTags gets a reference to the given []string and assigns it to the PurgeTags field.
// Deprecated
func (o *CachingPurge) SetPurgeTags(v []string) {
	o.PurgeTags = v
}

func (o CachingPurge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CachingPurge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["purge"] = o.Purge
	if !IsNil(o.PurgeUrls) {
		toSerialize["purge_urls"] = o.PurgeUrls
	}
	if !IsNil(o.PurgeTags) {
		toSerialize["purge_tags"] = o.PurgeTags
	}
	return toSerialize, nil
}

type NullableCachingPurge struct {
	value *CachingPurge
	isSet bool
}

func (v NullableCachingPurge) Get() *CachingPurge {
	return v.value
}

func (v *NullableCachingPurge) Set(val *CachingPurge) {
	v.value = val
	v.isSet = true
}

func (v NullableCachingPurge) IsSet() bool {
	return v.isSet
}

func (v *NullableCachingPurge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCachingPurge(val *CachingPurge) *NullableCachingPurge {
	return &NullableCachingPurge{value: val, isSet: true}
}

func (v NullableCachingPurge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCachingPurge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


