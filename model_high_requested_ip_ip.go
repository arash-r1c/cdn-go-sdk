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

// HighRequestedIpIp struct for HighRequestedIpIp
type HighRequestedIpIp struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *HighRequestedIpIp) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(HighRequestedIpIp)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *HighRequestedIpIp) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableHighRequestedIpIp struct {
	value *HighRequestedIpIp
	isSet bool
}

func (v NullableHighRequestedIpIp) Get() *HighRequestedIpIp {
	return v.value
}

func (v *NullableHighRequestedIpIp) Set(val *HighRequestedIpIp) {
	v.value = val
	v.isSet = true
}

func (v NullableHighRequestedIpIp) IsSet() bool {
	return v.isSet
}

func (v *NullableHighRequestedIpIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighRequestedIpIp(val *HighRequestedIpIp) *NullableHighRequestedIpIp {
	return &NullableHighRequestedIpIp{value: val, isSet: true}
}

func (v NullableHighRequestedIpIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighRequestedIpIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


