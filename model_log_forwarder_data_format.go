/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
	"fmt"
)

// LogForwarderDataFormat - struct for LogForwarderDataFormat
type LogForwarderDataFormat struct {
	LogForwarderAccessLogType *LogForwarderAccessLogType
	LogForwarderDNSType *LogForwarderDNSType
	LogForwarderErrorType *LogForwarderErrorType
	LogForwarderWAFType *LogForwarderWAFType
}

// LogForwarderAccessLogTypeAsLogForwarderDataFormat is a convenience function that returns LogForwarderAccessLogType wrapped in LogForwarderDataFormat
func LogForwarderAccessLogTypeAsLogForwarderDataFormat(v *LogForwarderAccessLogType) LogForwarderDataFormat {
	return LogForwarderDataFormat{
		LogForwarderAccessLogType: v,
	}
}

// LogForwarderDNSTypeAsLogForwarderDataFormat is a convenience function that returns LogForwarderDNSType wrapped in LogForwarderDataFormat
func LogForwarderDNSTypeAsLogForwarderDataFormat(v *LogForwarderDNSType) LogForwarderDataFormat {
	return LogForwarderDataFormat{
		LogForwarderDNSType: v,
	}
}

// LogForwarderErrorTypeAsLogForwarderDataFormat is a convenience function that returns LogForwarderErrorType wrapped in LogForwarderDataFormat
func LogForwarderErrorTypeAsLogForwarderDataFormat(v *LogForwarderErrorType) LogForwarderDataFormat {
	return LogForwarderDataFormat{
		LogForwarderErrorType: v,
	}
}

// LogForwarderWAFTypeAsLogForwarderDataFormat is a convenience function that returns LogForwarderWAFType wrapped in LogForwarderDataFormat
func LogForwarderWAFTypeAsLogForwarderDataFormat(v *LogForwarderWAFType) LogForwarderDataFormat {
	return LogForwarderDataFormat{
		LogForwarderWAFType: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogForwarderDataFormat) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogForwarderAccessLogType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderAccessLogType)
	if err == nil {
		jsonLogForwarderAccessLogType, _ := json.Marshal(dst.LogForwarderAccessLogType)
		if string(jsonLogForwarderAccessLogType) == "{}" { // empty struct
			dst.LogForwarderAccessLogType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderAccessLogType = nil
	}

	// try to unmarshal data into LogForwarderDNSType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderDNSType)
	if err == nil {
		jsonLogForwarderDNSType, _ := json.Marshal(dst.LogForwarderDNSType)
		if string(jsonLogForwarderDNSType) == "{}" { // empty struct
			dst.LogForwarderDNSType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderDNSType = nil
	}

	// try to unmarshal data into LogForwarderErrorType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderErrorType)
	if err == nil {
		jsonLogForwarderErrorType, _ := json.Marshal(dst.LogForwarderErrorType)
		if string(jsonLogForwarderErrorType) == "{}" { // empty struct
			dst.LogForwarderErrorType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderErrorType = nil
	}

	// try to unmarshal data into LogForwarderWAFType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderWAFType)
	if err == nil {
		jsonLogForwarderWAFType, _ := json.Marshal(dst.LogForwarderWAFType)
		if string(jsonLogForwarderWAFType) == "{}" { // empty struct
			dst.LogForwarderWAFType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderWAFType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LogForwarderAccessLogType = nil
		dst.LogForwarderDNSType = nil
		dst.LogForwarderErrorType = nil
		dst.LogForwarderWAFType = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LogForwarderDataFormat)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LogForwarderDataFormat)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogForwarderDataFormat) MarshalJSON() ([]byte, error) {
	if src.LogForwarderAccessLogType != nil {
		return json.Marshal(&src.LogForwarderAccessLogType)
	}

	if src.LogForwarderDNSType != nil {
		return json.Marshal(&src.LogForwarderDNSType)
	}

	if src.LogForwarderErrorType != nil {
		return json.Marshal(&src.LogForwarderErrorType)
	}

	if src.LogForwarderWAFType != nil {
		return json.Marshal(&src.LogForwarderWAFType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogForwarderDataFormat) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LogForwarderAccessLogType != nil {
		return obj.LogForwarderAccessLogType
	}

	if obj.LogForwarderDNSType != nil {
		return obj.LogForwarderDNSType
	}

	if obj.LogForwarderErrorType != nil {
		return obj.LogForwarderErrorType
	}

	if obj.LogForwarderWAFType != nil {
		return obj.LogForwarderWAFType
	}

	// all schemas are nil
	return nil
}

type NullableLogForwarderDataFormat struct {
	value *LogForwarderDataFormat
	isSet bool
}

func (v NullableLogForwarderDataFormat) Get() *LogForwarderDataFormat {
	return v.value
}

func (v *NullableLogForwarderDataFormat) Set(val *LogForwarderDataFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderDataFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderDataFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderDataFormat(val *LogForwarderDataFormat) *NullableLogForwarderDataFormat {
	return &NullableLogForwarderDataFormat{value: val, isSet: true}
}

func (v NullableLogForwarderDataFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderDataFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


