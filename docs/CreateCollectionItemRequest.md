# CreateCollectionItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | A unique itemId to use for this item within the collection. If an existing itemId is used, the current metadata will be overriden. Any number may be used.  The terms &#x60;itemId&#x60; and &#x60;collectionItemId&#x60; are used interchangeably and equivalent to one another throughout MetaFab documentation. | 
**Name** | **string** | The name of this item. | 
**Description** | **string** | A text description of this item. This is a great spot to include lore, game mechanics and more related to this item. | 
**ImageBase64** | Pointer to **string** | A base64 string of the image for this item. Either &#x60;imageBase64&#x60; or &#x60;imageUrl&#x60; must be provided. Supported image formats are &#x60;jpg&#x60;, &#x60;jpeg&#x60;, &#x60;png&#x60;, &#x60;gif&#x60;. Recommended size of 1:1 aspect ratio and no more than 1500x1500 pixels. | [optional] 
**ImageUrl** | Pointer to **string** | An external url that resolves to an image to use for this item. This can also be set to an ipfs:// uri. Recommended size of 1:1 aspect ratio and no more than 1500x1500 pixels. | [optional] 
**ExternalUrl** | Pointer to **string** | An optional URL where players can go to learn more about this item specifically, or your game, or any other link. | [optional] 
**Attributes** | Pointer to [**[]CreateCollectionItemRequestAttributesInner**](CreateCollectionItemRequestAttributesInner.md) | An array of objects that conform with the [metadata attributes standard that can be found here](https://docs.opensea.io/docs/metadata-standards#attributes) | [optional] 
**Data** | Pointer to **map[string]interface{}** | An arbitrary object of data attached to the top level metadata object. This is useful for including data or resource URLs specific to your game. Such as URLs that point to 3D models, music files, bitmaps, or anything else you need to reference. | [optional] 

## Methods

### NewCreateCollectionItemRequest

`func NewCreateCollectionItemRequest(id float32, name string, description string, ) *CreateCollectionItemRequest`

NewCreateCollectionItemRequest instantiates a new CreateCollectionItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionItemRequestWithDefaults

`func NewCreateCollectionItemRequestWithDefaults() *CreateCollectionItemRequest`

NewCreateCollectionItemRequestWithDefaults instantiates a new CreateCollectionItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCollectionItemRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCollectionItemRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCollectionItemRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *CreateCollectionItemRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCollectionItemRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCollectionItemRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCollectionItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageBase64

`func (o *CreateCollectionItemRequest) GetImageBase64() string`

GetImageBase64 returns the ImageBase64 field if non-nil, zero value otherwise.

### GetImageBase64Ok

`func (o *CreateCollectionItemRequest) GetImageBase64Ok() (*string, bool)`

GetImageBase64Ok returns a tuple with the ImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBase64

`func (o *CreateCollectionItemRequest) SetImageBase64(v string)`

SetImageBase64 sets ImageBase64 field to given value.

### HasImageBase64

`func (o *CreateCollectionItemRequest) HasImageBase64() bool`

HasImageBase64 returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateCollectionItemRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateCollectionItemRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateCollectionItemRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateCollectionItemRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExternalUrl

`func (o *CreateCollectionItemRequest) GetExternalUrl() string`

GetExternalUrl returns the ExternalUrl field if non-nil, zero value otherwise.

### GetExternalUrlOk

`func (o *CreateCollectionItemRequest) GetExternalUrlOk() (*string, bool)`

GetExternalUrlOk returns a tuple with the ExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrl

`func (o *CreateCollectionItemRequest) SetExternalUrl(v string)`

SetExternalUrl sets ExternalUrl field to given value.

### HasExternalUrl

`func (o *CreateCollectionItemRequest) HasExternalUrl() bool`

HasExternalUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *CreateCollectionItemRequest) GetAttributes() []CreateCollectionItemRequestAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateCollectionItemRequest) GetAttributesOk() (*[]CreateCollectionItemRequestAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateCollectionItemRequest) SetAttributes(v []CreateCollectionItemRequestAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateCollectionItemRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetData

`func (o *CreateCollectionItemRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateCollectionItemRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateCollectionItemRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CreateCollectionItemRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


