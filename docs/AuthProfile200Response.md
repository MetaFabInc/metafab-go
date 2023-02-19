# AuthProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This field has not had a description added. | [optional] 
**EcosystemId** | Pointer to **string** | This field has not had a description added. | [optional] 
**WalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**ConnectedWalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**Username** | Pointer to **string** | This field has not had a description added. | [optional] 
**AccessToken** | Pointer to **string** | This field has not had a description added. | [optional] 
**UpdatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CreatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**WalletDecryptKey** | Pointer to **string** | This field has not had a description added. | [optional] 
**Wallet** | Pointer to [**WalletModel**](WalletModel.md) |  | [optional] 
**CustodialWallet** | Pointer to [**WalletModel**](WalletModel.md) |  | [optional] 

## Methods

### NewAuthProfile200Response

`func NewAuthProfile200Response() *AuthProfile200Response`

NewAuthProfile200Response instantiates a new AuthProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthProfile200ResponseWithDefaults

`func NewAuthProfile200ResponseWithDefaults() *AuthProfile200Response`

NewAuthProfile200ResponseWithDefaults instantiates a new AuthProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthProfile200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthProfile200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthProfile200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthProfile200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEcosystemId

`func (o *AuthProfile200Response) GetEcosystemId() string`

GetEcosystemId returns the EcosystemId field if non-nil, zero value otherwise.

### GetEcosystemIdOk

`func (o *AuthProfile200Response) GetEcosystemIdOk() (*string, bool)`

GetEcosystemIdOk returns a tuple with the EcosystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystemId

`func (o *AuthProfile200Response) SetEcosystemId(v string)`

SetEcosystemId sets EcosystemId field to given value.

### HasEcosystemId

`func (o *AuthProfile200Response) HasEcosystemId() bool`

HasEcosystemId returns a boolean if a field has been set.

### GetWalletId

`func (o *AuthProfile200Response) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AuthProfile200Response) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AuthProfile200Response) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *AuthProfile200Response) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetConnectedWalletId

`func (o *AuthProfile200Response) GetConnectedWalletId() string`

GetConnectedWalletId returns the ConnectedWalletId field if non-nil, zero value otherwise.

### GetConnectedWalletIdOk

`func (o *AuthProfile200Response) GetConnectedWalletIdOk() (*string, bool)`

GetConnectedWalletIdOk returns a tuple with the ConnectedWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedWalletId

`func (o *AuthProfile200Response) SetConnectedWalletId(v string)`

SetConnectedWalletId sets ConnectedWalletId field to given value.

### HasConnectedWalletId

`func (o *AuthProfile200Response) HasConnectedWalletId() bool`

HasConnectedWalletId returns a boolean if a field has been set.

### GetUsername

`func (o *AuthProfile200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthProfile200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthProfile200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuthProfile200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthProfile200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthProfile200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthProfile200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthProfile200Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthProfile200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthProfile200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthProfile200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthProfile200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthProfile200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthProfile200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthProfile200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthProfile200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetWalletDecryptKey

`func (o *AuthProfile200Response) GetWalletDecryptKey() string`

GetWalletDecryptKey returns the WalletDecryptKey field if non-nil, zero value otherwise.

### GetWalletDecryptKeyOk

`func (o *AuthProfile200Response) GetWalletDecryptKeyOk() (*string, bool)`

GetWalletDecryptKeyOk returns a tuple with the WalletDecryptKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletDecryptKey

`func (o *AuthProfile200Response) SetWalletDecryptKey(v string)`

SetWalletDecryptKey sets WalletDecryptKey field to given value.

### HasWalletDecryptKey

`func (o *AuthProfile200Response) HasWalletDecryptKey() bool`

HasWalletDecryptKey returns a boolean if a field has been set.

### GetWallet

`func (o *AuthProfile200Response) GetWallet() WalletModel`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *AuthProfile200Response) GetWalletOk() (*WalletModel, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *AuthProfile200Response) SetWallet(v WalletModel)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *AuthProfile200Response) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetCustodialWallet

`func (o *AuthProfile200Response) GetCustodialWallet() WalletModel`

GetCustodialWallet returns the CustodialWallet field if non-nil, zero value otherwise.

### GetCustodialWalletOk

`func (o *AuthProfile200Response) GetCustodialWalletOk() (*WalletModel, bool)`

GetCustodialWalletOk returns a tuple with the CustodialWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodialWallet

`func (o *AuthProfile200Response) SetCustodialWallet(v WalletModel)`

SetCustodialWallet sets CustodialWallet field to given value.

### HasCustodialWallet

`func (o *AuthProfile200Response) HasCustodialWallet() bool`

HasCustodialWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


