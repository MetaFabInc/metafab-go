# UpgradeContractTrustedForwarderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwarderAddress** | **string** | A ERC2771 forwarder smart contract address to assign as the new trusted forwarder of the target smart contract. | 

## Methods

### NewUpgradeContractTrustedForwarderRequest

`func NewUpgradeContractTrustedForwarderRequest(forwarderAddress string, ) *UpgradeContractTrustedForwarderRequest`

NewUpgradeContractTrustedForwarderRequest instantiates a new UpgradeContractTrustedForwarderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeContractTrustedForwarderRequestWithDefaults

`func NewUpgradeContractTrustedForwarderRequestWithDefaults() *UpgradeContractTrustedForwarderRequest`

NewUpgradeContractTrustedForwarderRequestWithDefaults instantiates a new UpgradeContractTrustedForwarderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwarderAddress

`func (o *UpgradeContractTrustedForwarderRequest) GetForwarderAddress() string`

GetForwarderAddress returns the ForwarderAddress field if non-nil, zero value otherwise.

### GetForwarderAddressOk

`func (o *UpgradeContractTrustedForwarderRequest) GetForwarderAddressOk() (*string, bool)`

GetForwarderAddressOk returns a tuple with the ForwarderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarderAddress

`func (o *UpgradeContractTrustedForwarderRequest) SetForwarderAddress(v string)`

SetForwarderAddress sets ForwarderAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


