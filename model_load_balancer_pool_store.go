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

// checks if the LoadBalancerPoolStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPoolStore{}

// LoadBalancerPoolStore struct for LoadBalancerPoolStore
type LoadBalancerPoolStore struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status bool `json:"status"`
	// Zero means the default pool
	Priority *int32 `json:"priority,omitempty"`
	Method string `json:"method"`
	Keepalive string `json:"keepalive"`
	NextUpstreamTcp string `json:"next_upstream_tcp"`
	Regions []string `json:"regions,omitempty"`
	Origins []LoadBalancerOriginStore `json:"origins,omitempty"`
}

// NewLoadBalancerPoolStore instantiates a new LoadBalancerPoolStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPoolStore(name string, status bool, method string, keepalive string, nextUpstreamTcp string) *LoadBalancerPoolStore {
	this := LoadBalancerPoolStore{}
	this.Name = name
	this.Status = status
	this.Method = method
	this.Keepalive = keepalive
	this.NextUpstreamTcp = nextUpstreamTcp
	return &this
}

// NewLoadBalancerPoolStoreWithDefaults instantiates a new LoadBalancerPoolStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPoolStoreWithDefaults() *LoadBalancerPoolStore {
	this := LoadBalancerPoolStore{}
	var keepalive string = "off"
	this.Keepalive = keepalive
	var nextUpstreamTcp string = "off"
	this.NextUpstreamTcp = nextUpstreamTcp
	return &this
}

// GetName returns the Name field value
func (o *LoadBalancerPoolStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancerPoolStore) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerPoolStore) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerPoolStore) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerPoolStore) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *LoadBalancerPoolStore) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LoadBalancerPoolStore) SetStatus(v bool) {
	o.Status = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *LoadBalancerPoolStore) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *LoadBalancerPoolStore) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *LoadBalancerPoolStore) SetPriority(v int32) {
	o.Priority = &v
}

// GetMethod returns the Method field value
func (o *LoadBalancerPoolStore) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *LoadBalancerPoolStore) SetMethod(v string) {
	o.Method = v
}

// GetKeepalive returns the Keepalive field value
func (o *LoadBalancerPoolStore) GetKeepalive() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Keepalive
}

// GetKeepaliveOk returns a tuple with the Keepalive field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetKeepaliveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keepalive, true
}

// SetKeepalive sets field value
func (o *LoadBalancerPoolStore) SetKeepalive(v string) {
	o.Keepalive = v
}

// GetNextUpstreamTcp returns the NextUpstreamTcp field value
func (o *LoadBalancerPoolStore) GetNextUpstreamTcp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextUpstreamTcp
}

// GetNextUpstreamTcpOk returns a tuple with the NextUpstreamTcp field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetNextUpstreamTcpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextUpstreamTcp, true
}

// SetNextUpstreamTcp sets field value
func (o *LoadBalancerPoolStore) SetNextUpstreamTcp(v string) {
	o.NextUpstreamTcp = v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *LoadBalancerPoolStore) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *LoadBalancerPoolStore) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *LoadBalancerPoolStore) SetRegions(v []string) {
	o.Regions = v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *LoadBalancerPoolStore) GetOrigins() []LoadBalancerOriginStore {
	if o == nil || IsNil(o.Origins) {
		var ret []LoadBalancerOriginStore
		return ret
	}
	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolStore) GetOriginsOk() ([]LoadBalancerOriginStore, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *LoadBalancerPoolStore) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given []LoadBalancerOriginStore and assigns it to the Origins field.
func (o *LoadBalancerPoolStore) SetOrigins(v []LoadBalancerOriginStore) {
	o.Origins = v
}

func (o LoadBalancerPoolStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPoolStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["method"] = o.Method
	toSerialize["keepalive"] = o.Keepalive
	toSerialize["next_upstream_tcp"] = o.NextUpstreamTcp
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.Origins) {
		toSerialize["origins"] = o.Origins
	}
	return toSerialize, nil
}

type NullableLoadBalancerPoolStore struct {
	value *LoadBalancerPoolStore
	isSet bool
}

func (v NullableLoadBalancerPoolStore) Get() *LoadBalancerPoolStore {
	return v.value
}

func (v *NullableLoadBalancerPoolStore) Set(val *LoadBalancerPoolStore) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPoolStore) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPoolStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPoolStore(val *LoadBalancerPoolStore) *NullableLoadBalancerPoolStore {
	return &NullableLoadBalancerPoolStore{value: val, isSet: true}
}

func (v NullableLoadBalancerPoolStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPoolStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


