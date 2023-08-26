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

// checks if the UpstreamTimeout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpstreamTimeout{}

// UpstreamTimeout struct for UpstreamTimeout
type UpstreamTimeout struct {
	// Seconds to timeout
	ConnectTimeout *int32 `json:"connect_timeout,omitempty"`
	// Seconds to timeout
	ReadTimeout *int32 `json:"read_timeout,omitempty"`
	// Seconds to timeout
	SendTimeout *int32 `json:"send_timeout,omitempty"`
}

// NewUpstreamTimeout instantiates a new UpstreamTimeout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstreamTimeout() *UpstreamTimeout {
	this := UpstreamTimeout{}
	var connectTimeout int32 = 30
	this.ConnectTimeout = &connectTimeout
	var readTimeout int32 = 100
	this.ReadTimeout = &readTimeout
	var sendTimeout int32 = 300
	this.SendTimeout = &sendTimeout
	return &this
}

// NewUpstreamTimeoutWithDefaults instantiates a new UpstreamTimeout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamTimeoutWithDefaults() *UpstreamTimeout {
	this := UpstreamTimeout{}
	var connectTimeout int32 = 30
	this.ConnectTimeout = &connectTimeout
	var readTimeout int32 = 100
	this.ReadTimeout = &readTimeout
	var sendTimeout int32 = 300
	this.SendTimeout = &sendTimeout
	return &this
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *UpstreamTimeout) GetConnectTimeout() int32 {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret int32
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamTimeout) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *UpstreamTimeout) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given int32 and assigns it to the ConnectTimeout field.
func (o *UpstreamTimeout) SetConnectTimeout(v int32) {
	o.ConnectTimeout = &v
}

// GetReadTimeout returns the ReadTimeout field value if set, zero value otherwise.
func (o *UpstreamTimeout) GetReadTimeout() int32 {
	if o == nil || IsNil(o.ReadTimeout) {
		var ret int32
		return ret
	}
	return *o.ReadTimeout
}

// GetReadTimeoutOk returns a tuple with the ReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamTimeout) GetReadTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ReadTimeout) {
		return nil, false
	}
	return o.ReadTimeout, true
}

// HasReadTimeout returns a boolean if a field has been set.
func (o *UpstreamTimeout) HasReadTimeout() bool {
	if o != nil && !IsNil(o.ReadTimeout) {
		return true
	}

	return false
}

// SetReadTimeout gets a reference to the given int32 and assigns it to the ReadTimeout field.
func (o *UpstreamTimeout) SetReadTimeout(v int32) {
	o.ReadTimeout = &v
}

// GetSendTimeout returns the SendTimeout field value if set, zero value otherwise.
func (o *UpstreamTimeout) GetSendTimeout() int32 {
	if o == nil || IsNil(o.SendTimeout) {
		var ret int32
		return ret
	}
	return *o.SendTimeout
}

// GetSendTimeoutOk returns a tuple with the SendTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamTimeout) GetSendTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SendTimeout) {
		return nil, false
	}
	return o.SendTimeout, true
}

// HasSendTimeout returns a boolean if a field has been set.
func (o *UpstreamTimeout) HasSendTimeout() bool {
	if o != nil && !IsNil(o.SendTimeout) {
		return true
	}

	return false
}

// SetSendTimeout gets a reference to the given int32 and assigns it to the SendTimeout field.
func (o *UpstreamTimeout) SetSendTimeout(v int32) {
	o.SendTimeout = &v
}

func (o UpstreamTimeout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpstreamTimeout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connect_timeout"] = o.ConnectTimeout
	}
	if !IsNil(o.ReadTimeout) {
		toSerialize["read_timeout"] = o.ReadTimeout
	}
	if !IsNil(o.SendTimeout) {
		toSerialize["send_timeout"] = o.SendTimeout
	}
	return toSerialize, nil
}

type NullableUpstreamTimeout struct {
	value *UpstreamTimeout
	isSet bool
}

func (v NullableUpstreamTimeout) Get() *UpstreamTimeout {
	return v.value
}

func (v *NullableUpstreamTimeout) Set(val *UpstreamTimeout) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstreamTimeout) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstreamTimeout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstreamTimeout(val *UpstreamTimeout) *NullableUpstreamTimeout {
	return &NullableUpstreamTimeout{value: val, isSet: true}
}

func (v NullableUpstreamTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstreamTimeout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


