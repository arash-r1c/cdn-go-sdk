/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.115.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the TrafficChartsTraffics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficChartsTraffics{}

// TrafficChartsTraffics struct for TrafficChartsTraffics
type TrafficChartsTraffics struct {
	Title *string `json:"title,omitempty"`
	Categories []time.Time `json:"categories,omitempty"`
	Series []TrafficChartsTrafficsSeriesInner `json:"series,omitempty"`
}

// NewTrafficChartsTraffics instantiates a new TrafficChartsTraffics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficChartsTraffics() *TrafficChartsTraffics {
	this := TrafficChartsTraffics{}
	return &this
}

// NewTrafficChartsTrafficsWithDefaults instantiates a new TrafficChartsTraffics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficChartsTrafficsWithDefaults() *TrafficChartsTraffics {
	this := TrafficChartsTraffics{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TrafficChartsTraffics) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficChartsTraffics) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TrafficChartsTraffics) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TrafficChartsTraffics) SetTitle(v string) {
	o.Title = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *TrafficChartsTraffics) GetCategories() []time.Time {
	if o == nil || IsNil(o.Categories) {
		var ret []time.Time
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficChartsTraffics) GetCategoriesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *TrafficChartsTraffics) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []time.Time and assigns it to the Categories field.
func (o *TrafficChartsTraffics) SetCategories(v []time.Time) {
	o.Categories = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *TrafficChartsTraffics) GetSeries() []TrafficChartsTrafficsSeriesInner {
	if o == nil || IsNil(o.Series) {
		var ret []TrafficChartsTrafficsSeriesInner
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficChartsTraffics) GetSeriesOk() ([]TrafficChartsTrafficsSeriesInner, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *TrafficChartsTraffics) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []TrafficChartsTrafficsSeriesInner and assigns it to the Series field.
func (o *TrafficChartsTraffics) SetSeries(v []TrafficChartsTrafficsSeriesInner) {
	o.Series = v
}

func (o TrafficChartsTraffics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficChartsTraffics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullableTrafficChartsTraffics struct {
	value *TrafficChartsTraffics
	isSet bool
}

func (v NullableTrafficChartsTraffics) Get() *TrafficChartsTraffics {
	return v.value
}

func (v *NullableTrafficChartsTraffics) Set(val *TrafficChartsTraffics) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficChartsTraffics) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficChartsTraffics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficChartsTraffics(val *TrafficChartsTraffics) *NullableTrafficChartsTraffics {
	return &NullableTrafficChartsTraffics{value: val, isSet: true}
}

func (v NullableTrafficChartsTraffics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficChartsTraffics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


