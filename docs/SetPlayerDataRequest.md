# SetPlayerDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectedData** | Pointer to **map[string]interface{}** | protectedData can only be set if &#x60;X-Authorization&#x60; includes credentials for the game the target player is a part of. Expects an arbitrary object allowed to contain any set of properties and nested data within those properties, including arrays.  protectedData is great for storing sensitive player data like tracking experience points, off-chain inventories, save states, and more - things that players shouldn&#39;t have the ability to directly change themselves. | [optional] 
**PublicData** | Pointer to **map[string]interface{}** | publicData can be set if &#x60;X-Authorization&#x60; includes credentials for the target player or game the player is a part of. Expects an arbitrary object allowed to contain any set of properties and nested data within those properties, including arrays.  publicData is great for storing player preferences like in-game settings, non-sensitive data and more. Anything that a player should have the ability to directly change themselves without client or server verification can be stored in publicData. | [optional] 

## Methods

### NewSetPlayerDataRequest

`func NewSetPlayerDataRequest() *SetPlayerDataRequest`

NewSetPlayerDataRequest instantiates a new SetPlayerDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPlayerDataRequestWithDefaults

`func NewSetPlayerDataRequestWithDefaults() *SetPlayerDataRequest`

NewSetPlayerDataRequestWithDefaults instantiates a new SetPlayerDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectedData

`func (o *SetPlayerDataRequest) GetProtectedData() map[string]interface{}`

GetProtectedData returns the ProtectedData field if non-nil, zero value otherwise.

### GetProtectedDataOk

`func (o *SetPlayerDataRequest) GetProtectedDataOk() (*map[string]interface{}, bool)`

GetProtectedDataOk returns a tuple with the ProtectedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedData

`func (o *SetPlayerDataRequest) SetProtectedData(v map[string]interface{})`

SetProtectedData sets ProtectedData field to given value.

### HasProtectedData

`func (o *SetPlayerDataRequest) HasProtectedData() bool`

HasProtectedData returns a boolean if a field has been set.

### GetPublicData

`func (o *SetPlayerDataRequest) GetPublicData() map[string]interface{}`

GetPublicData returns the PublicData field if non-nil, zero value otherwise.

### GetPublicDataOk

`func (o *SetPlayerDataRequest) GetPublicDataOk() (*map[string]interface{}, bool)`

GetPublicDataOk returns a tuple with the PublicData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicData

`func (o *SetPlayerDataRequest) SetPublicData(v map[string]interface{})`

SetPublicData sets PublicData field to given value.

### HasPublicData

`func (o *SetPlayerDataRequest) HasPublicData() bool`

HasPublicData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


