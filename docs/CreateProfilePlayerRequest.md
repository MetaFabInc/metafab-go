# CreateProfilePlayerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username to assign to the created player. | 
**Permissions** | Pointer to [**map[string]ProfilePermissionsValue**](ProfilePermissionsValue.md) | A properly formatted permissions object that validates against the MetaFab profile permissions schema. | [optional] 

## Methods

### NewCreateProfilePlayerRequest

`func NewCreateProfilePlayerRequest(username string, ) *CreateProfilePlayerRequest`

NewCreateProfilePlayerRequest instantiates a new CreateProfilePlayerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfilePlayerRequestWithDefaults

`func NewCreateProfilePlayerRequestWithDefaults() *CreateProfilePlayerRequest`

NewCreateProfilePlayerRequestWithDefaults instantiates a new CreateProfilePlayerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateProfilePlayerRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateProfilePlayerRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateProfilePlayerRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPermissions

`func (o *CreateProfilePlayerRequest) GetPermissions() map[string]ProfilePermissionsValue`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateProfilePlayerRequest) GetPermissionsOk() (*map[string]ProfilePermissionsValue, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateProfilePlayerRequest) SetPermissions(v map[string]ProfilePermissionsValue)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateProfilePlayerRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


