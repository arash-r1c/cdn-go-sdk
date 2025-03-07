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

// checks if the TransportLayerProxyTrafficCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyTrafficCharts{}

// TransportLayerProxyTrafficCharts struct for TransportLayerProxyTrafficCharts
type TransportLayerProxyTrafficCharts struct {
	Traffics *TransportLayerProxyTrafficChartsTraffics `json:"traffics,omitempty"`
}

// NewTransportLayerProxyTrafficCharts instantiates a new TransportLayerProxyTrafficCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyTrafficCharts() *TransportLayerProxyTrafficCharts {
	this := TransportLayerProxyTrafficCharts{}
	return &this
}

// NewTransportLayerProxyTrafficChartsWithDefaults instantiates a new TransportLayerProxyTrafficCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyTrafficChartsWithDefaults() *TransportLayerProxyTrafficCharts {
	this := TransportLayerProxyTrafficCharts{}
	return &this
}

// GetTraffics returns the Traffics field value if set, zero value otherwise.
func (o *TransportLayerProxyTrafficCharts) GetTraffics() TransportLayerProxyTrafficChartsTraffics {
	if o == nil || IsNil(o.Traffics) {
		var ret TransportLayerProxyTrafficChartsTraffics
		return ret
	}
	return *o.Traffics
}

// GetTrafficsOk returns a tuple with the Traffics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyTrafficCharts) GetTrafficsOk() (*TransportLayerProxyTrafficChartsTraffics, bool) {
	if o == nil || IsNil(o.Traffics) {
		return nil, false
	}
	return o.Traffics, true
}

// HasTraffics returns a boolean if a field has been set.
func (o *TransportLayerProxyTrafficCharts) HasTraffics() bool {
	if o != nil && !IsNil(o.Traffics) {
		return true
	}

	return false
}

// SetTraffics gets a reference to the given TransportLayerProxyTrafficChartsTraffics and assigns it to the Traffics field.
func (o *TransportLayerProxyTrafficCharts) SetTraffics(v TransportLayerProxyTrafficChartsTraffics) {
	o.Traffics = &v
}

func (o TransportLayerProxyTrafficCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyTrafficCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Traffics) {
		toSerialize["traffics"] = o.Traffics
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyTrafficCharts struct {
	value *TransportLayerProxyTrafficCharts
	isSet bool
}

func (v NullableTransportLayerProxyTrafficCharts) Get() *TransportLayerProxyTrafficCharts {
	return v.value
}

func (v *NullableTransportLayerProxyTrafficCharts) Set(val *TransportLayerProxyTrafficCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyTrafficCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyTrafficCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyTrafficCharts(val *TransportLayerProxyTrafficCharts) *NullableTransportLayerProxyTrafficCharts {
	return &NullableTransportLayerProxyTrafficCharts{value: val, isSet: true}
}

func (v NullableTransportLayerProxyTrafficCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyTrafficCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


