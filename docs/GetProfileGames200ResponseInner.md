# GetProfileGames200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublishedKey** | Pointer to **string** |  | [optional] 
**IconImageUrl** | Pointer to **string** |  | [optional] 
**CoverImageUrl** | Pointer to **string** |  | [optional] 
**PrimaryColorHex** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]PublicPlayer**](PublicPlayer.md) |  | [optional] 

## Methods

### NewGetProfileGames200ResponseInner

`func NewGetProfileGames200ResponseInner() *GetProfileGames200ResponseInner`

NewGetProfileGames200ResponseInner instantiates a new GetProfileGames200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileGames200ResponseInnerWithDefaults

`func NewGetProfileGames200ResponseInnerWithDefaults() *GetProfileGames200ResponseInner`

NewGetProfileGames200ResponseInnerWithDefaults instantiates a new GetProfileGames200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetProfileGames200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProfileGames200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProfileGames200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetProfileGames200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetProfileGames200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetProfileGames200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetProfileGames200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetProfileGames200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishedKey

`func (o *GetProfileGames200ResponseInner) GetPublishedKey() string`

GetPublishedKey returns the PublishedKey field if non-nil, zero value otherwise.

### GetPublishedKeyOk

`func (o *GetProfileGames200ResponseInner) GetPublishedKeyOk() (*string, bool)`

GetPublishedKeyOk returns a tuple with the PublishedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedKey

`func (o *GetProfileGames200ResponseInner) SetPublishedKey(v string)`

SetPublishedKey sets PublishedKey field to given value.

### HasPublishedKey

`func (o *GetProfileGames200ResponseInner) HasPublishedKey() bool`

HasPublishedKey returns a boolean if a field has been set.

### GetIconImageUrl

`func (o *GetProfileGames200ResponseInner) GetIconImageUrl() string`

GetIconImageUrl returns the IconImageUrl field if non-nil, zero value otherwise.

### GetIconImageUrlOk

`func (o *GetProfileGames200ResponseInner) GetIconImageUrlOk() (*string, bool)`

GetIconImageUrlOk returns a tuple with the IconImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconImageUrl

`func (o *GetProfileGames200ResponseInner) SetIconImageUrl(v string)`

SetIconImageUrl sets IconImageUrl field to given value.

### HasIconImageUrl

`func (o *GetProfileGames200ResponseInner) HasIconImageUrl() bool`

HasIconImageUrl returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *GetProfileGames200ResponseInner) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *GetProfileGames200ResponseInner) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *GetProfileGames200ResponseInner) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *GetProfileGames200ResponseInner) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetPrimaryColorHex

`func (o *GetProfileGames200ResponseInner) GetPrimaryColorHex() string`

GetPrimaryColorHex returns the PrimaryColorHex field if non-nil, zero value otherwise.

### GetPrimaryColorHexOk

`func (o *GetProfileGames200ResponseInner) GetPrimaryColorHexOk() (*string, bool)`

GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColorHex

`func (o *GetProfileGames200ResponseInner) SetPrimaryColorHex(v string)`

SetPrimaryColorHex sets PrimaryColorHex field to given value.

### HasPrimaryColorHex

`func (o *GetProfileGames200ResponseInner) HasPrimaryColorHex() bool`

HasPrimaryColorHex returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetProfileGames200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetProfileGames200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetProfileGames200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetProfileGames200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPlayers

`func (o *GetProfileGames200ResponseInner) GetPlayers() []PublicPlayer`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GetProfileGames200ResponseInner) GetPlayersOk() (*[]PublicPlayer, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GetProfileGames200ResponseInner) SetPlayers(v []PublicPlayer)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *GetProfileGames200ResponseInner) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


