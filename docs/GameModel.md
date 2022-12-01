# GameModel

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
**Verified** | Pointer to **bool** | This field has not had a description added. | [optional] 
**UpdatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CreatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 

## Methods

### NewGameModel

`func NewGameModel() *GameModel`

NewGameModel instantiates a new GameModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameModelWithDefaults

`func NewGameModelWithDefaults() *GameModel`

NewGameModelWithDefaults instantiates a new GameModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GameModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWalletId

`func (o *GameModel) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *GameModel) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *GameModel) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *GameModel) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetFundingWalletId

`func (o *GameModel) GetFundingWalletId() string`

GetFundingWalletId returns the FundingWalletId field if non-nil, zero value otherwise.

### GetFundingWalletIdOk

`func (o *GameModel) GetFundingWalletIdOk() (*string, bool)`

GetFundingWalletIdOk returns a tuple with the FundingWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingWalletId

`func (o *GameModel) SetFundingWalletId(v string)`

SetFundingWalletId sets FundingWalletId field to given value.

### HasFundingWalletId

`func (o *GameModel) HasFundingWalletId() bool`

HasFundingWalletId returns a boolean if a field has been set.

### GetEmail

`func (o *GameModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GameModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GameModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GameModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *GameModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GameModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GameModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GameModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRpcs

`func (o *GameModel) GetRpcs() map[string]interface{}`

GetRpcs returns the Rpcs field if non-nil, zero value otherwise.

### GetRpcsOk

`func (o *GameModel) GetRpcsOk() (*map[string]interface{}, bool)`

GetRpcsOk returns a tuple with the Rpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpcs

`func (o *GameModel) SetRpcs(v map[string]interface{})`

SetRpcs sets Rpcs field to given value.

### HasRpcs

`func (o *GameModel) HasRpcs() bool`

HasRpcs returns a boolean if a field has been set.

### GetPublishedKey

`func (o *GameModel) GetPublishedKey() string`

GetPublishedKey returns the PublishedKey field if non-nil, zero value otherwise.

### GetPublishedKeyOk

`func (o *GameModel) GetPublishedKeyOk() (*string, bool)`

GetPublishedKeyOk returns a tuple with the PublishedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedKey

`func (o *GameModel) SetPublishedKey(v string)`

SetPublishedKey sets PublishedKey field to given value.

### HasPublishedKey

`func (o *GameModel) HasPublishedKey() bool`

HasPublishedKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *GameModel) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *GameModel) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *GameModel) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *GameModel) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetVerified

`func (o *GameModel) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *GameModel) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *GameModel) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *GameModel) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GameModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GameModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GameModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GameModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GameModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GameModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


