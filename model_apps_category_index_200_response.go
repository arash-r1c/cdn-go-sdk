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

// checks if the AppsCategoryIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsCategoryIndex200Response{}

// AppsCategoryIndex200Response struct for AppsCategoryIndex200Response
type AppsCategoryIndex200Response struct {
	Data []ApplicationCategory `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta *PaginatedResponseMeta `json:"meta,omitempty"`
}

// NewAppsCategoryIndex200Response instantiates a new AppsCategoryIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsCategoryIndex200Response() *AppsCategoryIndex200Response {
	this := AppsCategoryIndex200Response{}
	return &this
}

// NewAppsCategoryIndex200ResponseWithDefaults instantiates a new AppsCategoryIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsCategoryIndex200ResponseWithDefaults() *AppsCategoryIndex200Response {
	this := AppsCategoryIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppsCategoryIndex200Response) GetData() []ApplicationCategory {
	if o == nil || IsNil(o.Data) {
		var ret []ApplicationCategory
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCategoryIndex200Response) GetDataOk() ([]ApplicationCategory, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppsCategoryIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ApplicationCategory and assigns it to the Data field.
func (o *AppsCategoryIndex200Response) SetData(v []ApplicationCategory) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppsCategoryIndex200Response) GetLinks() PaginatedResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCategoryIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppsCategoryIndex200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedResponseLinks and assigns it to the Links field.
func (o *AppsCategoryIndex200Response) SetLinks(v PaginatedResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppsCategoryIndex200Response) GetMeta() PaginatedResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatedResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCategoryIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppsCategoryIndex200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatedResponseMeta and assigns it to the Meta field.
func (o *AppsCategoryIndex200Response) SetMeta(v PaginatedResponseMeta) {
	o.Meta = &v
}

func (o AppsCategoryIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsCategoryIndex200Response) ToMap() (map[string]interface{}, error) {
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

type NullableAppsCategoryIndex200Response struct {
	value *AppsCategoryIndex200Response
	isSet bool
}

func (v NullableAppsCategoryIndex200Response) Get() *AppsCategoryIndex200Response {
	return v.value
}

func (v *NullableAppsCategoryIndex200Response) Set(val *AppsCategoryIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsCategoryIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsCategoryIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsCategoryIndex200Response(val *AppsCategoryIndex200Response) *NullableAppsCategoryIndex200Response {
	return &NullableAppsCategoryIndex200Response{value: val, isSet: true}
}

func (v NullableAppsCategoryIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsCategoryIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


