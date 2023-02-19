# UpdateProfilePlayerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**map[string]ProfilePermissionsValue**](ProfilePermissionsValue.md) | A properly formatted permissions object that validates against the MetaFab profile permissions schema. | [optional] 

## Methods

### NewUpdateProfilePlayerRequest

`func NewUpdateProfilePlayerRequest() *UpdateProfilePlayerRequest`

NewUpdateProfilePlayerRequest instantiates a new UpdateProfilePlayerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfilePlayerRequestWithDefaults

`func NewUpdateProfilePlayerRequestWithDefaults() *UpdateProfilePlayerRequest`

NewUpdateProfilePlayerRequestWithDefaults instantiates a new UpdateProfilePlayerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *UpdateProfilePlayerRequest) GetPermissions() map[string]ProfilePermissionsValue`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateProfilePlayerRequest) GetPermissionsOk() (*map[string]ProfilePermissionsValue, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateProfilePlayerRequest) SetPermissions(v map[string]ProfilePermissionsValue)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateProfilePlayerRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


