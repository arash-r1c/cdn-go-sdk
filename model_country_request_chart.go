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

// checks if the CountryRequestChart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryRequestChart{}

// CountryRequestChart struct for CountryRequestChart
type CountryRequestChart struct {
	// The fill key for the country
	FillKey *int32 `json:"fillKey,omitempty"`
	// The name of the country
	Name *string `json:"name,omitempty"`
	// The number of requests from the country
	Value *int32 `json:"value,omitempty"`
}

// NewCountryRequestChart instantiates a new CountryRequestChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryRequestChart() *CountryRequestChart {
	this := CountryRequestChart{}
	return &this
}

// NewCountryRequestChartWithDefaults instantiates a new CountryRequestChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryRequestChartWithDefaults() *CountryRequestChart {
	this := CountryRequestChart{}
	return &this
}

// GetFillKey returns the FillKey field value if set, zero value otherwise.
func (o *CountryRequestChart) GetFillKey() int32 {
	if o == nil || IsNil(o.FillKey) {
		var ret int32
		return ret
	}
	return *o.FillKey
}

// GetFillKeyOk returns a tuple with the FillKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryRequestChart) GetFillKeyOk() (*int32, bool) {
	if o == nil || IsNil(o.FillKey) {
		return nil, false
	}
	return o.FillKey, true
}

// HasFillKey returns a boolean if a field has been set.
func (o *CountryRequestChart) HasFillKey() bool {
	if o != nil && !IsNil(o.FillKey) {
		return true
	}

	return false
}

// SetFillKey gets a reference to the given int32 and assigns it to the FillKey field.
func (o *CountryRequestChart) SetFillKey(v int32) {
	o.FillKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CountryRequestChart) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryRequestChart) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CountryRequestChart) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CountryRequestChart) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CountryRequestChart) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryRequestChart) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CountryRequestChart) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *CountryRequestChart) SetValue(v int32) {
	o.Value = &v
}

func (o CountryRequestChart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryRequestChart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FillKey) {
		toSerialize["fillKey"] = o.FillKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCountryRequestChart struct {
	value *CountryRequestChart
	isSet bool
}

func (v NullableCountryRequestChart) Get() *CountryRequestChart {
	return v.value
}

func (v *NullableCountryRequestChart) Set(val *CountryRequestChart) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryRequestChart) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryRequestChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryRequestChart(val *CountryRequestChart) *NullableCountryRequestChart {
	return &NullableCountryRequestChart{value: val, isSet: true}
}

func (v NullableCountryRequestChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryRequestChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


