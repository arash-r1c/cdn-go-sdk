/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.106.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"fmt"
)

// PlansIndexDomainParameter struct for PlansIndexDomainParameter
type PlansIndexDomainParameter struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PlansIndexDomainParameter) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PlansIndexDomainParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PlansIndexDomainParameter) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePlansIndexDomainParameter struct {
	value *PlansIndexDomainParameter
	isSet bool
}

func (v NullablePlansIndexDomainParameter) Get() *PlansIndexDomainParameter {
	return v.value
}

func (v *NullablePlansIndexDomainParameter) Set(val *PlansIndexDomainParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePlansIndexDomainParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePlansIndexDomainParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlansIndexDomainParameter(val *PlansIndexDomainParameter) *NullablePlansIndexDomainParameter {
	return &NullablePlansIndexDomainParameter{value: val, isSet: true}
}

func (v NullablePlansIndexDomainParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlansIndexDomainParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


