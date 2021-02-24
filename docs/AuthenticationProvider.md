# AuthenticationProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationProvider

`func NewAuthenticationProvider() *AuthenticationProvider`

NewAuthenticationProvider instantiates a new AuthenticationProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderWithDefaults

`func NewAuthenticationProviderWithDefaults() *AuthenticationProvider`

NewAuthenticationProviderWithDefaults instantiates a new AuthenticationProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *AuthenticationProvider) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthenticationProvider) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthenticationProvider) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthenticationProvider) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

