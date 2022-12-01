# SetCollectionItemTimelockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timelock** | **int32** | A unix timestamp (in seconds) defining when the set timelock expires. | 

## Methods

### NewSetCollectionItemTimelockRequest

`func NewSetCollectionItemTimelockRequest(timelock int32, ) *SetCollectionItemTimelockRequest`

NewSetCollectionItemTimelockRequest instantiates a new SetCollectionItemTimelockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCollectionItemTimelockRequestWithDefaults

`func NewSetCollectionItemTimelockRequestWithDefaults() *SetCollectionItemTimelockRequest`

NewSetCollectionItemTimelockRequestWithDefaults instantiates a new SetCollectionItemTimelockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimelock

`func (o *SetCollectionItemTimelockRequest) GetTimelock() int32`

GetTimelock returns the Timelock field if non-nil, zero value otherwise.

### GetTimelockOk

`func (o *SetCollectionItemTimelockRequest) GetTimelockOk() (*int32, bool)`

GetTimelockOk returns a tuple with the Timelock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelock

`func (o *SetCollectionItemTimelockRequest) SetTimelock(v int32)`

SetTimelock sets Timelock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


