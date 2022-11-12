# SetCurrencyFeesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientAddress** | **string** | The recipient address of currency transaction fees. | 
**BasisPoints** | **float32** | A percentage fee for every transaction represented in basis points. To set a 1.5% fee, you would use a value of 150. This value can be 0, denoting no percentage fees. | 
**FixedAmount** | **float32** | A fixed fee for every transaction. A value of 0.5 would mean 0.5 of the currency of a transaction is always taken as a fee. This value can be 0, denoting no fixed fees. | 
**CapAmount** | **float32** | The maximum fee amount for any single transaction. The total fee of a transaction is calculated as the sum of the basis points (percentage) fee, and fixed fee. If a calculated fee is greater than this maximum fee value, the maximum fee will be used instead. | 

## Methods

### NewSetCurrencyFeesRequest

`func NewSetCurrencyFeesRequest(recipientAddress string, basisPoints float32, fixedAmount float32, capAmount float32, ) *SetCurrencyFeesRequest`

NewSetCurrencyFeesRequest instantiates a new SetCurrencyFeesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCurrencyFeesRequestWithDefaults

`func NewSetCurrencyFeesRequestWithDefaults() *SetCurrencyFeesRequest`

NewSetCurrencyFeesRequestWithDefaults instantiates a new SetCurrencyFeesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientAddress

`func (o *SetCurrencyFeesRequest) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *SetCurrencyFeesRequest) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *SetCurrencyFeesRequest) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.


### GetBasisPoints

`func (o *SetCurrencyFeesRequest) GetBasisPoints() float32`

GetBasisPoints returns the BasisPoints field if non-nil, zero value otherwise.

### GetBasisPointsOk

`func (o *SetCurrencyFeesRequest) GetBasisPointsOk() (*float32, bool)`

GetBasisPointsOk returns a tuple with the BasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisPoints

`func (o *SetCurrencyFeesRequest) SetBasisPoints(v float32)`

SetBasisPoints sets BasisPoints field to given value.


### GetFixedAmount

`func (o *SetCurrencyFeesRequest) GetFixedAmount() float32`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *SetCurrencyFeesRequest) GetFixedAmountOk() (*float32, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *SetCurrencyFeesRequest) SetFixedAmount(v float32)`

SetFixedAmount sets FixedAmount field to given value.


### GetCapAmount

`func (o *SetCurrencyFeesRequest) GetCapAmount() float32`

GetCapAmount returns the CapAmount field if non-nil, zero value otherwise.

### GetCapAmountOk

`func (o *SetCurrencyFeesRequest) GetCapAmountOk() (*float32, bool)`

GetCapAmountOk returns a tuple with the CapAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapAmount

`func (o *SetCurrencyFeesRequest) SetCapAmount(v float32)`

SetCapAmount sets CapAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


