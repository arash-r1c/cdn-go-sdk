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

// checks if the LogForwarderErrorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderErrorType{}

// LogForwarderErrorType Error log type
type LogForwarderErrorType struct {
	ClientIp *bool `json:"client_ip,omitempty"`
	UpstreamProto *bool `json:"upstream_proto,omitempty"`
	UpstreamUri *bool `json:"upstream_uri,omitempty"`
	UpstreamPort *bool `json:"upstream_port,omitempty"`
	UpstreamIp *bool `json:"upstream_ip,omitempty"`
	DomainName *bool `json:"domain_name,omitempty"`
	HttpVersion *bool `json:"http_version,omitempty"`
	RequestMethod *bool `json:"request_method,omitempty"`
	RequestUri *bool `json:"request_uri,omitempty"`
	RealTimestamp *bool `json:"real_timestamp,omitempty"`
	ErrorMessage *bool `json:"error_message,omitempty"`
	PopSite *bool `json:"pop_site,omitempty"`
	RequestId *bool `json:"request_id,omitempty"`
}

// NewLogForwarderErrorType instantiates a new LogForwarderErrorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderErrorType() *LogForwarderErrorType {
	this := LogForwarderErrorType{}
	return &this
}

// NewLogForwarderErrorTypeWithDefaults instantiates a new LogForwarderErrorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderErrorTypeWithDefaults() *LogForwarderErrorType {
	this := LogForwarderErrorType{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetClientIp() bool {
	if o == nil || IsNil(o.ClientIp) {
		var ret bool
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetClientIpOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasClientIp() bool {
	if o != nil && !IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given bool and assigns it to the ClientIp field.
func (o *LogForwarderErrorType) SetClientIp(v bool) {
	o.ClientIp = &v
}

// GetUpstreamProto returns the UpstreamProto field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetUpstreamProto() bool {
	if o == nil || IsNil(o.UpstreamProto) {
		var ret bool
		return ret
	}
	return *o.UpstreamProto
}

// GetUpstreamProtoOk returns a tuple with the UpstreamProto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetUpstreamProtoOk() (*bool, bool) {
	if o == nil || IsNil(o.UpstreamProto) {
		return nil, false
	}
	return o.UpstreamProto, true
}

// HasUpstreamProto returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasUpstreamProto() bool {
	if o != nil && !IsNil(o.UpstreamProto) {
		return true
	}

	return false
}

// SetUpstreamProto gets a reference to the given bool and assigns it to the UpstreamProto field.
func (o *LogForwarderErrorType) SetUpstreamProto(v bool) {
	o.UpstreamProto = &v
}

// GetUpstreamUri returns the UpstreamUri field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetUpstreamUri() bool {
	if o == nil || IsNil(o.UpstreamUri) {
		var ret bool
		return ret
	}
	return *o.UpstreamUri
}

// GetUpstreamUriOk returns a tuple with the UpstreamUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetUpstreamUriOk() (*bool, bool) {
	if o == nil || IsNil(o.UpstreamUri) {
		return nil, false
	}
	return o.UpstreamUri, true
}

// HasUpstreamUri returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasUpstreamUri() bool {
	if o != nil && !IsNil(o.UpstreamUri) {
		return true
	}

	return false
}

// SetUpstreamUri gets a reference to the given bool and assigns it to the UpstreamUri field.
func (o *LogForwarderErrorType) SetUpstreamUri(v bool) {
	o.UpstreamUri = &v
}

// GetUpstreamPort returns the UpstreamPort field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetUpstreamPort() bool {
	if o == nil || IsNil(o.UpstreamPort) {
		var ret bool
		return ret
	}
	return *o.UpstreamPort
}

// GetUpstreamPortOk returns a tuple with the UpstreamPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetUpstreamPortOk() (*bool, bool) {
	if o == nil || IsNil(o.UpstreamPort) {
		return nil, false
	}
	return o.UpstreamPort, true
}

// HasUpstreamPort returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasUpstreamPort() bool {
	if o != nil && !IsNil(o.UpstreamPort) {
		return true
	}

	return false
}

// SetUpstreamPort gets a reference to the given bool and assigns it to the UpstreamPort field.
func (o *LogForwarderErrorType) SetUpstreamPort(v bool) {
	o.UpstreamPort = &v
}

// GetUpstreamIp returns the UpstreamIp field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetUpstreamIp() bool {
	if o == nil || IsNil(o.UpstreamIp) {
		var ret bool
		return ret
	}
	return *o.UpstreamIp
}

// GetUpstreamIpOk returns a tuple with the UpstreamIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetUpstreamIpOk() (*bool, bool) {
	if o == nil || IsNil(o.UpstreamIp) {
		return nil, false
	}
	return o.UpstreamIp, true
}

// HasUpstreamIp returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasUpstreamIp() bool {
	if o != nil && !IsNil(o.UpstreamIp) {
		return true
	}

	return false
}

// SetUpstreamIp gets a reference to the given bool and assigns it to the UpstreamIp field.
func (o *LogForwarderErrorType) SetUpstreamIp(v bool) {
	o.UpstreamIp = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetDomainName() bool {
	if o == nil || IsNil(o.DomainName) {
		var ret bool
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetDomainNameOk() (*bool, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given bool and assigns it to the DomainName field.
func (o *LogForwarderErrorType) SetDomainName(v bool) {
	o.DomainName = &v
}

// GetHttpVersion returns the HttpVersion field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetHttpVersion() bool {
	if o == nil || IsNil(o.HttpVersion) {
		var ret bool
		return ret
	}
	return *o.HttpVersion
}

// GetHttpVersionOk returns a tuple with the HttpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetHttpVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.HttpVersion) {
		return nil, false
	}
	return o.HttpVersion, true
}

// HasHttpVersion returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasHttpVersion() bool {
	if o != nil && !IsNil(o.HttpVersion) {
		return true
	}

	return false
}

// SetHttpVersion gets a reference to the given bool and assigns it to the HttpVersion field.
func (o *LogForwarderErrorType) SetHttpVersion(v bool) {
	o.HttpVersion = &v
}

// GetRequestMethod returns the RequestMethod field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetRequestMethod() bool {
	if o == nil || IsNil(o.RequestMethod) {
		var ret bool
		return ret
	}
	return *o.RequestMethod
}

// GetRequestMethodOk returns a tuple with the RequestMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetRequestMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestMethod) {
		return nil, false
	}
	return o.RequestMethod, true
}

// HasRequestMethod returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasRequestMethod() bool {
	if o != nil && !IsNil(o.RequestMethod) {
		return true
	}

	return false
}

// SetRequestMethod gets a reference to the given bool and assigns it to the RequestMethod field.
func (o *LogForwarderErrorType) SetRequestMethod(v bool) {
	o.RequestMethod = &v
}

// GetRequestUri returns the RequestUri field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetRequestUri() bool {
	if o == nil || IsNil(o.RequestUri) {
		var ret bool
		return ret
	}
	return *o.RequestUri
}

// GetRequestUriOk returns a tuple with the RequestUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetRequestUriOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestUri) {
		return nil, false
	}
	return o.RequestUri, true
}

// HasRequestUri returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasRequestUri() bool {
	if o != nil && !IsNil(o.RequestUri) {
		return true
	}

	return false
}

// SetRequestUri gets a reference to the given bool and assigns it to the RequestUri field.
func (o *LogForwarderErrorType) SetRequestUri(v bool) {
	o.RequestUri = &v
}

// GetRealTimestamp returns the RealTimestamp field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetRealTimestamp() bool {
	if o == nil || IsNil(o.RealTimestamp) {
		var ret bool
		return ret
	}
	return *o.RealTimestamp
}

// GetRealTimestampOk returns a tuple with the RealTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetRealTimestampOk() (*bool, bool) {
	if o == nil || IsNil(o.RealTimestamp) {
		return nil, false
	}
	return o.RealTimestamp, true
}

// HasRealTimestamp returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasRealTimestamp() bool {
	if o != nil && !IsNil(o.RealTimestamp) {
		return true
	}

	return false
}

// SetRealTimestamp gets a reference to the given bool and assigns it to the RealTimestamp field.
func (o *LogForwarderErrorType) SetRealTimestamp(v bool) {
	o.RealTimestamp = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetErrorMessage() bool {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret bool
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetErrorMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given bool and assigns it to the ErrorMessage field.
func (o *LogForwarderErrorType) SetErrorMessage(v bool) {
	o.ErrorMessage = &v
}

// GetPopSite returns the PopSite field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetPopSite() bool {
	if o == nil || IsNil(o.PopSite) {
		var ret bool
		return ret
	}
	return *o.PopSite
}

// GetPopSiteOk returns a tuple with the PopSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetPopSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.PopSite) {
		return nil, false
	}
	return o.PopSite, true
}

// HasPopSite returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasPopSite() bool {
	if o != nil && !IsNil(o.PopSite) {
		return true
	}

	return false
}

// SetPopSite gets a reference to the given bool and assigns it to the PopSite field.
func (o *LogForwarderErrorType) SetPopSite(v bool) {
	o.PopSite = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *LogForwarderErrorType) GetRequestId() bool {
	if o == nil || IsNil(o.RequestId) {
		var ret bool
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderErrorType) GetRequestIdOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *LogForwarderErrorType) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given bool and assigns it to the RequestId field.
func (o *LogForwarderErrorType) SetRequestId(v bool) {
	o.RequestId = &v
}

func (o LogForwarderErrorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderErrorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !IsNil(o.UpstreamProto) {
		toSerialize["upstream_proto"] = o.UpstreamProto
	}
	if !IsNil(o.UpstreamUri) {
		toSerialize["upstream_uri"] = o.UpstreamUri
	}
	if !IsNil(o.UpstreamPort) {
		toSerialize["upstream_port"] = o.UpstreamPort
	}
	if !IsNil(o.UpstreamIp) {
		toSerialize["upstream_ip"] = o.UpstreamIp
	}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.HttpVersion) {
		toSerialize["http_version"] = o.HttpVersion
	}
	if !IsNil(o.RequestMethod) {
		toSerialize["request_method"] = o.RequestMethod
	}
	if !IsNil(o.RequestUri) {
		toSerialize["request_uri"] = o.RequestUri
	}
	if !IsNil(o.RealTimestamp) {
		toSerialize["real_timestamp"] = o.RealTimestamp
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.PopSite) {
		toSerialize["pop_site"] = o.PopSite
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableLogForwarderErrorType struct {
	value *LogForwarderErrorType
	isSet bool
}

func (v NullableLogForwarderErrorType) Get() *LogForwarderErrorType {
	return v.value
}

func (v *NullableLogForwarderErrorType) Set(val *LogForwarderErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderErrorType(val *LogForwarderErrorType) *NullableLogForwarderErrorType {
	return &NullableLogForwarderErrorType{value: val, isSet: true}
}

func (v NullableLogForwarderErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


