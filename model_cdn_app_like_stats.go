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

// checks if the CdnAppLikeStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdnAppLikeStats{}

// CdnAppLikeStats struct for CdnAppLikeStats
type CdnAppLikeStats struct {
	LikesCount *int32 `json:"likes_count,omitempty"`
	DislikesCount *int32 `json:"dislikes_count,omitempty"`
}

// NewCdnAppLikeStats instantiates a new CdnAppLikeStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnAppLikeStats() *CdnAppLikeStats {
	this := CdnAppLikeStats{}
	return &this
}

// NewCdnAppLikeStatsWithDefaults instantiates a new CdnAppLikeStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnAppLikeStatsWithDefaults() *CdnAppLikeStats {
	this := CdnAppLikeStats{}
	return &this
}

// GetLikesCount returns the LikesCount field value if set, zero value otherwise.
func (o *CdnAppLikeStats) GetLikesCount() int32 {
	if o == nil || IsNil(o.LikesCount) {
		var ret int32
		return ret
	}
	return *o.LikesCount
}

// GetLikesCountOk returns a tuple with the LikesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnAppLikeStats) GetLikesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LikesCount) {
		return nil, false
	}
	return o.LikesCount, true
}

// HasLikesCount returns a boolean if a field has been set.
func (o *CdnAppLikeStats) HasLikesCount() bool {
	if o != nil && !IsNil(o.LikesCount) {
		return true
	}

	return false
}

// SetLikesCount gets a reference to the given int32 and assigns it to the LikesCount field.
func (o *CdnAppLikeStats) SetLikesCount(v int32) {
	o.LikesCount = &v
}

// GetDislikesCount returns the DislikesCount field value if set, zero value otherwise.
func (o *CdnAppLikeStats) GetDislikesCount() int32 {
	if o == nil || IsNil(o.DislikesCount) {
		var ret int32
		return ret
	}
	return *o.DislikesCount
}

// GetDislikesCountOk returns a tuple with the DislikesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnAppLikeStats) GetDislikesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DislikesCount) {
		return nil, false
	}
	return o.DislikesCount, true
}

// HasDislikesCount returns a boolean if a field has been set.
func (o *CdnAppLikeStats) HasDislikesCount() bool {
	if o != nil && !IsNil(o.DislikesCount) {
		return true
	}

	return false
}

// SetDislikesCount gets a reference to the given int32 and assigns it to the DislikesCount field.
func (o *CdnAppLikeStats) SetDislikesCount(v int32) {
	o.DislikesCount = &v
}

func (o CdnAppLikeStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdnAppLikeStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LikesCount) {
		toSerialize["likes_count"] = o.LikesCount
	}
	if !IsNil(o.DislikesCount) {
		toSerialize["dislikes_count"] = o.DislikesCount
	}
	return toSerialize, nil
}

type NullableCdnAppLikeStats struct {
	value *CdnAppLikeStats
	isSet bool
}

func (v NullableCdnAppLikeStats) Get() *CdnAppLikeStats {
	return v.value
}

func (v *NullableCdnAppLikeStats) Set(val *CdnAppLikeStats) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnAppLikeStats) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnAppLikeStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnAppLikeStats(val *CdnAppLikeStats) *NullableCdnAppLikeStats {
	return &NullableCdnAppLikeStats{value: val, isSet: true}
}

func (v NullableCdnAppLikeStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnAppLikeStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


