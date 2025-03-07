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

// checks if the LogForwarderDatadogConnectionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderDatadogConnectionType{}

// LogForwarderDatadogConnectionType Datadog connection
type LogForwarderDatadogConnectionType struct {
	SampleRate *int32 `json:"sample_rate,omitempty"`
	Url *string `json:"url,omitempty"`
	ApiKey *string `json:"api_key,omitempty"`
	AppKey *string `json:"app_key,omitempty"`
	FlushInterval *int32 `json:"flush_interval,omitempty"`
	BufferSize *int32 `json:"buffer_size,omitempty"`
}

// NewLogForwarderDatadogConnectionType instantiates a new LogForwarderDatadogConnectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderDatadogConnectionType() *LogForwarderDatadogConnectionType {
	this := LogForwarderDatadogConnectionType{}
	return &this
}

// NewLogForwarderDatadogConnectionTypeWithDefaults instantiates a new LogForwarderDatadogConnectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderDatadogConnectionTypeWithDefaults() *LogForwarderDatadogConnectionType {
	this := LogForwarderDatadogConnectionType{}
	return &this
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *LogForwarderDatadogConnectionType) GetSampleRate() int32 {
	if o == nil || IsNil(o.SampleRate) {
		var ret int32
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDatadogConnectionType) GetSampleRateOk() (*int32, bool) {
	if o == nil || IsNil(o.SampleRate) {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *LogForwarderDatadogConnectionType) HasSampleRate() bool {
	if o != nil && !IsNil(o.SampleRate) {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given int32 and assigns it to the SampleRate field.
func (o *LogForwarderDatadogConnectionType) SetSampleRate(v int32) {
	o.SampleRate = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LogForwarderDatadogConnectionType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDatadogConnectionType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LogForwarderDatadogConnectionType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LogForwarderDatadogConnectionType) SetUrl(v string) {
	o.Url = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *LogForwarderDatadogConnectionType) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDatadogConnectionType) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *LogForwarderDatadogConnectionType) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *LogForwarderDatadogConnectionType) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAppKey returns the AppKey field value if set, zero value otherwise.
func (o *LogForwarderDatadogConnectionType) GetAppKey() string {
	if o == nil || IsNil(o.AppKey) {
		var ret string
		return ret
	}
	return *o.AppKey
}

// GetAppKeyOk returns a tuple with the AppKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDatadogConnectionType) GetAppKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AppKey) {
		return nil, false
	}
	return o.AppKey, true
}

// HasAppKey returns a boolean if a field has been set.
func (o *LogForwarderDatadogConnectionType) HasAppKey() bool {
	if o != nil && !IsNil(o.AppKey) {
		return true
	}

	return false
}

// SetAppKey gets a reference to the given string and assigns it to the AppKey field.
func (o *LogForwarderDatadogConnectionType) SetAppKey(v string) {
	o.AppKey = &v
}

// GetFlushInterval returns the FlushInterval field value if set, zero value otherwise.
func (o *LogForwarderDatadogConnectionType) GetFlushInterval() int32 {
	if o == nil || IsNil(o.FlushInterval) {
		var ret int32
		return ret
	}
	return *o.FlushInterval
}

// GetFlushIntervalOk returns a tuple with the FlushInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDatadogConnectionType) GetFlushIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.FlushInterval) {
		return nil, false
	}
	return o.FlushInterval, true
}

// HasFlushInterval returns a boolean if a field has been set.
func (o *LogForwarderDatadogConnectionType) HasFlushInterval() bool {
	if o != nil && !IsNil(o.FlushInterval) {
		return true
	}

	return false
}

// SetFlushInterval gets a reference to the given int32 and assigns it to the FlushInterval field.
func (o *LogForwarderDatadogConnectionType) SetFlushInterval(v int32) {
	o.FlushInterval = &v
}

// GetBufferSize returns the BufferSize field value if set, zero value otherwise.
func (o *LogForwarderDatadogConnectionType) GetBufferSize() int32 {
	if o == nil || IsNil(o.BufferSize) {
		var ret int32
		return ret
	}
	return *o.BufferSize
}

// GetBufferSizeOk returns a tuple with the BufferSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDatadogConnectionType) GetBufferSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BufferSize) {
		return nil, false
	}
	return o.BufferSize, true
}

// HasBufferSize returns a boolean if a field has been set.
func (o *LogForwarderDatadogConnectionType) HasBufferSize() bool {
	if o != nil && !IsNil(o.BufferSize) {
		return true
	}

	return false
}

// SetBufferSize gets a reference to the given int32 and assigns it to the BufferSize field.
func (o *LogForwarderDatadogConnectionType) SetBufferSize(v int32) {
	o.BufferSize = &v
}

func (o LogForwarderDatadogConnectionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderDatadogConnectionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SampleRate) {
		toSerialize["sample_rate"] = o.SampleRate
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.AppKey) {
		toSerialize["app_key"] = o.AppKey
	}
	if !IsNil(o.FlushInterval) {
		toSerialize["flush_interval"] = o.FlushInterval
	}
	if !IsNil(o.BufferSize) {
		toSerialize["buffer_size"] = o.BufferSize
	}
	return toSerialize, nil
}

type NullableLogForwarderDatadogConnectionType struct {
	value *LogForwarderDatadogConnectionType
	isSet bool
}

func (v NullableLogForwarderDatadogConnectionType) Get() *LogForwarderDatadogConnectionType {
	return v.value
}

func (v *NullableLogForwarderDatadogConnectionType) Set(val *LogForwarderDatadogConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderDatadogConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderDatadogConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderDatadogConnectionType(val *LogForwarderDatadogConnectionType) *NullableLogForwarderDatadogConnectionType {
	return &NullableLogForwarderDatadogConnectionType{value: val, isSet: true}
}

func (v NullableLogForwarderDatadogConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderDatadogConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


