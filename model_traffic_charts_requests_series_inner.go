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

// checks if the TrafficChartsRequestsSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficChartsRequestsSeriesInner{}

// TrafficChartsRequestsSeriesInner struct for TrafficChartsRequestsSeriesInner
type TrafficChartsRequestsSeriesInner struct {
	Name *string `json:"name,omitempty"`
	Data []int32 `json:"data,omitempty"`
}

// NewTrafficChartsRequestsSeriesInner instantiates a new TrafficChartsRequestsSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficChartsRequestsSeriesInner() *TrafficChartsRequestsSeriesInner {
	this := TrafficChartsRequestsSeriesInner{}
	return &this
}

// NewTrafficChartsRequestsSeriesInnerWithDefaults instantiates a new TrafficChartsRequestsSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficChartsRequestsSeriesInnerWithDefaults() *TrafficChartsRequestsSeriesInner {
	this := TrafficChartsRequestsSeriesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrafficChartsRequestsSeriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficChartsRequestsSeriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrafficChartsRequestsSeriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrafficChartsRequestsSeriesInner) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TrafficChartsRequestsSeriesInner) GetData() []int32 {
	if o == nil || IsNil(o.Data) {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficChartsRequestsSeriesInner) GetDataOk() ([]int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TrafficChartsRequestsSeriesInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *TrafficChartsRequestsSeriesInner) SetData(v []int32) {
	o.Data = v
}

func (o TrafficChartsRequestsSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficChartsRequestsSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTrafficChartsRequestsSeriesInner struct {
	value *TrafficChartsRequestsSeriesInner
	isSet bool
}

func (v NullableTrafficChartsRequestsSeriesInner) Get() *TrafficChartsRequestsSeriesInner {
	return v.value
}

func (v *NullableTrafficChartsRequestsSeriesInner) Set(val *TrafficChartsRequestsSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficChartsRequestsSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficChartsRequestsSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficChartsRequestsSeriesInner(val *TrafficChartsRequestsSeriesInner) *NullableTrafficChartsRequestsSeriesInner {
	return &NullableTrafficChartsRequestsSeriesInner{value: val, isSet: true}
}

func (v NullableTrafficChartsRequestsSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficChartsRequestsSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


