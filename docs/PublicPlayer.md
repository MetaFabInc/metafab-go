# PublicPlayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This field has not had a description added. | [optional] 
**GameId** | Pointer to **string** | This field has not had a description added. | [optional] 
**WalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**ConnectedWalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**ProfileId** | Pointer to **string** | This field has not had a description added. | [optional] 
**Username** | Pointer to **string** | This field has not had a description added. | [optional] 
**UpdatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CreatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CustodialWallet** | Pointer to [**PublicPlayerCustodialWallet**](PublicPlayerCustodialWallet.md) |  | [optional] 
**Wallet** | Pointer to [**PublicPlayerCustodialWallet**](PublicPlayerCustodialWallet.md) |  | [optional] 

## Methods

### NewPublicPlayer

`func NewPublicPlayer() *PublicPlayer`

NewPublicPlayer instantiates a new PublicPlayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPlayerWithDefaults

`func NewPublicPlayerWithDefaults() *PublicPlayer`

NewPublicPlayerWithDefaults instantiates a new PublicPlayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicPlayer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicPlayer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicPlayer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicPlayer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGameId

`func (o *PublicPlayer) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *PublicPlayer) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *PublicPlayer) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *PublicPlayer) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetWalletId

`func (o *PublicPlayer) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *PublicPlayer) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *PublicPlayer) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *PublicPlayer) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetConnectedWalletId

`func (o *PublicPlayer) GetConnectedWalletId() string`

GetConnectedWalletId returns the ConnectedWalletId field if non-nil, zero value otherwise.

### GetConnectedWalletIdOk

`func (o *PublicPlayer) GetConnectedWalletIdOk() (*string, bool)`

GetConnectedWalletIdOk returns a tuple with the ConnectedWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedWalletId

`func (o *PublicPlayer) SetConnectedWalletId(v string)`

SetConnectedWalletId sets ConnectedWalletId field to given value.

### HasConnectedWalletId

`func (o *PublicPlayer) HasConnectedWalletId() bool`

HasConnectedWalletId returns a boolean if a field has been set.

### GetProfileId

`func (o *PublicPlayer) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PublicPlayer) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PublicPlayer) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *PublicPlayer) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetUsername

`func (o *PublicPlayer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PublicPlayer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PublicPlayer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PublicPlayer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicPlayer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicPlayer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicPlayer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicPlayer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicPlayer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicPlayer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicPlayer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicPlayer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustodialWallet

`func (o *PublicPlayer) GetCustodialWallet() PublicPlayerCustodialWallet`

GetCustodialWallet returns the CustodialWallet field if non-nil, zero value otherwise.

### GetCustodialWalletOk

`func (o *PublicPlayer) GetCustodialWalletOk() (*PublicPlayerCustodialWallet, bool)`

GetCustodialWalletOk returns a tuple with the CustodialWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodialWallet

`func (o *PublicPlayer) SetCustodialWallet(v PublicPlayerCustodialWallet)`

SetCustodialWallet sets CustodialWallet field to given value.

### HasCustodialWallet

`func (o *PublicPlayer) HasCustodialWallet() bool`

HasCustodialWallet returns a boolean if a field has been set.

### GetWallet

`func (o *PublicPlayer) GetWallet() PublicPlayerCustodialWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *PublicPlayer) GetWalletOk() (*PublicPlayerCustodialWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *PublicPlayer) SetWallet(v PublicPlayerCustodialWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *PublicPlayer) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


