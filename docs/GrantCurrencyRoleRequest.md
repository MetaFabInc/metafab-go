# GrantCurrencyRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | A valid MetaFab role or bytes string representing a role, such as &#x60;minter&#x60; or &#x60;0xc9eb32e43bf5ecbceacf00b32281dfc5d6d700a0db676ea26ccf938a385ac3b7&#x60; | 
**Address** | Pointer to **string** | A valid EVM based address to grant the role to. | [optional] 
**WalletId** | Pointer to **string** | A wallet id within the MetaFab ecosystem to grant the role to. | [optional] 

## Methods

### NewGrantCurrencyRoleRequest

`func NewGrantCurrencyRoleRequest(role string, ) *GrantCurrencyRoleRequest`

NewGrantCurrencyRoleRequest instantiates a new GrantCurrencyRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantCurrencyRoleRequestWithDefaults

`func NewGrantCurrencyRoleRequestWithDefaults() *GrantCurrencyRoleRequest`

NewGrantCurrencyRoleRequestWithDefaults instantiates a new GrantCurrencyRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *GrantCurrencyRoleRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GrantCurrencyRoleRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GrantCurrencyRoleRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetAddress

`func (o *GrantCurrencyRoleRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GrantCurrencyRoleRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GrantCurrencyRoleRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GrantCurrencyRoleRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetWalletId

`func (o *GrantCurrencyRoleRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *GrantCurrencyRoleRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *GrantCurrencyRoleRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *GrantCurrencyRoleRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


