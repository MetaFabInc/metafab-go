# AuthPlayer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This field has not had a description added. | [optional] 
**GameId** | Pointer to **string** | This field has not had a description added. | [optional] 
**WalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**ConnectedWalletId** | Pointer to **string** | This field has not had a description added. | [optional] 
**ProfileId** | Pointer to **string** | This field has not had a description added. | [optional] 
**Username** | Pointer to **string** | This field has not had a description added. | [optional] 
**AccessToken** | Pointer to **string** | This field has not had a description added. | [optional] 
**ProfilePermissions** | Pointer to **map[string]interface{}** | This field has not had a description added. | [optional] 
**UpdatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CreatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**WalletDecryptKey** | Pointer to **string** | This field has not had a description added. | [optional] 
**Wallet** | Pointer to [**WalletModel**](WalletModel.md) |  | [optional] 
**CustodialWallet** | Pointer to [**WalletModel**](WalletModel.md) |  | [optional] 

## Methods

### NewAuthPlayer200Response

`func NewAuthPlayer200Response() *AuthPlayer200Response`

NewAuthPlayer200Response instantiates a new AuthPlayer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthPlayer200ResponseWithDefaults

`func NewAuthPlayer200ResponseWithDefaults() *AuthPlayer200Response`

NewAuthPlayer200ResponseWithDefaults instantiates a new AuthPlayer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthPlayer200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthPlayer200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthPlayer200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthPlayer200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGameId

`func (o *AuthPlayer200Response) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *AuthPlayer200Response) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *AuthPlayer200Response) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *AuthPlayer200Response) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetWalletId

`func (o *AuthPlayer200Response) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *AuthPlayer200Response) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *AuthPlayer200Response) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *AuthPlayer200Response) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetConnectedWalletId

`func (o *AuthPlayer200Response) GetConnectedWalletId() string`

GetConnectedWalletId returns the ConnectedWalletId field if non-nil, zero value otherwise.

### GetConnectedWalletIdOk

`func (o *AuthPlayer200Response) GetConnectedWalletIdOk() (*string, bool)`

GetConnectedWalletIdOk returns a tuple with the ConnectedWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedWalletId

`func (o *AuthPlayer200Response) SetConnectedWalletId(v string)`

SetConnectedWalletId sets ConnectedWalletId field to given value.

### HasConnectedWalletId

`func (o *AuthPlayer200Response) HasConnectedWalletId() bool`

HasConnectedWalletId returns a boolean if a field has been set.

### GetProfileId

`func (o *AuthPlayer200Response) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AuthPlayer200Response) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AuthPlayer200Response) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AuthPlayer200Response) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetUsername

`func (o *AuthPlayer200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthPlayer200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthPlayer200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuthPlayer200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthPlayer200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthPlayer200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthPlayer200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthPlayer200Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetProfilePermissions

`func (o *AuthPlayer200Response) GetProfilePermissions() map[string]interface{}`

GetProfilePermissions returns the ProfilePermissions field if non-nil, zero value otherwise.

### GetProfilePermissionsOk

`func (o *AuthPlayer200Response) GetProfilePermissionsOk() (*map[string]interface{}, bool)`

GetProfilePermissionsOk returns a tuple with the ProfilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePermissions

`func (o *AuthPlayer200Response) SetProfilePermissions(v map[string]interface{})`

SetProfilePermissions sets ProfilePermissions field to given value.

### HasProfilePermissions

`func (o *AuthPlayer200Response) HasProfilePermissions() bool`

HasProfilePermissions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthPlayer200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthPlayer200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthPlayer200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthPlayer200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthPlayer200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthPlayer200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthPlayer200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthPlayer200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetWalletDecryptKey

`func (o *AuthPlayer200Response) GetWalletDecryptKey() string`

GetWalletDecryptKey returns the WalletDecryptKey field if non-nil, zero value otherwise.

### GetWalletDecryptKeyOk

`func (o *AuthPlayer200Response) GetWalletDecryptKeyOk() (*string, bool)`

GetWalletDecryptKeyOk returns a tuple with the WalletDecryptKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletDecryptKey

`func (o *AuthPlayer200Response) SetWalletDecryptKey(v string)`

SetWalletDecryptKey sets WalletDecryptKey field to given value.

### HasWalletDecryptKey

`func (o *AuthPlayer200Response) HasWalletDecryptKey() bool`

HasWalletDecryptKey returns a boolean if a field has been set.

### GetWallet

`func (o *AuthPlayer200Response) GetWallet() WalletModel`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *AuthPlayer200Response) GetWalletOk() (*WalletModel, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *AuthPlayer200Response) SetWallet(v WalletModel)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *AuthPlayer200Response) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetCustodialWallet

`func (o *AuthPlayer200Response) GetCustodialWallet() WalletModel`

GetCustodialWallet returns the CustodialWallet field if non-nil, zero value otherwise.

### GetCustodialWalletOk

`func (o *AuthPlayer200Response) GetCustodialWalletOk() (*WalletModel, bool)`

GetCustodialWalletOk returns a tuple with the CustodialWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodialWallet

`func (o *AuthPlayer200Response) SetCustodialWallet(v WalletModel)`

SetCustodialWallet sets CustodialWallet field to given value.

### HasCustodialWallet

`func (o *AuthPlayer200Response) HasCustodialWallet() bool`

HasCustodialWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


