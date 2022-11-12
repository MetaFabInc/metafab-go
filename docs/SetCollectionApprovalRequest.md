# SetCollectionApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approved** | **bool** | A true or false value approves or disapproves the provided address or address associated with the provided walletId. | 
**Address** | Pointer to **string** | A valid EVM based address to allow control over the authenticating game or player&#39;s wallet items for this collection. | [optional] 
**WalletId** | Pointer to **[]string** | A wallet id within the MetaFab ecosystem to allow control over the authenticating game or player&#39;s wallet items for this collection. | [optional] 

## Methods

### NewSetCollectionApprovalRequest

`func NewSetCollectionApprovalRequest(approved bool, ) *SetCollectionApprovalRequest`

NewSetCollectionApprovalRequest instantiates a new SetCollectionApprovalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCollectionApprovalRequestWithDefaults

`func NewSetCollectionApprovalRequestWithDefaults() *SetCollectionApprovalRequest`

NewSetCollectionApprovalRequestWithDefaults instantiates a new SetCollectionApprovalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproved

`func (o *SetCollectionApprovalRequest) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *SetCollectionApprovalRequest) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *SetCollectionApprovalRequest) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetAddress

`func (o *SetCollectionApprovalRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SetCollectionApprovalRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SetCollectionApprovalRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SetCollectionApprovalRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetWalletId

`func (o *SetCollectionApprovalRequest) GetWalletId() []string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SetCollectionApprovalRequest) GetWalletIdOk() (*[]string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SetCollectionApprovalRequest) SetWalletId(v []string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *SetCollectionApprovalRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


