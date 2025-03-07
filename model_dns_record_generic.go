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

// DnsRecordGeneric struct for DnsRecordGeneric
type DnsRecordGeneric struct {
	DnsRecordGenericArrayValue *DnsRecordGenericArrayValue
	DnsRecordGenericObjectValue *DnsRecordGenericObjectValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnsRecordGeneric) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DnsRecordGenericArrayValue
	err = json.Unmarshal(data, &dst.DnsRecordGenericArrayValue);
	if err == nil {
		jsonDnsRecordGenericArrayValue, _ := json.Marshal(dst.DnsRecordGenericArrayValue)
		if string(jsonDnsRecordGenericArrayValue) == "{}" { // empty struct
			dst.DnsRecordGenericArrayValue = nil
		} else {
			return nil // data stored in dst.DnsRecordGenericArrayValue, return on the first match
		}
	} else {
		dst.DnsRecordGenericArrayValue = nil
	}

	// try to unmarshal JSON data into DnsRecordGenericObjectValue
	err = json.Unmarshal(data, &dst.DnsRecordGenericObjectValue);
	if err == nil {
		jsonDnsRecordGenericObjectValue, _ := json.Marshal(dst.DnsRecordGenericObjectValue)
		if string(jsonDnsRecordGenericObjectValue) == "{}" { // empty struct
			dst.DnsRecordGenericObjectValue = nil
		} else {
			return nil // data stored in dst.DnsRecordGenericObjectValue, return on the first match
		}
	} else {
		dst.DnsRecordGenericObjectValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DnsRecordGeneric)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnsRecordGeneric) MarshalJSON() ([]byte, error) {
	if src.DnsRecordGenericArrayValue != nil {
		return json.Marshal(&src.DnsRecordGenericArrayValue)
	}

	if src.DnsRecordGenericObjectValue != nil {
		return json.Marshal(&src.DnsRecordGenericObjectValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnsRecordGeneric struct {
	value *DnsRecordGeneric
	isSet bool
}

func (v NullableDnsRecordGeneric) Get() *DnsRecordGeneric {
	return v.value
}

func (v *NullableDnsRecordGeneric) Set(val *DnsRecordGeneric) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecordGeneric) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecordGeneric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecordGeneric(val *DnsRecordGeneric) *NullableDnsRecordGeneric {
	return &NullableDnsRecordGeneric{value: val, isSet: true}
}

func (v NullableDnsRecordGeneric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecordGeneric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


