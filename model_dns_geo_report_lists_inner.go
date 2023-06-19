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

// checks if the DnsGeoReportListsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsGeoReportListsInner{}

// DnsGeoReportListsInner struct for DnsGeoReportListsInner
type DnsGeoReportListsInner struct {
	Country *string `json:"country,omitempty"`
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Requests *int32 `json:"requests,omitempty"`
}

// NewDnsGeoReportListsInner instantiates a new DnsGeoReportListsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsGeoReportListsInner() *DnsGeoReportListsInner {
	this := DnsGeoReportListsInner{}
	return &this
}

// NewDnsGeoReportListsInnerWithDefaults instantiates a new DnsGeoReportListsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsGeoReportListsInnerWithDefaults() *DnsGeoReportListsInner {
	this := DnsGeoReportListsInner{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DnsGeoReportListsInner) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportListsInner) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DnsGeoReportListsInner) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *DnsGeoReportListsInner) SetCountry(v string) {
	o.Country = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnsGeoReportListsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportListsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnsGeoReportListsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnsGeoReportListsInner) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DnsGeoReportListsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportListsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DnsGeoReportListsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DnsGeoReportListsInner) SetCode(v string) {
	o.Code = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *DnsGeoReportListsInner) GetRequests() int32 {
	if o == nil || IsNil(o.Requests) {
		var ret int32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportListsInner) GetRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *DnsGeoReportListsInner) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int32 and assigns it to the Requests field.
func (o *DnsGeoReportListsInner) SetRequests(v int32) {
	o.Requests = &v
}

func (o DnsGeoReportListsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsGeoReportListsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableDnsGeoReportListsInner struct {
	value *DnsGeoReportListsInner
	isSet bool
}

func (v NullableDnsGeoReportListsInner) Get() *DnsGeoReportListsInner {
	return v.value
}

func (v *NullableDnsGeoReportListsInner) Set(val *DnsGeoReportListsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsGeoReportListsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsGeoReportListsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsGeoReportListsInner(val *DnsGeoReportListsInner) *NullableDnsGeoReportListsInner {
	return &NullableDnsGeoReportListsInner{value: val, isSet: true}
}

func (v NullableDnsGeoReportListsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsGeoReportListsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


