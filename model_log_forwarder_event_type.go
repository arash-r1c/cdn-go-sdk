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

// checks if the LogForwarderEventType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderEventType{}

// LogForwarderEventType Event log type
type LogForwarderEventType struct {
	Domain *bool `json:"domain,omitempty"`
	Uuid *bool `json:"uuid,omitempty"`
	Timestamp *bool `json:"timestamp,omitempty"`
	Method *bool `json:"method,omitempty"`
	Scheme *bool `json:"scheme,omitempty"`
	Ip *bool `json:"ip,omitempty"`
	Country *bool `json:"country,omitempty"`
	Status *bool `json:"status,omitempty"`
	ServerIp *bool `json:"server_ip,omitempty"`
	ServerPort *bool `json:"server_port,omitempty"`
	Uri *bool `json:"uri,omitempty"`
	QueryString *bool `json:"query_string,omitempty"`
	Firewall *bool `json:"firewall,omitempty"`
	Proxy *bool `json:"proxy,omitempty"`
	DnsResolver *bool `json:"dns_resolver,omitempty"`
	Ddos *bool `json:"ddos,omitempty"`
	Ratelimit *bool `json:"ratelimit,omitempty"`
	Waf *bool `json:"waf,omitempty"`
}

// NewLogForwarderEventType instantiates a new LogForwarderEventType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderEventType() *LogForwarderEventType {
	this := LogForwarderEventType{}
	return &this
}

// NewLogForwarderEventTypeWithDefaults instantiates a new LogForwarderEventType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderEventTypeWithDefaults() *LogForwarderEventType {
	this := LogForwarderEventType{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetDomain() bool {
	if o == nil || IsNil(o.Domain) {
		var ret bool
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetDomainOk() (*bool, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given bool and assigns it to the Domain field.
func (o *LogForwarderEventType) SetDomain(v bool) {
	o.Domain = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetUuid() bool {
	if o == nil || IsNil(o.Uuid) {
		var ret bool
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetUuidOk() (*bool, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given bool and assigns it to the Uuid field.
func (o *LogForwarderEventType) SetUuid(v bool) {
	o.Uuid = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetTimestamp() bool {
	if o == nil || IsNil(o.Timestamp) {
		var ret bool
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetTimestampOk() (*bool, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given bool and assigns it to the Timestamp field.
func (o *LogForwarderEventType) SetTimestamp(v bool) {
	o.Timestamp = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetMethod() bool {
	if o == nil || IsNil(o.Method) {
		var ret bool
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given bool and assigns it to the Method field.
func (o *LogForwarderEventType) SetMethod(v bool) {
	o.Method = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetScheme() bool {
	if o == nil || IsNil(o.Scheme) {
		var ret bool
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetSchemeOk() (*bool, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given bool and assigns it to the Scheme field.
func (o *LogForwarderEventType) SetScheme(v bool) {
	o.Scheme = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetIp() bool {
	if o == nil || IsNil(o.Ip) {
		var ret bool
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetIpOk() (*bool, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given bool and assigns it to the Ip field.
func (o *LogForwarderEventType) SetIp(v bool) {
	o.Ip = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetCountry() bool {
	if o == nil || IsNil(o.Country) {
		var ret bool
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetCountryOk() (*bool, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given bool and assigns it to the Country field.
func (o *LogForwarderEventType) SetCountry(v bool) {
	o.Country = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LogForwarderEventType) SetStatus(v bool) {
	o.Status = &v
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetServerIp() bool {
	if o == nil || IsNil(o.ServerIp) {
		var ret bool
		return ret
	}
	return *o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetServerIpOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerIp) {
		return nil, false
	}
	return o.ServerIp, true
}

// HasServerIp returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasServerIp() bool {
	if o != nil && !IsNil(o.ServerIp) {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given bool and assigns it to the ServerIp field.
func (o *LogForwarderEventType) SetServerIp(v bool) {
	o.ServerIp = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetServerPort() bool {
	if o == nil || IsNil(o.ServerPort) {
		var ret bool
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetServerPortOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerPort) {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasServerPort() bool {
	if o != nil && !IsNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given bool and assigns it to the ServerPort field.
func (o *LogForwarderEventType) SetServerPort(v bool) {
	o.ServerPort = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetUri() bool {
	if o == nil || IsNil(o.Uri) {
		var ret bool
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetUriOk() (*bool, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given bool and assigns it to the Uri field.
func (o *LogForwarderEventType) SetUri(v bool) {
	o.Uri = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetQueryString() bool {
	if o == nil || IsNil(o.QueryString) {
		var ret bool
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetQueryStringOk() (*bool, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given bool and assigns it to the QueryString field.
func (o *LogForwarderEventType) SetQueryString(v bool) {
	o.QueryString = &v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetFirewall() bool {
	if o == nil || IsNil(o.Firewall) {
		var ret bool
		return ret
	}
	return *o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.Firewall) {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given bool and assigns it to the Firewall field.
func (o *LogForwarderEventType) SetFirewall(v bool) {
	o.Firewall = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetProxy() bool {
	if o == nil || IsNil(o.Proxy) {
		var ret bool
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasProxy() bool {
	if o != nil && !IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given bool and assigns it to the Proxy field.
func (o *LogForwarderEventType) SetProxy(v bool) {
	o.Proxy = &v
}

// GetDnsResolver returns the DnsResolver field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetDnsResolver() bool {
	if o == nil || IsNil(o.DnsResolver) {
		var ret bool
		return ret
	}
	return *o.DnsResolver
}

// GetDnsResolverOk returns a tuple with the DnsResolver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetDnsResolverOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsResolver) {
		return nil, false
	}
	return o.DnsResolver, true
}

// HasDnsResolver returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasDnsResolver() bool {
	if o != nil && !IsNil(o.DnsResolver) {
		return true
	}

	return false
}

// SetDnsResolver gets a reference to the given bool and assigns it to the DnsResolver field.
func (o *LogForwarderEventType) SetDnsResolver(v bool) {
	o.DnsResolver = &v
}

// GetDdos returns the Ddos field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetDdos() bool {
	if o == nil || IsNil(o.Ddos) {
		var ret bool
		return ret
	}
	return *o.Ddos
}

// GetDdosOk returns a tuple with the Ddos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetDdosOk() (*bool, bool) {
	if o == nil || IsNil(o.Ddos) {
		return nil, false
	}
	return o.Ddos, true
}

// HasDdos returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasDdos() bool {
	if o != nil && !IsNil(o.Ddos) {
		return true
	}

	return false
}

// SetDdos gets a reference to the given bool and assigns it to the Ddos field.
func (o *LogForwarderEventType) SetDdos(v bool) {
	o.Ddos = &v
}

// GetRatelimit returns the Ratelimit field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetRatelimit() bool {
	if o == nil || IsNil(o.Ratelimit) {
		var ret bool
		return ret
	}
	return *o.Ratelimit
}

// GetRatelimitOk returns a tuple with the Ratelimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetRatelimitOk() (*bool, bool) {
	if o == nil || IsNil(o.Ratelimit) {
		return nil, false
	}
	return o.Ratelimit, true
}

// HasRatelimit returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasRatelimit() bool {
	if o != nil && !IsNil(o.Ratelimit) {
		return true
	}

	return false
}

// SetRatelimit gets a reference to the given bool and assigns it to the Ratelimit field.
func (o *LogForwarderEventType) SetRatelimit(v bool) {
	o.Ratelimit = &v
}

// GetWaf returns the Waf field value if set, zero value otherwise.
func (o *LogForwarderEventType) GetWaf() bool {
	if o == nil || IsNil(o.Waf) {
		var ret bool
		return ret
	}
	return *o.Waf
}

// GetWafOk returns a tuple with the Waf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderEventType) GetWafOk() (*bool, bool) {
	if o == nil || IsNil(o.Waf) {
		return nil, false
	}
	return o.Waf, true
}

// HasWaf returns a boolean if a field has been set.
func (o *LogForwarderEventType) HasWaf() bool {
	if o != nil && !IsNil(o.Waf) {
		return true
	}

	return false
}

// SetWaf gets a reference to the given bool and assigns it to the Waf field.
func (o *LogForwarderEventType) SetWaf(v bool) {
	o.Waf = &v
}

func (o LogForwarderEventType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderEventType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ServerIp) {
		toSerialize["server_ip"] = o.ServerIp
	}
	if !IsNil(o.ServerPort) {
		toSerialize["server_port"] = o.ServerPort
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.QueryString) {
		toSerialize["query_string"] = o.QueryString
	}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !IsNil(o.DnsResolver) {
		toSerialize["dns_resolver"] = o.DnsResolver
	}
	if !IsNil(o.Ddos) {
		toSerialize["ddos"] = o.Ddos
	}
	if !IsNil(o.Ratelimit) {
		toSerialize["ratelimit"] = o.Ratelimit
	}
	if !IsNil(o.Waf) {
		toSerialize["waf"] = o.Waf
	}
	return toSerialize, nil
}

type NullableLogForwarderEventType struct {
	value *LogForwarderEventType
	isSet bool
}

func (v NullableLogForwarderEventType) Get() *LogForwarderEventType {
	return v.value
}

func (v *NullableLogForwarderEventType) Set(val *LogForwarderEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderEventType(val *LogForwarderEventType) *NullableLogForwarderEventType {
	return &NullableLogForwarderEventType{value: val, isSet: true}
}

func (v NullableLogForwarderEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


