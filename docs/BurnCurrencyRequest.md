# BurnCurrencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | The amount of currency to remove (burn). The currency balance of the authenticating game or player&#39;s wallet must be equal to or greater than this amount. | 

## Methods

### NewBurnCurrencyRequest

`func NewBurnCurrencyRequest(amount float32, ) *BurnCurrencyRequest`

NewBurnCurrencyRequest instantiates a new BurnCurrencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurnCurrencyRequestWithDefaults

`func NewBurnCurrencyRequestWithDefaults() *BurnCurrencyRequest`

NewBurnCurrencyRequestWithDefaults instantiates a new BurnCurrencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BurnCurrencyRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BurnCurrencyRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BurnCurrencyRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


