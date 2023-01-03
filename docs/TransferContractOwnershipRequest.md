# TransferContractOwnershipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerAddress** | **string** | A wallet address to assign as the new owner and administrator of the target smart contract. | 

## Methods

### NewTransferContractOwnershipRequest

`func NewTransferContractOwnershipRequest(ownerAddress string, ) *TransferContractOwnershipRequest`

NewTransferContractOwnershipRequest instantiates a new TransferContractOwnershipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferContractOwnershipRequestWithDefaults

`func NewTransferContractOwnershipRequestWithDefaults() *TransferContractOwnershipRequest`

NewTransferContractOwnershipRequestWithDefaults instantiates a new TransferContractOwnershipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerAddress

`func (o *TransferContractOwnershipRequest) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *TransferContractOwnershipRequest) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *TransferContractOwnershipRequest) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


