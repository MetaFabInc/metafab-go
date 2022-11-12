# GetCurrencyFees200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientAddress** | Pointer to **string** | The wallet address that fees from all applicable transactions are automatically sent to. | [optional] 
**BasisPoints** | Pointer to **float32** | The number of fee basis points. 100 basisPoints &#x3D; 1% fee of the total transaction amount deducted from the total received by the recipient. | [optional] 
**FixedAmount** | Pointer to **float32** | The fixed number of currency as a fee regardless of the total transaction amount. 10 &#x3D; 10 of the currency as a fee for any transaction, deducted from the total received by the recipient. | [optional] 
**CapAmount** | Pointer to **float32** | The maximum combined fee between basisPoints and fixedAmount. If the total transaction fee is over this amount, the capAmount will be used as the transaction fee deducted from the total received by the recipient. | [optional] 

## Methods

### NewGetCurrencyFees200Response

`func NewGetCurrencyFees200Response() *GetCurrencyFees200Response`

NewGetCurrencyFees200Response instantiates a new GetCurrencyFees200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrencyFees200ResponseWithDefaults

`func NewGetCurrencyFees200ResponseWithDefaults() *GetCurrencyFees200Response`

NewGetCurrencyFees200ResponseWithDefaults instantiates a new GetCurrencyFees200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientAddress

`func (o *GetCurrencyFees200Response) GetRecipientAddress() string`

GetRecipientAddress returns the RecipientAddress field if non-nil, zero value otherwise.

### GetRecipientAddressOk

`func (o *GetCurrencyFees200Response) GetRecipientAddressOk() (*string, bool)`

GetRecipientAddressOk returns a tuple with the RecipientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddress

`func (o *GetCurrencyFees200Response) SetRecipientAddress(v string)`

SetRecipientAddress sets RecipientAddress field to given value.

### HasRecipientAddress

`func (o *GetCurrencyFees200Response) HasRecipientAddress() bool`

HasRecipientAddress returns a boolean if a field has been set.

### GetBasisPoints

`func (o *GetCurrencyFees200Response) GetBasisPoints() float32`

GetBasisPoints returns the BasisPoints field if non-nil, zero value otherwise.

### GetBasisPointsOk

`func (o *GetCurrencyFees200Response) GetBasisPointsOk() (*float32, bool)`

GetBasisPointsOk returns a tuple with the BasisPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasisPoints

`func (o *GetCurrencyFees200Response) SetBasisPoints(v float32)`

SetBasisPoints sets BasisPoints field to given value.

### HasBasisPoints

`func (o *GetCurrencyFees200Response) HasBasisPoints() bool`

HasBasisPoints returns a boolean if a field has been set.

### GetFixedAmount

`func (o *GetCurrencyFees200Response) GetFixedAmount() float32`

GetFixedAmount returns the FixedAmount field if non-nil, zero value otherwise.

### GetFixedAmountOk

`func (o *GetCurrencyFees200Response) GetFixedAmountOk() (*float32, bool)`

GetFixedAmountOk returns a tuple with the FixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmount

`func (o *GetCurrencyFees200Response) SetFixedAmount(v float32)`

SetFixedAmount sets FixedAmount field to given value.

### HasFixedAmount

`func (o *GetCurrencyFees200Response) HasFixedAmount() bool`

HasFixedAmount returns a boolean if a field has been set.

### GetCapAmount

`func (o *GetCurrencyFees200Response) GetCapAmount() float32`

GetCapAmount returns the CapAmount field if non-nil, zero value otherwise.

### GetCapAmountOk

`func (o *GetCurrencyFees200Response) GetCapAmountOk() (*float32, bool)`

GetCapAmountOk returns a tuple with the CapAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapAmount

`func (o *GetCurrencyFees200Response) SetCapAmount(v float32)`

SetCapAmount sets CapAmount field to given value.

### HasCapAmount

`func (o *GetCurrencyFees200Response) HasCapAmount() bool`

HasCapAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


