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

// checks if the SavedTrafficsStatisticsTraffic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrafficsStatisticsTraffic{}

// SavedTrafficsStatisticsTraffic struct for SavedTrafficsStatisticsTraffic
type SavedTrafficsStatisticsTraffic struct {
	Saved *int64 `json:"saved,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewSavedTrafficsStatisticsTraffic instantiates a new SavedTrafficsStatisticsTraffic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrafficsStatisticsTraffic() *SavedTrafficsStatisticsTraffic {
	this := SavedTrafficsStatisticsTraffic{}
	return &this
}

// NewSavedTrafficsStatisticsTrafficWithDefaults instantiates a new SavedTrafficsStatisticsTraffic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrafficsStatisticsTrafficWithDefaults() *SavedTrafficsStatisticsTraffic {
	this := SavedTrafficsStatisticsTraffic{}
	return &this
}

// GetSaved returns the Saved field value if set, zero value otherwise.
func (o *SavedTrafficsStatisticsTraffic) GetSaved() int64 {
	if o == nil || IsNil(o.Saved) {
		var ret int64
		return ret
	}
	return *o.Saved
}

// GetSavedOk returns a tuple with the Saved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsStatisticsTraffic) GetSavedOk() (*int64, bool) {
	if o == nil || IsNil(o.Saved) {
		return nil, false
	}
	return o.Saved, true
}

// HasSaved returns a boolean if a field has been set.
func (o *SavedTrafficsStatisticsTraffic) HasSaved() bool {
	if o != nil && !IsNil(o.Saved) {
		return true
	}

	return false
}

// SetSaved gets a reference to the given int64 and assigns it to the Saved field.
func (o *SavedTrafficsStatisticsTraffic) SetSaved(v int64) {
	o.Saved = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SavedTrafficsStatisticsTraffic) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsStatisticsTraffic) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SavedTrafficsStatisticsTraffic) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *SavedTrafficsStatisticsTraffic) SetTotal(v int64) {
	o.Total = &v
}

func (o SavedTrafficsStatisticsTraffic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrafficsStatisticsTraffic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Saved) {
		toSerialize["saved"] = o.Saved
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableSavedTrafficsStatisticsTraffic struct {
	value *SavedTrafficsStatisticsTraffic
	isSet bool
}

func (v NullableSavedTrafficsStatisticsTraffic) Get() *SavedTrafficsStatisticsTraffic {
	return v.value
}

func (v *NullableSavedTrafficsStatisticsTraffic) Set(val *SavedTrafficsStatisticsTraffic) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrafficsStatisticsTraffic) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrafficsStatisticsTraffic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrafficsStatisticsTraffic(val *SavedTrafficsStatisticsTraffic) *NullableSavedTrafficsStatisticsTraffic {
	return &NullableSavedTrafficsStatisticsTraffic{value: val, isSet: true}
}

func (v NullableSavedTrafficsStatisticsTraffic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrafficsStatisticsTraffic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


