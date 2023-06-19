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

// checks if the DnsRecordIpFilterMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRecordIpFilterMode{}

// DnsRecordIpFilterMode struct for DnsRecordIpFilterMode
type DnsRecordIpFilterMode struct {
	Count *string `json:"count,omitempty"`
	Order *string `json:"order,omitempty"`
	GeoFilter *string `json:"geo_filter,omitempty"`
}

// NewDnsRecordIpFilterMode instantiates a new DnsRecordIpFilterMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRecordIpFilterMode() *DnsRecordIpFilterMode {
	this := DnsRecordIpFilterMode{}
	return &this
}

// NewDnsRecordIpFilterModeWithDefaults instantiates a new DnsRecordIpFilterMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRecordIpFilterModeWithDefaults() *DnsRecordIpFilterMode {
	this := DnsRecordIpFilterMode{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DnsRecordIpFilterMode) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordIpFilterMode) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DnsRecordIpFilterMode) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *DnsRecordIpFilterMode) SetCount(v string) {
	o.Count = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DnsRecordIpFilterMode) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordIpFilterMode) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DnsRecordIpFilterMode) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *DnsRecordIpFilterMode) SetOrder(v string) {
	o.Order = &v
}

// GetGeoFilter returns the GeoFilter field value if set, zero value otherwise.
func (o *DnsRecordIpFilterMode) GetGeoFilter() string {
	if o == nil || IsNil(o.GeoFilter) {
		var ret string
		return ret
	}
	return *o.GeoFilter
}

// GetGeoFilterOk returns a tuple with the GeoFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordIpFilterMode) GetGeoFilterOk() (*string, bool) {
	if o == nil || IsNil(o.GeoFilter) {
		return nil, false
	}
	return o.GeoFilter, true
}

// HasGeoFilter returns a boolean if a field has been set.
func (o *DnsRecordIpFilterMode) HasGeoFilter() bool {
	if o != nil && !IsNil(o.GeoFilter) {
		return true
	}

	return false
}

// SetGeoFilter gets a reference to the given string and assigns it to the GeoFilter field.
func (o *DnsRecordIpFilterMode) SetGeoFilter(v string) {
	o.GeoFilter = &v
}

func (o DnsRecordIpFilterMode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRecordIpFilterMode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.GeoFilter) {
		toSerialize["geo_filter"] = o.GeoFilter
	}
	return toSerialize, nil
}

type NullableDnsRecordIpFilterMode struct {
	value *DnsRecordIpFilterMode
	isSet bool
}

func (v NullableDnsRecordIpFilterMode) Get() *DnsRecordIpFilterMode {
	return v.value
}

func (v *NullableDnsRecordIpFilterMode) Set(val *DnsRecordIpFilterMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecordIpFilterMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecordIpFilterMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecordIpFilterMode(val *DnsRecordIpFilterMode) *NullableDnsRecordIpFilterMode {
	return &NullableDnsRecordIpFilterMode{value: val, isSet: true}
}

func (v NullableDnsRecordIpFilterMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecordIpFilterMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


