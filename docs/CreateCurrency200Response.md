# CreateCurrency200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This field has not had a description added. | [optional] 
**GameId** | Pointer to **string** | This field has not had a description added. | [optional] 
**ContractId** | Pointer to **string** | This field has not had a description added. | [optional] 
**Name** | Pointer to **string** | This field has not had a description added. | [optional] 
**Symbol** | Pointer to **string** | This field has not had a description added. | [optional] 
**SupplyCap** | Pointer to **int32** | This field has not had a description added. | [optional] 
**UpdatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**CreatedAt** | Pointer to **string** | This field has not had a description added. | [optional] 
**Contract** | Pointer to [**CreateCollection200ResponseAllOfContract**](CreateCollection200ResponseAllOfContract.md) |  | [optional] 

## Methods

### NewCreateCurrency200Response

`func NewCreateCurrency200Response() *CreateCurrency200Response`

NewCreateCurrency200Response instantiates a new CreateCurrency200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCurrency200ResponseWithDefaults

`func NewCreateCurrency200ResponseWithDefaults() *CreateCurrency200Response`

NewCreateCurrency200ResponseWithDefaults instantiates a new CreateCurrency200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCurrency200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCurrency200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCurrency200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCurrency200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGameId

`func (o *CreateCurrency200Response) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *CreateCurrency200Response) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *CreateCurrency200Response) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *CreateCurrency200Response) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetContractId

`func (o *CreateCurrency200Response) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *CreateCurrency200Response) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *CreateCurrency200Response) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *CreateCurrency200Response) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetName

`func (o *CreateCurrency200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCurrency200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCurrency200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCurrency200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *CreateCurrency200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CreateCurrency200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CreateCurrency200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CreateCurrency200Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSupplyCap

`func (o *CreateCurrency200Response) GetSupplyCap() int32`

GetSupplyCap returns the SupplyCap field if non-nil, zero value otherwise.

### GetSupplyCapOk

`func (o *CreateCurrency200Response) GetSupplyCapOk() (*int32, bool)`

GetSupplyCapOk returns a tuple with the SupplyCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyCap

`func (o *CreateCurrency200Response) SetSupplyCap(v int32)`

SetSupplyCap sets SupplyCap field to given value.

### HasSupplyCap

`func (o *CreateCurrency200Response) HasSupplyCap() bool`

HasSupplyCap returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateCurrency200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateCurrency200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateCurrency200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateCurrency200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateCurrency200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCurrency200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCurrency200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateCurrency200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContract

`func (o *CreateCurrency200Response) GetContract() CreateCollection200ResponseAllOfContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *CreateCurrency200Response) GetContractOk() (*CreateCollection200ResponseAllOfContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *CreateCurrency200Response) SetContract(v CreateCollection200ResponseAllOfContract)`

SetContract sets Contract field to given value.

### HasContract

`func (o *CreateCurrency200Response) HasContract() bool`

HasContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


