# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | **string** |  | 
**NewPassword** | Pointer to **string** |  | [optional] 
**MustChangePassword** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**Preferences** | Pointer to **map[string]interface{}** |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**ViewFilter** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUser

`func NewUser(username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetNewPassword

`func (o *User) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *User) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *User) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *User) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetMustChangePassword

`func (o *User) GetMustChangePassword() bool`

GetMustChangePassword returns the MustChangePassword field if non-nil, zero value otherwise.

### GetMustChangePasswordOk

`func (o *User) GetMustChangePasswordOk() (*bool, bool)`

GetMustChangePasswordOk returns a tuple with the MustChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustChangePassword

`func (o *User) SetMustChangePassword(v bool)`

SetMustChangePassword sets MustChangePassword field to given value.

### HasMustChangePassword

`func (o *User) HasMustChangePassword() bool`

HasMustChangePassword returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *User) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *User) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *User) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *User) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPreferences

`func (o *User) GetPreferences() map[string]interface{}`

GetPreferences returns the Preferences field if non-nil, zero value otherwise.

### GetPreferencesOk

`func (o *User) GetPreferencesOk() (*map[string]interface{}, bool)`

GetPreferencesOk returns a tuple with the Preferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferences

`func (o *User) SetPreferences(v map[string]interface{})`

SetPreferences sets Preferences field to given value.

### HasPreferences

`func (o *User) HasPreferences() bool`

HasPreferences returns a boolean if a field has been set.

### GetRoleIds

`func (o *User) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *User) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *User) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *User) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetViewFilter

`func (o *User) GetViewFilter() map[string]interface{}`

GetViewFilter returns the ViewFilter field if non-nil, zero value otherwise.

### GetViewFilterOk

`func (o *User) GetViewFilterOk() (*map[string]interface{}, bool)`

GetViewFilterOk returns a tuple with the ViewFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewFilter

`func (o *User) SetViewFilter(v map[string]interface{})`

SetViewFilter sets ViewFilter field to given value.

### HasViewFilter

`func (o *User) HasViewFilter() bool`

HasViewFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


