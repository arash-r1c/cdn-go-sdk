/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.99.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the EmailForwardingStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingStats{}

// EmailForwardingStats struct for EmailForwardingStats
type EmailForwardingStats struct {
	Id *string `json:"id,omitempty"`
	DnsActivation *bool `json:"dns_activation,omitempty"`
	RecipientsCount *int32 `json:"recipients_count,omitempty"`
	AliasesCount *int32 `json:"aliases_count,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	IsConfigured *bool `json:"is_configured,omitempty"`
	EmailsForwarded *int32 `json:"emails_forwarded,omitempty"`
	EmailsBlocked *int32 `json:"emails_blocked,omitempty"`
	EmailsReplied *int32 `json:"emails_replied,omitempty"`
}

// NewEmailForwardingStats instantiates a new EmailForwardingStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingStats() *EmailForwardingStats {
	this := EmailForwardingStats{}
	return &this
}

// NewEmailForwardingStatsWithDefaults instantiates a new EmailForwardingStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingStatsWithDefaults() *EmailForwardingStats {
	this := EmailForwardingStats{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailForwardingStats) SetId(v string) {
	o.Id = &v
}

// GetDnsActivation returns the DnsActivation field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetDnsActivation() bool {
	if o == nil || IsNil(o.DnsActivation) {
		var ret bool
		return ret
	}
	return *o.DnsActivation
}

// GetDnsActivationOk returns a tuple with the DnsActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetDnsActivationOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsActivation) {
		return nil, false
	}
	return o.DnsActivation, true
}

// HasDnsActivation returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasDnsActivation() bool {
	if o != nil && !IsNil(o.DnsActivation) {
		return true
	}

	return false
}

// SetDnsActivation gets a reference to the given bool and assigns it to the DnsActivation field.
func (o *EmailForwardingStats) SetDnsActivation(v bool) {
	o.DnsActivation = &v
}

// GetRecipientsCount returns the RecipientsCount field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetRecipientsCount() int32 {
	if o == nil || IsNil(o.RecipientsCount) {
		var ret int32
		return ret
	}
	return *o.RecipientsCount
}

// GetRecipientsCountOk returns a tuple with the RecipientsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetRecipientsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RecipientsCount) {
		return nil, false
	}
	return o.RecipientsCount, true
}

// HasRecipientsCount returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasRecipientsCount() bool {
	if o != nil && !IsNil(o.RecipientsCount) {
		return true
	}

	return false
}

// SetRecipientsCount gets a reference to the given int32 and assigns it to the RecipientsCount field.
func (o *EmailForwardingStats) SetRecipientsCount(v int32) {
	o.RecipientsCount = &v
}

// GetAliasesCount returns the AliasesCount field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetAliasesCount() int32 {
	if o == nil || IsNil(o.AliasesCount) {
		var ret int32
		return ret
	}
	return *o.AliasesCount
}

// GetAliasesCountOk returns a tuple with the AliasesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetAliasesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AliasesCount) {
		return nil, false
	}
	return o.AliasesCount, true
}

// HasAliasesCount returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasAliasesCount() bool {
	if o != nil && !IsNil(o.AliasesCount) {
		return true
	}

	return false
}

// SetAliasesCount gets a reference to the given int32 and assigns it to the AliasesCount field.
func (o *EmailForwardingStats) SetAliasesCount(v int32) {
	o.AliasesCount = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EmailForwardingStats) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsConfigured returns the IsConfigured field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetIsConfigured() bool {
	if o == nil || IsNil(o.IsConfigured) {
		var ret bool
		return ret
	}
	return *o.IsConfigured
}

// GetIsConfiguredOk returns a tuple with the IsConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetIsConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConfigured) {
		return nil, false
	}
	return o.IsConfigured, true
}

// HasIsConfigured returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasIsConfigured() bool {
	if o != nil && !IsNil(o.IsConfigured) {
		return true
	}

	return false
}

// SetIsConfigured gets a reference to the given bool and assigns it to the IsConfigured field.
func (o *EmailForwardingStats) SetIsConfigured(v bool) {
	o.IsConfigured = &v
}

// GetEmailsForwarded returns the EmailsForwarded field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetEmailsForwarded() int32 {
	if o == nil || IsNil(o.EmailsForwarded) {
		var ret int32
		return ret
	}
	return *o.EmailsForwarded
}

// GetEmailsForwardedOk returns a tuple with the EmailsForwarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetEmailsForwardedOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsForwarded) {
		return nil, false
	}
	return o.EmailsForwarded, true
}

// HasEmailsForwarded returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasEmailsForwarded() bool {
	if o != nil && !IsNil(o.EmailsForwarded) {
		return true
	}

	return false
}

// SetEmailsForwarded gets a reference to the given int32 and assigns it to the EmailsForwarded field.
func (o *EmailForwardingStats) SetEmailsForwarded(v int32) {
	o.EmailsForwarded = &v
}

// GetEmailsBlocked returns the EmailsBlocked field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetEmailsBlocked() int32 {
	if o == nil || IsNil(o.EmailsBlocked) {
		var ret int32
		return ret
	}
	return *o.EmailsBlocked
}

// GetEmailsBlockedOk returns a tuple with the EmailsBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetEmailsBlockedOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsBlocked) {
		return nil, false
	}
	return o.EmailsBlocked, true
}

// HasEmailsBlocked returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasEmailsBlocked() bool {
	if o != nil && !IsNil(o.EmailsBlocked) {
		return true
	}

	return false
}

// SetEmailsBlocked gets a reference to the given int32 and assigns it to the EmailsBlocked field.
func (o *EmailForwardingStats) SetEmailsBlocked(v int32) {
	o.EmailsBlocked = &v
}

// GetEmailsReplied returns the EmailsReplied field value if set, zero value otherwise.
func (o *EmailForwardingStats) GetEmailsReplied() int32 {
	if o == nil || IsNil(o.EmailsReplied) {
		var ret int32
		return ret
	}
	return *o.EmailsReplied
}

// GetEmailsRepliedOk returns a tuple with the EmailsReplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStats) GetEmailsRepliedOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsReplied) {
		return nil, false
	}
	return o.EmailsReplied, true
}

// HasEmailsReplied returns a boolean if a field has been set.
func (o *EmailForwardingStats) HasEmailsReplied() bool {
	if o != nil && !IsNil(o.EmailsReplied) {
		return true
	}

	return false
}

// SetEmailsReplied gets a reference to the given int32 and assigns it to the EmailsReplied field.
func (o *EmailForwardingStats) SetEmailsReplied(v int32) {
	o.EmailsReplied = &v
}

func (o EmailForwardingStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DnsActivation) {
		toSerialize["dns_activation"] = o.DnsActivation
	}
	if !IsNil(o.RecipientsCount) {
		toSerialize["recipients_count"] = o.RecipientsCount
	}
	if !IsNil(o.AliasesCount) {
		toSerialize["aliases_count"] = o.AliasesCount
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsConfigured) {
		toSerialize["is_configured"] = o.IsConfigured
	}
	if !IsNil(o.EmailsForwarded) {
		toSerialize["emails_forwarded"] = o.EmailsForwarded
	}
	if !IsNil(o.EmailsBlocked) {
		toSerialize["emails_blocked"] = o.EmailsBlocked
	}
	if !IsNil(o.EmailsReplied) {
		toSerialize["emails_replied"] = o.EmailsReplied
	}
	return toSerialize, nil
}

type NullableEmailForwardingStats struct {
	value *EmailForwardingStats
	isSet bool
}

func (v NullableEmailForwardingStats) Get() *EmailForwardingStats {
	return v.value
}

func (v *NullableEmailForwardingStats) Set(val *EmailForwardingStats) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingStats) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingStats(val *EmailForwardingStats) *NullableEmailForwardingStats {
	return &NullableEmailForwardingStats{value: val, isSet: true}
}

func (v NullableEmailForwardingStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


