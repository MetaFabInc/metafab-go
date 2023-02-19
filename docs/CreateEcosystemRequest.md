# CreateEcosystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the ecosystem you&#39;re creating. | 
**Email** | **string** | The email address associated with this ecosystem and used to login/authenticate as the ecosystem. | 
**Password** | **string** | The password to authenticate as the ecosystem. | 

## Methods

### NewCreateEcosystemRequest

`func NewCreateEcosystemRequest(name string, email string, password string, ) *CreateEcosystemRequest`

NewCreateEcosystemRequest instantiates a new CreateEcosystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEcosystemRequestWithDefaults

`func NewCreateEcosystemRequestWithDefaults() *CreateEcosystemRequest`

NewCreateEcosystemRequestWithDefaults instantiates a new CreateEcosystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEcosystemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEcosystemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEcosystemRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateEcosystemRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateEcosystemRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateEcosystemRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CreateEcosystemRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateEcosystemRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateEcosystemRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


