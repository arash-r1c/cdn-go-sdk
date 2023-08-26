/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.106.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the ReportsVisitorsHighRequestIps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsVisitorsHighRequestIps200Response{}

// ReportsVisitorsHighRequestIps200Response struct for ReportsVisitorsHighRequestIps200Response
type ReportsVisitorsHighRequestIps200Response struct {
	Data []HighRequestedIp `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta *PaginatedResponseMeta `json:"meta,omitempty"`
}

// NewReportsVisitorsHighRequestIps200Response instantiates a new ReportsVisitorsHighRequestIps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsVisitorsHighRequestIps200Response() *ReportsVisitorsHighRequestIps200Response {
	this := ReportsVisitorsHighRequestIps200Response{}
	return &this
}

// NewReportsVisitorsHighRequestIps200ResponseWithDefaults instantiates a new ReportsVisitorsHighRequestIps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsVisitorsHighRequestIps200ResponseWithDefaults() *ReportsVisitorsHighRequestIps200Response {
	this := ReportsVisitorsHighRequestIps200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReportsVisitorsHighRequestIps200Response) GetData() []HighRequestedIp {
	if o == nil || IsNil(o.Data) {
		var ret []HighRequestedIp
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsVisitorsHighRequestIps200Response) GetDataOk() ([]HighRequestedIp, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReportsVisitorsHighRequestIps200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []HighRequestedIp and assigns it to the Data field.
func (o *ReportsVisitorsHighRequestIps200Response) SetData(v []HighRequestedIp) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReportsVisitorsHighRequestIps200Response) GetLinks() PaginatedResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsVisitorsHighRequestIps200Response) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReportsVisitorsHighRequestIps200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedResponseLinks and assigns it to the Links field.
func (o *ReportsVisitorsHighRequestIps200Response) SetLinks(v PaginatedResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ReportsVisitorsHighRequestIps200Response) GetMeta() PaginatedResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatedResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsVisitorsHighRequestIps200Response) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ReportsVisitorsHighRequestIps200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatedResponseMeta and assigns it to the Meta field.
func (o *ReportsVisitorsHighRequestIps200Response) SetMeta(v PaginatedResponseMeta) {
	o.Meta = &v
}

func (o ReportsVisitorsHighRequestIps200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsVisitorsHighRequestIps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableReportsVisitorsHighRequestIps200Response struct {
	value *ReportsVisitorsHighRequestIps200Response
	isSet bool
}

func (v NullableReportsVisitorsHighRequestIps200Response) Get() *ReportsVisitorsHighRequestIps200Response {
	return v.value
}

func (v *NullableReportsVisitorsHighRequestIps200Response) Set(val *ReportsVisitorsHighRequestIps200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsVisitorsHighRequestIps200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsVisitorsHighRequestIps200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsVisitorsHighRequestIps200Response(val *ReportsVisitorsHighRequestIps200Response) *NullableReportsVisitorsHighRequestIps200Response {
	return &NullableReportsVisitorsHighRequestIps200Response{value: val, isSet: true}
}

func (v NullableReportsVisitorsHighRequestIps200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsVisitorsHighRequestIps200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


