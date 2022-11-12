# WriteContractRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Func** | **string** | A contract function name. This can be any valid function from the the ABI of the contract you are interacting with. For example, &#x60;mint&#x60;. | 
**Args** | Pointer to [**[]WriteContractRequestArgsInner**](WriteContractRequestArgsInner.md) | An array of args. This is optional and only necessary if the function being invoked requires arguments per the contract ABI. For example, &#x60;[123, \&quot;Hello\&quot;, false]&#x60;. | [optional] 

## Methods

### NewWriteContractRequest

`func NewWriteContractRequest(func_ string, ) *WriteContractRequest`

NewWriteContractRequest instantiates a new WriteContractRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteContractRequestWithDefaults

`func NewWriteContractRequestWithDefaults() *WriteContractRequest`

NewWriteContractRequestWithDefaults instantiates a new WriteContractRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunc

`func (o *WriteContractRequest) GetFunc() string`

GetFunc returns the Func field if non-nil, zero value otherwise.

### GetFuncOk

`func (o *WriteContractRequest) GetFuncOk() (*string, bool)`

GetFuncOk returns a tuple with the Func field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunc

`func (o *WriteContractRequest) SetFunc(v string)`

SetFunc sets Func field to given value.


### GetArgs

`func (o *WriteContractRequest) GetArgs() []WriteContractRequestArgsInner`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *WriteContractRequest) GetArgsOk() (*[]WriteContractRequestArgsInner, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *WriteContractRequest) SetArgs(v []WriteContractRequestArgsInner)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *WriteContractRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


