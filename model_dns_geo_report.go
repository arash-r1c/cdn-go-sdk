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

// checks if the DnsGeoReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsGeoReport{}

// DnsGeoReport struct for DnsGeoReport
type DnsGeoReport struct {
	// Deprecated
	Statistics []interface{} `json:"statistics,omitempty"`
	Charts *DnsGeoReportCharts `json:"charts,omitempty"`
	Lists []DnsGeoReportListsInner `json:"lists,omitempty"`
}

// NewDnsGeoReport instantiates a new DnsGeoReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsGeoReport() *DnsGeoReport {
	this := DnsGeoReport{}
	return &this
}

// NewDnsGeoReportWithDefaults instantiates a new DnsGeoReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsGeoReportWithDefaults() *DnsGeoReport {
	this := DnsGeoReport{}
	return &this
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
// Deprecated
func (o *DnsGeoReport) GetStatistics() []interface{} {
	if o == nil || IsNil(o.Statistics) {
		var ret []interface{}
		return ret
	}
	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DnsGeoReport) GetStatisticsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *DnsGeoReport) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given []interface{} and assigns it to the Statistics field.
// Deprecated
func (o *DnsGeoReport) SetStatistics(v []interface{}) {
	o.Statistics = v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *DnsGeoReport) GetCharts() DnsGeoReportCharts {
	if o == nil || IsNil(o.Charts) {
		var ret DnsGeoReportCharts
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReport) GetChartsOk() (*DnsGeoReportCharts, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *DnsGeoReport) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given DnsGeoReportCharts and assigns it to the Charts field.
func (o *DnsGeoReport) SetCharts(v DnsGeoReportCharts) {
	o.Charts = &v
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *DnsGeoReport) GetLists() []DnsGeoReportListsInner {
	if o == nil || IsNil(o.Lists) {
		var ret []DnsGeoReportListsInner
		return ret
	}
	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReport) GetListsOk() ([]DnsGeoReportListsInner, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *DnsGeoReport) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given []DnsGeoReportListsInner and assigns it to the Lists field.
func (o *DnsGeoReport) SetLists(v []DnsGeoReportListsInner) {
	o.Lists = v
}

func (o DnsGeoReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsGeoReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	return toSerialize, nil
}

type NullableDnsGeoReport struct {
	value *DnsGeoReport
	isSet bool
}

func (v NullableDnsGeoReport) Get() *DnsGeoReport {
	return v.value
}

func (v *NullableDnsGeoReport) Set(val *DnsGeoReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsGeoReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsGeoReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsGeoReport(val *DnsGeoReport) *NullableDnsGeoReport {
	return &NullableDnsGeoReport{value: val, isSet: true}
}

func (v NullableDnsGeoReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsGeoReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


