/*
 * Seq API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package seq

import (
	"encoding/json"
)

// InstanceTitle struct for InstanceTitle
type InstanceTitle struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewInstanceTitle instantiates a new InstanceTitle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTitle() *InstanceTitle {
	this := InstanceTitle{}
	return &this
}

// NewInstanceTitleWithDefaults instantiates a new InstanceTitle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTitleWithDefaults() *InstanceTitle {
	this := InstanceTitle{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceTitle) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTitle) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceTitle) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InstanceTitle) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceTitle) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTitle) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceTitle) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceTitle) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InstanceTitle) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTitle) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InstanceTitle) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InstanceTitle) SetValue(v string) {
	o.Value = &v
}

func (o InstanceTitle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceTitle struct {
	value *InstanceTitle
	isSet bool
}

func (v NullableInstanceTitle) Get() *InstanceTitle {
	return v.value
}

func (v *NullableInstanceTitle) Set(val *InstanceTitle) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTitle) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTitle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTitle(val *InstanceTitle) *NullableInstanceTitle {
	return &NullableInstanceTitle{value: val, isSet: true}
}

func (v NullableInstanceTitle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTitle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


