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

// checks if the PageRulesIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRulesIndex200Response{}

// PageRulesIndex200Response struct for PageRulesIndex200Response
type PageRulesIndex200Response struct {
	Data []PageRuleSummary `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta *PaginatedResponseMeta `json:"meta,omitempty"`
}

// NewPageRulesIndex200Response instantiates a new PageRulesIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRulesIndex200Response() *PageRulesIndex200Response {
	this := PageRulesIndex200Response{}
	return &this
}

// NewPageRulesIndex200ResponseWithDefaults instantiates a new PageRulesIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRulesIndex200ResponseWithDefaults() *PageRulesIndex200Response {
	this := PageRulesIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PageRulesIndex200Response) GetData() []PageRuleSummary {
	if o == nil || IsNil(o.Data) {
		var ret []PageRuleSummary
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRulesIndex200Response) GetDataOk() ([]PageRuleSummary, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PageRulesIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PageRuleSummary and assigns it to the Data field.
func (o *PageRulesIndex200Response) SetData(v []PageRuleSummary) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PageRulesIndex200Response) GetLinks() PaginatedResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRulesIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PageRulesIndex200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedResponseLinks and assigns it to the Links field.
func (o *PageRulesIndex200Response) SetLinks(v PaginatedResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PageRulesIndex200Response) GetMeta() PaginatedResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatedResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRulesIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PageRulesIndex200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatedResponseMeta and assigns it to the Meta field.
func (o *PageRulesIndex200Response) SetMeta(v PaginatedResponseMeta) {
	o.Meta = &v
}

func (o PageRulesIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRulesIndex200Response) ToMap() (map[string]interface{}, error) {
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

type NullablePageRulesIndex200Response struct {
	value *PageRulesIndex200Response
	isSet bool
}

func (v NullablePageRulesIndex200Response) Get() *PageRulesIndex200Response {
	return v.value
}

func (v *NullablePageRulesIndex200Response) Set(val *PageRulesIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRulesIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRulesIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRulesIndex200Response(val *PageRulesIndex200Response) *NullablePageRulesIndex200Response {
	return &NullablePageRulesIndex200Response{value: val, isSet: true}
}

func (v NullablePageRulesIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRulesIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


