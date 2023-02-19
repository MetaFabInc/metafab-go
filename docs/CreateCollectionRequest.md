# CreateCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this item collection. This can be anything, such as &#x60;Production - Item Collection&#x60;, &#x60;Testing - My Game Item Collection&#x60;, etc. | [optional] 
**Chain** | **string** | The blockchain you want to deploy this item collection on. Support for new blockchains are added over time. | 

## Methods

### NewCreateCollectionRequest

`func NewCreateCollectionRequest(chain string, ) *CreateCollectionRequest`

NewCreateCollectionRequest instantiates a new CreateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionRequestWithDefaults

`func NewCreateCollectionRequestWithDefaults() *CreateCollectionRequest`

NewCreateCollectionRequestWithDefaults instantiates a new CreateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCollectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCollectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCollectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCollectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChain

`func (o *CreateCollectionRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateCollectionRequest) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateCollectionRequest) SetChain(v string)`

SetChain sets Chain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


