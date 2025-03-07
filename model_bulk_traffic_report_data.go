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

// checks if the BulkTrafficReportData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkTrafficReportData{}

// BulkTrafficReportData struct for BulkTrafficReportData
type BulkTrafficReportData struct {
	Requests *int64 `json:"requests,omitempty"`
	IngressBytes *int64 `json:"ingress_bytes,omitempty"`
	EgressBytes *BulkTrafficReportDataEgressBytes `json:"egress_bytes,omitempty"`
}

// NewBulkTrafficReportData instantiates a new BulkTrafficReportData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkTrafficReportData() *BulkTrafficReportData {
	this := BulkTrafficReportData{}
	return &this
}

// NewBulkTrafficReportDataWithDefaults instantiates a new BulkTrafficReportData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkTrafficReportDataWithDefaults() *BulkTrafficReportData {
	this := BulkTrafficReportData{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *BulkTrafficReportData) GetRequests() int64 {
	if o == nil || IsNil(o.Requests) {
		var ret int64
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTrafficReportData) GetRequestsOk() (*int64, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *BulkTrafficReportData) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int64 and assigns it to the Requests field.
func (o *BulkTrafficReportData) SetRequests(v int64) {
	o.Requests = &v
}

// GetIngressBytes returns the IngressBytes field value if set, zero value otherwise.
func (o *BulkTrafficReportData) GetIngressBytes() int64 {
	if o == nil || IsNil(o.IngressBytes) {
		var ret int64
		return ret
	}
	return *o.IngressBytes
}

// GetIngressBytesOk returns a tuple with the IngressBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTrafficReportData) GetIngressBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.IngressBytes) {
		return nil, false
	}
	return o.IngressBytes, true
}

// HasIngressBytes returns a boolean if a field has been set.
func (o *BulkTrafficReportData) HasIngressBytes() bool {
	if o != nil && !IsNil(o.IngressBytes) {
		return true
	}

	return false
}

// SetIngressBytes gets a reference to the given int64 and assigns it to the IngressBytes field.
func (o *BulkTrafficReportData) SetIngressBytes(v int64) {
	o.IngressBytes = &v
}

// GetEgressBytes returns the EgressBytes field value if set, zero value otherwise.
func (o *BulkTrafficReportData) GetEgressBytes() BulkTrafficReportDataEgressBytes {
	if o == nil || IsNil(o.EgressBytes) {
		var ret BulkTrafficReportDataEgressBytes
		return ret
	}
	return *o.EgressBytes
}

// GetEgressBytesOk returns a tuple with the EgressBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTrafficReportData) GetEgressBytesOk() (*BulkTrafficReportDataEgressBytes, bool) {
	if o == nil || IsNil(o.EgressBytes) {
		return nil, false
	}
	return o.EgressBytes, true
}

// HasEgressBytes returns a boolean if a field has been set.
func (o *BulkTrafficReportData) HasEgressBytes() bool {
	if o != nil && !IsNil(o.EgressBytes) {
		return true
	}

	return false
}

// SetEgressBytes gets a reference to the given BulkTrafficReportDataEgressBytes and assigns it to the EgressBytes field.
func (o *BulkTrafficReportData) SetEgressBytes(v BulkTrafficReportDataEgressBytes) {
	o.EgressBytes = &v
}

func (o BulkTrafficReportData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkTrafficReportData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.IngressBytes) {
		toSerialize["ingress_bytes"] = o.IngressBytes
	}
	if !IsNil(o.EgressBytes) {
		toSerialize["egress_bytes"] = o.EgressBytes
	}
	return toSerialize, nil
}

type NullableBulkTrafficReportData struct {
	value *BulkTrafficReportData
	isSet bool
}

func (v NullableBulkTrafficReportData) Get() *BulkTrafficReportData {
	return v.value
}

func (v *NullableBulkTrafficReportData) Set(val *BulkTrafficReportData) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkTrafficReportData) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkTrafficReportData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkTrafficReportData(val *BulkTrafficReportData) *NullableBulkTrafficReportData {
	return &NullableBulkTrafficReportData{value: val, isSet: true}
}

func (v NullableBulkTrafficReportData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkTrafficReportData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


