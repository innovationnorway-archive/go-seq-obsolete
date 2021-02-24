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

// License struct for License
type License struct {
	LicenceText *string `json:"LicenceText,omitempty"`
	IsValid *bool `json:"IsValid,omitempty"`
	IsSingleUser *bool `json:"IsSingleUser,omitempty"`
	SubscriptionId *string `json:"SubscriptionId,omitempty"`
	StatusDescription *string `json:"StatusDescription,omitempty"`
	IsWarning *bool `json:"IsWarning,omitempty"`
	CanRenewOnlineNow *bool `json:"CanRenewOnlineNow,omitempty"`
	LicensedUsers *int32 `json:"LicensedUsers,omitempty"`
	AutomaticallyRefresh *bool `json:"AutomaticallyRefresh,omitempty"`
	Id *string `json:"Id,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetLicenceText returns the LicenceText field value if set, zero value otherwise.
func (o *License) GetLicenceText() string {
	if o == nil || o.LicenceText == nil {
		var ret string
		return ret
	}
	return *o.LicenceText
}

// GetLicenceTextOk returns a tuple with the LicenceText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicenceTextOk() (*string, bool) {
	if o == nil || o.LicenceText == nil {
		return nil, false
	}
	return o.LicenceText, true
}

// HasLicenceText returns a boolean if a field has been set.
func (o *License) HasLicenceText() bool {
	if o != nil && o.LicenceText != nil {
		return true
	}

	return false
}

// SetLicenceText gets a reference to the given string and assigns it to the LicenceText field.
func (o *License) SetLicenceText(v string) {
	o.LicenceText = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *License) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *License) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *License) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetIsSingleUser returns the IsSingleUser field value if set, zero value otherwise.
func (o *License) GetIsSingleUser() bool {
	if o == nil || o.IsSingleUser == nil {
		var ret bool
		return ret
	}
	return *o.IsSingleUser
}

// GetIsSingleUserOk returns a tuple with the IsSingleUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIsSingleUserOk() (*bool, bool) {
	if o == nil || o.IsSingleUser == nil {
		return nil, false
	}
	return o.IsSingleUser, true
}

// HasIsSingleUser returns a boolean if a field has been set.
func (o *License) HasIsSingleUser() bool {
	if o != nil && o.IsSingleUser != nil {
		return true
	}

	return false
}

// SetIsSingleUser gets a reference to the given bool and assigns it to the IsSingleUser field.
func (o *License) SetIsSingleUser(v bool) {
	o.IsSingleUser = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *License) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *License) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *License) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *License) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *License) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *License) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetIsWarning returns the IsWarning field value if set, zero value otherwise.
func (o *License) GetIsWarning() bool {
	if o == nil || o.IsWarning == nil {
		var ret bool
		return ret
	}
	return *o.IsWarning
}

// GetIsWarningOk returns a tuple with the IsWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIsWarningOk() (*bool, bool) {
	if o == nil || o.IsWarning == nil {
		return nil, false
	}
	return o.IsWarning, true
}

// HasIsWarning returns a boolean if a field has been set.
func (o *License) HasIsWarning() bool {
	if o != nil && o.IsWarning != nil {
		return true
	}

	return false
}

// SetIsWarning gets a reference to the given bool and assigns it to the IsWarning field.
func (o *License) SetIsWarning(v bool) {
	o.IsWarning = &v
}

// GetCanRenewOnlineNow returns the CanRenewOnlineNow field value if set, zero value otherwise.
func (o *License) GetCanRenewOnlineNow() bool {
	if o == nil || o.CanRenewOnlineNow == nil {
		var ret bool
		return ret
	}
	return *o.CanRenewOnlineNow
}

// GetCanRenewOnlineNowOk returns a tuple with the CanRenewOnlineNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetCanRenewOnlineNowOk() (*bool, bool) {
	if o == nil || o.CanRenewOnlineNow == nil {
		return nil, false
	}
	return o.CanRenewOnlineNow, true
}

// HasCanRenewOnlineNow returns a boolean if a field has been set.
func (o *License) HasCanRenewOnlineNow() bool {
	if o != nil && o.CanRenewOnlineNow != nil {
		return true
	}

	return false
}

// SetCanRenewOnlineNow gets a reference to the given bool and assigns it to the CanRenewOnlineNow field.
func (o *License) SetCanRenewOnlineNow(v bool) {
	o.CanRenewOnlineNow = &v
}

// GetLicensedUsers returns the LicensedUsers field value if set, zero value otherwise.
func (o *License) GetLicensedUsers() int32 {
	if o == nil || o.LicensedUsers == nil {
		var ret int32
		return ret
	}
	return *o.LicensedUsers
}

// GetLicensedUsersOk returns a tuple with the LicensedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicensedUsersOk() (*int32, bool) {
	if o == nil || o.LicensedUsers == nil {
		return nil, false
	}
	return o.LicensedUsers, true
}

// HasLicensedUsers returns a boolean if a field has been set.
func (o *License) HasLicensedUsers() bool {
	if o != nil && o.LicensedUsers != nil {
		return true
	}

	return false
}

// SetLicensedUsers gets a reference to the given int32 and assigns it to the LicensedUsers field.
func (o *License) SetLicensedUsers(v int32) {
	o.LicensedUsers = &v
}

// GetAutomaticallyRefresh returns the AutomaticallyRefresh field value if set, zero value otherwise.
func (o *License) GetAutomaticallyRefresh() bool {
	if o == nil || o.AutomaticallyRefresh == nil {
		var ret bool
		return ret
	}
	return *o.AutomaticallyRefresh
}

// GetAutomaticallyRefreshOk returns a tuple with the AutomaticallyRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetAutomaticallyRefreshOk() (*bool, bool) {
	if o == nil || o.AutomaticallyRefresh == nil {
		return nil, false
	}
	return o.AutomaticallyRefresh, true
}

// HasAutomaticallyRefresh returns a boolean if a field has been set.
func (o *License) HasAutomaticallyRefresh() bool {
	if o != nil && o.AutomaticallyRefresh != nil {
		return true
	}

	return false
}

// SetAutomaticallyRefresh gets a reference to the given bool and assigns it to the AutomaticallyRefresh field.
func (o *License) SetAutomaticallyRefresh(v bool) {
	o.AutomaticallyRefresh = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *License) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *License) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *License) SetId(v string) {
	o.Id = &v
}

func (o License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LicenceText != nil {
		toSerialize["LicenceText"] = o.LicenceText
	}
	if o.IsValid != nil {
		toSerialize["IsValid"] = o.IsValid
	}
	if o.IsSingleUser != nil {
		toSerialize["IsSingleUser"] = o.IsSingleUser
	}
	if o.SubscriptionId != nil {
		toSerialize["SubscriptionId"] = o.SubscriptionId
	}
	if o.StatusDescription != nil {
		toSerialize["StatusDescription"] = o.StatusDescription
	}
	if o.IsWarning != nil {
		toSerialize["IsWarning"] = o.IsWarning
	}
	if o.CanRenewOnlineNow != nil {
		toSerialize["CanRenewOnlineNow"] = o.CanRenewOnlineNow
	}
	if o.LicensedUsers != nil {
		toSerialize["LicensedUsers"] = o.LicensedUsers
	}
	if o.AutomaticallyRefresh != nil {
		toSerialize["AutomaticallyRefresh"] = o.AutomaticallyRefresh
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


