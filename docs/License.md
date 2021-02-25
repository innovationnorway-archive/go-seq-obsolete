# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseText** | **string** |  | 
**IsValid** | Pointer to **bool** |  | [optional] [readonly] 
**IsSingleUser** | Pointer to **bool** |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **string** |  | [optional] [readonly] 
**StatusDescription** | Pointer to **string** |  | [optional] [readonly] 
**IsWarning** | Pointer to **bool** |  | [optional] [readonly] 
**CanRenewOnlineNow** | Pointer to **bool** |  | [optional] [readonly] 
**LicensedUsers** | Pointer to **int32** |  | [optional] [readonly] 
**AutomaticallyRefresh** | Pointer to **bool** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLicense

`func NewLicense(licenseText string, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseText

`func (o *License) GetLicenseText() string`

GetLicenseText returns the LicenseText field if non-nil, zero value otherwise.

### GetLicenseTextOk

`func (o *License) GetLicenseTextOk() (*string, bool)`

GetLicenseTextOk returns a tuple with the LicenseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseText

`func (o *License) SetLicenseText(v string)`

SetLicenseText sets LicenseText field to given value.


### GetIsValid

`func (o *License) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *License) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *License) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *License) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetIsSingleUser

`func (o *License) GetIsSingleUser() bool`

GetIsSingleUser returns the IsSingleUser field if non-nil, zero value otherwise.

### GetIsSingleUserOk

`func (o *License) GetIsSingleUserOk() (*bool, bool)`

GetIsSingleUserOk returns a tuple with the IsSingleUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleUser

`func (o *License) SetIsSingleUser(v bool)`

SetIsSingleUser sets IsSingleUser field to given value.

### HasIsSingleUser

`func (o *License) HasIsSingleUser() bool`

HasIsSingleUser returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *License) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *License) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *License) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *License) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetStatusDescription

`func (o *License) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *License) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *License) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *License) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetIsWarning

`func (o *License) GetIsWarning() bool`

GetIsWarning returns the IsWarning field if non-nil, zero value otherwise.

### GetIsWarningOk

`func (o *License) GetIsWarningOk() (*bool, bool)`

GetIsWarningOk returns a tuple with the IsWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWarning

`func (o *License) SetIsWarning(v bool)`

SetIsWarning sets IsWarning field to given value.

### HasIsWarning

`func (o *License) HasIsWarning() bool`

HasIsWarning returns a boolean if a field has been set.

### GetCanRenewOnlineNow

`func (o *License) GetCanRenewOnlineNow() bool`

GetCanRenewOnlineNow returns the CanRenewOnlineNow field if non-nil, zero value otherwise.

### GetCanRenewOnlineNowOk

`func (o *License) GetCanRenewOnlineNowOk() (*bool, bool)`

GetCanRenewOnlineNowOk returns a tuple with the CanRenewOnlineNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRenewOnlineNow

`func (o *License) SetCanRenewOnlineNow(v bool)`

SetCanRenewOnlineNow sets CanRenewOnlineNow field to given value.

### HasCanRenewOnlineNow

`func (o *License) HasCanRenewOnlineNow() bool`

HasCanRenewOnlineNow returns a boolean if a field has been set.

### GetLicensedUsers

`func (o *License) GetLicensedUsers() int32`

GetLicensedUsers returns the LicensedUsers field if non-nil, zero value otherwise.

### GetLicensedUsersOk

`func (o *License) GetLicensedUsersOk() (*int32, bool)`

GetLicensedUsersOk returns a tuple with the LicensedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedUsers

`func (o *License) SetLicensedUsers(v int32)`

SetLicensedUsers sets LicensedUsers field to given value.

### HasLicensedUsers

`func (o *License) HasLicensedUsers() bool`

HasLicensedUsers returns a boolean if a field has been set.

### GetAutomaticallyRefresh

`func (o *License) GetAutomaticallyRefresh() bool`

GetAutomaticallyRefresh returns the AutomaticallyRefresh field if non-nil, zero value otherwise.

### GetAutomaticallyRefreshOk

`func (o *License) GetAutomaticallyRefreshOk() (*bool, bool)`

GetAutomaticallyRefreshOk returns a tuple with the AutomaticallyRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyRefresh

`func (o *License) SetAutomaticallyRefresh(v bool)`

SetAutomaticallyRefresh sets AutomaticallyRefresh field to given value.

### HasAutomaticallyRefresh

`func (o *License) HasAutomaticallyRefresh() bool`

HasAutomaticallyRefresh returns a boolean if a field has been set.

### GetId

`func (o *License) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *License) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *License) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *License) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


