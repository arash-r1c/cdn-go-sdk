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

// checks if the PageRuleImageResize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRuleImageResize{}

// PageRuleImageResize struct for PageRuleImageResize
type PageRuleImageResize struct {
	Status *string `json:"status,omitempty"`
	HeightBy *string `json:"height_by,omitempty"`
	WidthBy *string `json:"width_by,omitempty"`
}

// NewPageRuleImageResize instantiates a new PageRuleImageResize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRuleImageResize() *PageRuleImageResize {
	this := PageRuleImageResize{}
	var status string = "off"
	this.Status = &status
	var heightBy string = "height"
	this.HeightBy = &heightBy
	var widthBy string = "width"
	this.WidthBy = &widthBy
	return &this
}

// NewPageRuleImageResizeWithDefaults instantiates a new PageRuleImageResize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRuleImageResizeWithDefaults() *PageRuleImageResize {
	this := PageRuleImageResize{}
	var heightBy string = "height"
	this.HeightBy = &heightBy
	var widthBy string = "width"
	this.WidthBy = &widthBy
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PageRuleImageResize) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleImageResize) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PageRuleImageResize) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PageRuleImageResize) SetStatus(v string) {
	o.Status = &v
}

// GetHeightBy returns the HeightBy field value if set, zero value otherwise.
func (o *PageRuleImageResize) GetHeightBy() string {
	if o == nil || IsNil(o.HeightBy) {
		var ret string
		return ret
	}
	return *o.HeightBy
}

// GetHeightByOk returns a tuple with the HeightBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleImageResize) GetHeightByOk() (*string, bool) {
	if o == nil || IsNil(o.HeightBy) {
		return nil, false
	}
	return o.HeightBy, true
}

// HasHeightBy returns a boolean if a field has been set.
func (o *PageRuleImageResize) HasHeightBy() bool {
	if o != nil && !IsNil(o.HeightBy) {
		return true
	}

	return false
}

// SetHeightBy gets a reference to the given string and assigns it to the HeightBy field.
func (o *PageRuleImageResize) SetHeightBy(v string) {
	o.HeightBy = &v
}

// GetWidthBy returns the WidthBy field value if set, zero value otherwise.
func (o *PageRuleImageResize) GetWidthBy() string {
	if o == nil || IsNil(o.WidthBy) {
		var ret string
		return ret
	}
	return *o.WidthBy
}

// GetWidthByOk returns a tuple with the WidthBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleImageResize) GetWidthByOk() (*string, bool) {
	if o == nil || IsNil(o.WidthBy) {
		return nil, false
	}
	return o.WidthBy, true
}

// HasWidthBy returns a boolean if a field has been set.
func (o *PageRuleImageResize) HasWidthBy() bool {
	if o != nil && !IsNil(o.WidthBy) {
		return true
	}

	return false
}

// SetWidthBy gets a reference to the given string and assigns it to the WidthBy field.
func (o *PageRuleImageResize) SetWidthBy(v string) {
	o.WidthBy = &v
}

func (o PageRuleImageResize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRuleImageResize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.HeightBy) {
		toSerialize["height_by"] = o.HeightBy
	}
	if !IsNil(o.WidthBy) {
		toSerialize["width_by"] = o.WidthBy
	}
	return toSerialize, nil
}

type NullablePageRuleImageResize struct {
	value *PageRuleImageResize
	isSet bool
}

func (v NullablePageRuleImageResize) Get() *PageRuleImageResize {
	return v.value
}

func (v *NullablePageRuleImageResize) Set(val *PageRuleImageResize) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRuleImageResize) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRuleImageResize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRuleImageResize(val *PageRuleImageResize) *NullablePageRuleImageResize {
	return &NullablePageRuleImageResize{value: val, isSet: true}
}

func (v NullablePageRuleImageResize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRuleImageResize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


