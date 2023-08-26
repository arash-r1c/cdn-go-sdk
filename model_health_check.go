/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.106.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the HealthCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheck{}

// HealthCheck struct for HealthCheck
type HealthCheck struct {
	RequestConfig *HealthCheckRequestConfig `json:"request_config,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// can be IP/Host when type is `upstream`, otherwise it must be a valid record ID
	Origin *string `json:"origin,omitempty"`
	OriginType *string `json:"origin_type,omitempty"`
	Upstreams []string `json:"upstreams,omitempty"`
	// In milliseconds
	Interval *int32 `json:"interval,omitempty"`
	Threshold *int32 `json:"threshold,omitempty"`
	Type *string `json:"type,omitempty"`
	// The health-check is off or on
	Status *bool `json:"status,omitempty"`
	// Number of immediate retries in case of a timeout
	Retries *int32 `json:"retries,omitempty"`
	Zones []HealthCheckZone `json:"zones,omitempty"`
	MonitoringUpdatedAt NullableTime `json:"monitoring_updated_at,omitempty"`
}

// NewHealthCheck instantiates a new HealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheck() *HealthCheck {
	this := HealthCheck{}
	var status bool = true
	this.Status = &status
	return &this
}

// NewHealthCheckWithDefaults instantiates a new HealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckWithDefaults() *HealthCheck {
	this := HealthCheck{}
	var status bool = true
	this.Status = &status
	return &this
}

// GetRequestConfig returns the RequestConfig field value if set, zero value otherwise.
func (o *HealthCheck) GetRequestConfig() HealthCheckRequestConfig {
	if o == nil || IsNil(o.RequestConfig) {
		var ret HealthCheckRequestConfig
		return ret
	}
	return *o.RequestConfig
}

// GetRequestConfigOk returns a tuple with the RequestConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetRequestConfigOk() (*HealthCheckRequestConfig, bool) {
	if o == nil || IsNil(o.RequestConfig) {
		return nil, false
	}
	return o.RequestConfig, true
}

// HasRequestConfig returns a boolean if a field has been set.
func (o *HealthCheck) HasRequestConfig() bool {
	if o != nil && !IsNil(o.RequestConfig) {
		return true
	}

	return false
}

// SetRequestConfig gets a reference to the given HealthCheckRequestConfig and assigns it to the RequestConfig field.
func (o *HealthCheck) SetRequestConfig(v HealthCheckRequestConfig) {
	o.RequestConfig = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HealthCheck) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HealthCheck) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HealthCheck) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HealthCheck) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HealthCheck) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HealthCheck) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HealthCheck) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HealthCheck) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HealthCheck) SetDescription(v string) {
	o.Description = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *HealthCheck) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *HealthCheck) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *HealthCheck) SetOrigin(v string) {
	o.Origin = &v
}

// GetOriginType returns the OriginType field value if set, zero value otherwise.
func (o *HealthCheck) GetOriginType() string {
	if o == nil || IsNil(o.OriginType) {
		var ret string
		return ret
	}
	return *o.OriginType
}

// GetOriginTypeOk returns a tuple with the OriginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetOriginTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OriginType) {
		return nil, false
	}
	return o.OriginType, true
}

// HasOriginType returns a boolean if a field has been set.
func (o *HealthCheck) HasOriginType() bool {
	if o != nil && !IsNil(o.OriginType) {
		return true
	}

	return false
}

// SetOriginType gets a reference to the given string and assigns it to the OriginType field.
func (o *HealthCheck) SetOriginType(v string) {
	o.OriginType = &v
}

// GetUpstreams returns the Upstreams field value if set, zero value otherwise.
func (o *HealthCheck) GetUpstreams() []string {
	if o == nil || IsNil(o.Upstreams) {
		var ret []string
		return ret
	}
	return o.Upstreams
}

// GetUpstreamsOk returns a tuple with the Upstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetUpstreamsOk() ([]string, bool) {
	if o == nil || IsNil(o.Upstreams) {
		return nil, false
	}
	return o.Upstreams, true
}

// HasUpstreams returns a boolean if a field has been set.
func (o *HealthCheck) HasUpstreams() bool {
	if o != nil && !IsNil(o.Upstreams) {
		return true
	}

	return false
}

// SetUpstreams gets a reference to the given []string and assigns it to the Upstreams field.
func (o *HealthCheck) SetUpstreams(v []string) {
	o.Upstreams = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *HealthCheck) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *HealthCheck) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *HealthCheck) SetInterval(v int32) {
	o.Interval = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *HealthCheck) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *HealthCheck) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *HealthCheck) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HealthCheck) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HealthCheck) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HealthCheck) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheck) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheck) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HealthCheck) SetStatus(v bool) {
	o.Status = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *HealthCheck) GetRetries() int32 {
	if o == nil || IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *HealthCheck) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *HealthCheck) SetRetries(v int32) {
	o.Retries = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *HealthCheck) GetZones() []HealthCheckZone {
	if o == nil || IsNil(o.Zones) {
		var ret []HealthCheckZone
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck) GetZonesOk() ([]HealthCheckZone, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *HealthCheck) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []HealthCheckZone and assigns it to the Zones field.
func (o *HealthCheck) SetZones(v []HealthCheckZone) {
	o.Zones = v
}

// GetMonitoringUpdatedAt returns the MonitoringUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HealthCheck) GetMonitoringUpdatedAt() time.Time {
	if o == nil || IsNil(o.MonitoringUpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MonitoringUpdatedAt.Get()
}

// GetMonitoringUpdatedAtOk returns a tuple with the MonitoringUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HealthCheck) GetMonitoringUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitoringUpdatedAt.Get(), o.MonitoringUpdatedAt.IsSet()
}

// HasMonitoringUpdatedAt returns a boolean if a field has been set.
func (o *HealthCheck) HasMonitoringUpdatedAt() bool {
	if o != nil && o.MonitoringUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetMonitoringUpdatedAt gets a reference to the given NullableTime and assigns it to the MonitoringUpdatedAt field.
func (o *HealthCheck) SetMonitoringUpdatedAt(v time.Time) {
	o.MonitoringUpdatedAt.Set(&v)
}
// SetMonitoringUpdatedAtNil sets the value for MonitoringUpdatedAt to be an explicit nil
func (o *HealthCheck) SetMonitoringUpdatedAtNil() {
	o.MonitoringUpdatedAt.Set(nil)
}

// UnsetMonitoringUpdatedAt ensures that no value is present for MonitoringUpdatedAt, not even an explicit nil
func (o *HealthCheck) UnsetMonitoringUpdatedAt() {
	o.MonitoringUpdatedAt.Unset()
}

func (o HealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestConfig) {
		toSerialize["request_config"] = o.RequestConfig
	}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.OriginType) {
		toSerialize["origin_type"] = o.OriginType
	}
	if !IsNil(o.Upstreams) {
		toSerialize["upstreams"] = o.Upstreams
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	if o.MonitoringUpdatedAt.IsSet() {
		toSerialize["monitoring_updated_at"] = o.MonitoringUpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableHealthCheck struct {
	value *HealthCheck
	isSet bool
}

func (v NullableHealthCheck) Get() *HealthCheck {
	return v.value
}

func (v *NullableHealthCheck) Set(val *HealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheck(val *HealthCheck) *NullableHealthCheck {
	return &NullableHealthCheck{value: val, isSet: true}
}

func (v NullableHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


