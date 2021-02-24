# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedPermissions** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**InputSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] [readonly] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 
**TokenPrefix** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewApiKey

`func NewApiKey(title string, ) *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedPermissions

`func (o *ApiKey) GetAssignedPermissions() []string`

GetAssignedPermissions returns the AssignedPermissions field if non-nil, zero value otherwise.

### GetAssignedPermissionsOk

`func (o *ApiKey) GetAssignedPermissionsOk() (*[]string, bool)`

GetAssignedPermissionsOk returns a tuple with the AssignedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedPermissions

`func (o *ApiKey) SetAssignedPermissions(v []string)`

SetAssignedPermissions sets AssignedPermissions field to given value.

### HasAssignedPermissions

`func (o *ApiKey) HasAssignedPermissions() bool`

HasAssignedPermissions returns a boolean if a field has been set.

### GetId

`func (o *ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputSettings

`func (o *ApiKey) GetInputSettings() map[string]interface{}`

GetInputSettings returns the InputSettings field if non-nil, zero value otherwise.

### GetInputSettingsOk

`func (o *ApiKey) GetInputSettingsOk() (*map[string]interface{}, bool)`

GetInputSettingsOk returns a tuple with the InputSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSettings

`func (o *ApiKey) SetInputSettings(v map[string]interface{})`

SetInputSettings sets InputSettings field to given value.

### HasInputSettings

`func (o *ApiKey) HasInputSettings() bool`

HasInputSettings returns a boolean if a field has been set.

### GetIsDefault

`func (o *ApiKey) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ApiKey) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ApiKey) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ApiKey) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiKey) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiKey) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiKey) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiKey) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTitle

`func (o *ApiKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiKey) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetToken

`func (o *ApiKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenPrefix

`func (o *ApiKey) GetTokenPrefix() string`

GetTokenPrefix returns the TokenPrefix field if non-nil, zero value otherwise.

### GetTokenPrefixOk

`func (o *ApiKey) GetTokenPrefixOk() (*string, bool)`

GetTokenPrefixOk returns a tuple with the TokenPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPrefix

`func (o *ApiKey) SetTokenPrefix(v string)`

SetTokenPrefix sets TokenPrefix field to given value.

### HasTokenPrefix

`func (o *ApiKey) HasTokenPrefix() bool`

HasTokenPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


