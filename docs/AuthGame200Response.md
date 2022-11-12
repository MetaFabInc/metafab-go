# AuthGame200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This field has not had a description added. | [optional] 
**WalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**FundingWalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**Email** | Pointer to **string** | This field has not had a description added. | [optional] 
**Name** | Pointer to **string** | This field has not had a description added. | [optional] 
**Rpcs** | Pointer to **map[string]interface{}** | This field has not had a description added. | [optional] 
**PublishedKey** | Pointer to **string** | This field has not had a description added. | [optional] 
**SecretKey** | Pointer to **string** | This field has not had a description added. | [optional] 
**UpdatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CreatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**Wallet** | Pointer to [**WalletModel**](WalletModel.md) |  | [optional] 
**FundingWallet** | Pointer to [**WalletModel**](WalletModel.md) |  | [optional] 

## Methods

### NewAuthGame200Response

`func NewAuthGame200Response() *AuthGame200Response`

NewAuthGame200Response instantiates a new AuthGame200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGame200ResponseWithDefaults

`func NewAuthGame200ResponseWithDefaults() *AuthGame200Response`

NewAuthGame200ResponseWithDefaults instantiates a new AuthGame200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthGame200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthGame200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthGame200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthGame200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWalletId

`func (o *AuthGame200Response) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AuthGame200Response) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AuthGame200Response) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *AuthGame200Response) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetFundingWalletId

`func (o *AuthGame200Response) GetFundingWalletId() string`

GetFundingWalletId returns the FundingWalletId field if non-nil, zero value otherwise.

### GetFundingWalletIdOk

`func (o *AuthGame200Response) GetFundingWalletIdOk() (*string, bool)`

GetFundingWalletIdOk returns a tuple with the FundingWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingWalletId

`func (o *AuthGame200Response) SetFundingWalletId(v string)`

SetFundingWalletId sets FundingWalletId field to given value.

### HasFundingWalletId

`func (o *AuthGame200Response) HasFundingWalletId() bool`

HasFundingWalletId returns a boolean if a field has been set.

### GetEmail

`func (o *AuthGame200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthGame200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthGame200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthGame200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *AuthGame200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthGame200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthGame200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthGame200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRpcs

`func (o *AuthGame200Response) GetRpcs() map[string]interface{}`

GetRpcs returns the Rpcs field if non-nil, zero value otherwise.

### GetRpcsOk

`func (o *AuthGame200Response) GetRpcsOk() (*map[string]interface{}, bool)`

GetRpcsOk returns a tuple with the Rpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcs

`func (o *AuthGame200Response) SetRpcs(v map[string]interface{})`

SetRpcs sets Rpcs field to given value.

### HasRpcs

`func (o *AuthGame200Response) HasRpcs() bool`

HasRpcs returns a boolean if a field has been set.

### GetPublishedKey

`func (o *AuthGame200Response) GetPublishedKey() string`

GetPublishedKey returns the PublishedKey field if non-nil, zero value otherwise.

### GetPublishedKeyOk

`func (o *AuthGame200Response) GetPublishedKeyOk() (*string, bool)`

GetPublishedKeyOk returns a tuple with the PublishedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedKey

`func (o *AuthGame200Response) SetPublishedKey(v string)`

SetPublishedKey sets PublishedKey field to given value.

### HasPublishedKey

`func (o *AuthGame200Response) HasPublishedKey() bool`

HasPublishedKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AuthGame200Response) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AuthGame200Response) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AuthGame200Response) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AuthGame200Response) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthGame200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthGame200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthGame200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthGame200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthGame200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthGame200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthGame200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthGame200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetWallet

`func (o *AuthGame200Response) GetWallet() WalletModel`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *AuthGame200Response) GetWalletOk() (*WalletModel, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *AuthGame200Response) SetWallet(v WalletModel)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *AuthGame200Response) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetFundingWallet

`func (o *AuthGame200Response) GetFundingWallet() WalletModel`

GetFundingWallet returns the FundingWallet field if non-nil, zero value otherwise.

### GetFundingWalletOk

`func (o *AuthGame200Response) GetFundingWalletOk() (*WalletModel, bool)`

GetFundingWalletOk returns a tuple with the FundingWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingWallet

`func (o *AuthGame200Response) SetFundingWallet(v WalletModel)`

SetFundingWallet sets FundingWallet field to given value.

### HasFundingWallet

`func (o *AuthGame200Response) HasFundingWallet() bool`

HasFundingWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


