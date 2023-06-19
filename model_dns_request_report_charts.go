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

// checks if the DnsRequestReportCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRequestReportCharts{}

// DnsRequestReportCharts struct for DnsRequestReportCharts
type DnsRequestReportCharts struct {
	Requests *DnsRequestReportChartsRequests `json:"requests,omitempty"`
}

// NewDnsRequestReportCharts instantiates a new DnsRequestReportCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRequestReportCharts() *DnsRequestReportCharts {
	this := DnsRequestReportCharts{}
	return &this
}

// NewDnsRequestReportChartsWithDefaults instantiates a new DnsRequestReportCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRequestReportChartsWithDefaults() *DnsRequestReportCharts {
	this := DnsRequestReportCharts{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *DnsRequestReportCharts) GetRequests() DnsRequestReportChartsRequests {
	if o == nil || IsNil(o.Requests) {
		var ret DnsRequestReportChartsRequests
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReportCharts) GetRequestsOk() (*DnsRequestReportChartsRequests, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *DnsRequestReportCharts) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given DnsRequestReportChartsRequests and assigns it to the Requests field.
func (o *DnsRequestReportCharts) SetRequests(v DnsRequestReportChartsRequests) {
	o.Requests = &v
}

func (o DnsRequestReportCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRequestReportCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableDnsRequestReportCharts struct {
	value *DnsRequestReportCharts
	isSet bool
}

func (v NullableDnsRequestReportCharts) Get() *DnsRequestReportCharts {
	return v.value
}

func (v *NullableDnsRequestReportCharts) Set(val *DnsRequestReportCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRequestReportCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRequestReportCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRequestReportCharts(val *DnsRequestReportCharts) *NullableDnsRequestReportCharts {
	return &NullableDnsRequestReportCharts{value: val, isSet: true}
}

func (v NullableDnsRequestReportCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRequestReportCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


