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

// checks if the LogForwarderLogglyConnectionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderLogglyConnectionType{}

// LogForwarderLogglyConnectionType Loggly connection
type LogForwarderLogglyConnectionType struct {
	SampleRate *int32 `json:"sample_rate,omitempty"`
	Token *string `json:"token,omitempty"`
	Url *string `json:"url,omitempty"`
	FlushInterval *int32 `json:"flush_interval,omitempty"`
	BufferSize *int32 `json:"buffer_size,omitempty"`
}

// NewLogForwarderLogglyConnectionType instantiates a new LogForwarderLogglyConnectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderLogglyConnectionType() *LogForwarderLogglyConnectionType {
	this := LogForwarderLogglyConnectionType{}
	return &this
}

// NewLogForwarderLogglyConnectionTypeWithDefaults instantiates a new LogForwarderLogglyConnectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderLogglyConnectionTypeWithDefaults() *LogForwarderLogglyConnectionType {
	this := LogForwarderLogglyConnectionType{}
	return &this
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *LogForwarderLogglyConnectionType) GetSampleRate() int32 {
	if o == nil || IsNil(o.SampleRate) {
		var ret int32
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderLogglyConnectionType) GetSampleRateOk() (*int32, bool) {
	if o == nil || IsNil(o.SampleRate) {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *LogForwarderLogglyConnectionType) HasSampleRate() bool {
	if o != nil && !IsNil(o.SampleRate) {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given int32 and assigns it to the SampleRate field.
func (o *LogForwarderLogglyConnectionType) SetSampleRate(v int32) {
	o.SampleRate = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LogForwarderLogglyConnectionType) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderLogglyConnectionType) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LogForwarderLogglyConnectionType) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LogForwarderLogglyConnectionType) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LogForwarderLogglyConnectionType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderLogglyConnectionType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LogForwarderLogglyConnectionType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LogForwarderLogglyConnectionType) SetUrl(v string) {
	o.Url = &v
}

// GetFlushInterval returns the FlushInterval field value if set, zero value otherwise.
func (o *LogForwarderLogglyConnectionType) GetFlushInterval() int32 {
	if o == nil || IsNil(o.FlushInterval) {
		var ret int32
		return ret
	}
	return *o.FlushInterval
}

// GetFlushIntervalOk returns a tuple with the FlushInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderLogglyConnectionType) GetFlushIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.FlushInterval) {
		return nil, false
	}
	return o.FlushInterval, true
}

// HasFlushInterval returns a boolean if a field has been set.
func (o *LogForwarderLogglyConnectionType) HasFlushInterval() bool {
	if o != nil && !IsNil(o.FlushInterval) {
		return true
	}

	return false
}

// SetFlushInterval gets a reference to the given int32 and assigns it to the FlushInterval field.
func (o *LogForwarderLogglyConnectionType) SetFlushInterval(v int32) {
	o.FlushInterval = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *LogForwarderLogglyConnectionType) GetBufferSize() int32 {
	if o == nil || IsNil(o.BufferSize) {
		var ret int32
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderLogglyConnectionType) GetBufferSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *LogForwarderLogglyConnectionType) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given int32 and assigns it to the BufferSize field.
func (o *LogForwarderLogglyConnectionType) SetBufferSize(v int32) {
	o.BufferSize = &v
}

func (o LogForwarderLogglyConnectionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderLogglyConnectionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SampleRate) {
		toSerialize["sample_rate"] = o.SampleRate
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.FlushInterval) {
		toSerialize["flush_interval"] = o.FlushInterval
	}
	if !IsNil(o.BufferSize) {
		toSerialize["buffer_size"] = o.BufferSize
	}
	return toSerialize, nil
}

type NullableLogForwarderLogglyConnectionType struct {
	value *LogForwarderLogglyConnectionType
	isSet bool
}

func (v NullableLogForwarderLogglyConnectionType) Get() *LogForwarderLogglyConnectionType {
	return v.value
}

func (v *NullableLogForwarderLogglyConnectionType) Set(val *LogForwarderLogglyConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderLogglyConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderLogglyConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderLogglyConnectionType(val *LogForwarderLogglyConnectionType) *NullableLogForwarderLogglyConnectionType {
	return &NullableLogForwarderLogglyConnectionType{value: val, isSet: true}
}

func (v NullableLogForwarderLogglyConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderLogglyConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


