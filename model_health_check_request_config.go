/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.115.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"fmt"
)

// HealthCheckRequestConfig - struct for HealthCheckRequestConfig
type HealthCheckRequestConfig struct {
	HttpConfig *HttpConfig
	TcpConfig *TcpConfig
}

// HttpConfigAsHealthCheckRequestConfig is a convenience function that returns HttpConfig wrapped in HealthCheckRequestConfig
func HttpConfigAsHealthCheckRequestConfig(v *HttpConfig) HealthCheckRequestConfig {
	return HealthCheckRequestConfig{
		HttpConfig: v,
	}
}

// TcpConfigAsHealthCheckRequestConfig is a convenience function that returns TcpConfig wrapped in HealthCheckRequestConfig
func TcpConfigAsHealthCheckRequestConfig(v *TcpConfig) HealthCheckRequestConfig {
	return HealthCheckRequestConfig{
		TcpConfig: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *HealthCheckRequestConfig) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into HttpConfig
	err = newStrictDecoder(data).Decode(&dst.HttpConfig)
	if err == nil {
		jsonHttpConfig, _ := json.Marshal(dst.HttpConfig)
		if string(jsonHttpConfig) == "{}" { // empty struct
			dst.HttpConfig = nil
		} else {
			match++
		}
	} else {
		dst.HttpConfig = nil
	}

	// try to unmarshal data into TcpConfig
	err = newStrictDecoder(data).Decode(&dst.TcpConfig)
	if err == nil {
		jsonTcpConfig, _ := json.Marshal(dst.TcpConfig)
		if string(jsonTcpConfig) == "{}" { // empty struct
			dst.TcpConfig = nil
		} else {
			match++
		}
	} else {
		dst.TcpConfig = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.HttpConfig = nil
		dst.TcpConfig = nil

		return fmt.Errorf("data matches more than one schema in oneOf(HealthCheckRequestConfig)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(HealthCheckRequestConfig)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HealthCheckRequestConfig) MarshalJSON() ([]byte, error) {
	if src.HttpConfig != nil {
		return json.Marshal(&src.HttpConfig)
	}

	if src.TcpConfig != nil {
		return json.Marshal(&src.TcpConfig)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HealthCheckRequestConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.HttpConfig != nil {
		return obj.HttpConfig
	}

	if obj.TcpConfig != nil {
		return obj.TcpConfig
	}

	// all schemas are nil
	return nil
}

type NullableHealthCheckRequestConfig struct {
	value *HealthCheckRequestConfig
	isSet bool
}

func (v NullableHealthCheckRequestConfig) Get() *HealthCheckRequestConfig {
	return v.value
}

func (v *NullableHealthCheckRequestConfig) Set(val *HealthCheckRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckRequestConfig(val *HealthCheckRequestConfig) *NullableHealthCheckRequestConfig {
	return &NullableHealthCheckRequestConfig{value: val, isSet: true}
}

func (v NullableHealthCheckRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


