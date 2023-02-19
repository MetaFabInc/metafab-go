# ProfilePermissionsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | The target chain for the contract and related permissions. | [optional] 
**Scopes** | Pointer to **[]string** | An optional array of valid permissioning scopes. | [optional] 
**Functions** | Pointer to **[]string** | An optional array of contract functions to request permission for. | [optional] 
**Erc20Limit** | Pointer to **int32** | A maximum lifetime limit of erc20 that can be tranferred for this contract address. | [optional] 
**Erc1155Limits** | Pointer to **map[string]int32** | An object mapping erc1155 ids to maximum lifetime transfer limits of each permitted item id supplied for this contract address. | [optional] 

## Methods

### NewProfilePermissionsValue

`func NewProfilePermissionsValue() *ProfilePermissionsValue`

NewProfilePermissionsValue instantiates a new ProfilePermissionsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilePermissionsValueWithDefaults

`func NewProfilePermissionsValueWithDefaults() *ProfilePermissionsValue`

NewProfilePermissionsValueWithDefaults instantiates a new ProfilePermissionsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ProfilePermissionsValue) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ProfilePermissionsValue) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ProfilePermissionsValue) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *ProfilePermissionsValue) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetScopes

`func (o *ProfilePermissionsValue) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProfilePermissionsValue) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProfilePermissionsValue) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProfilePermissionsValue) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetFunctions

`func (o *ProfilePermissionsValue) GetFunctions() []string`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ProfilePermissionsValue) GetFunctionsOk() (*[]string, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ProfilePermissionsValue) SetFunctions(v []string)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *ProfilePermissionsValue) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### GetErc20Limit

`func (o *ProfilePermissionsValue) GetErc20Limit() int32`

GetErc20Limit returns the Erc20Limit field if non-nil, zero value otherwise.

### GetErc20LimitOk

`func (o *ProfilePermissionsValue) GetErc20LimitOk() (*int32, bool)`

GetErc20LimitOk returns a tuple with the Erc20Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErc20Limit

`func (o *ProfilePermissionsValue) SetErc20Limit(v int32)`

SetErc20Limit sets Erc20Limit field to given value.

### HasErc20Limit

`func (o *ProfilePermissionsValue) HasErc20Limit() bool`

HasErc20Limit returns a boolean if a field has been set.

### GetErc1155Limits

`func (o *ProfilePermissionsValue) GetErc1155Limits() map[string]int32`

GetErc1155Limits returns the Erc1155Limits field if non-nil, zero value otherwise.

### GetErc1155LimitsOk

`func (o *ProfilePermissionsValue) GetErc1155LimitsOk() (*map[string]int32, bool)`

GetErc1155LimitsOk returns a tuple with the Erc1155Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErc1155Limits

`func (o *ProfilePermissionsValue) SetErc1155Limits(v map[string]int32)`

SetErc1155Limits sets Erc1155Limits field to given value.

### HasErc1155Limits

`func (o *ProfilePermissionsValue) HasErc1155Limits() bool`

HasErc1155Limits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


