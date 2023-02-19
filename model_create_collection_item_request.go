/*
MetaFab API

Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.5.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// CreateCollectionItemRequest struct for CreateCollectionItemRequest
type CreateCollectionItemRequest struct {
	// A unique itemId to use for this item within the collection. If an existing itemId is used, the current metadata will be overriden. Any number may be used.  The terms `itemId` and `collectionItemId` are used interchangeably and equivalent to one another throughout MetaFab documentation.
	Id int32 `json:"id"`
	// The name of this item.
	Name string `json:"name"`
	// A text description of this item. This is a great spot to include lore, game mechanics and more related to this item.
	Description string `json:"description"`
	// A base64 string of the image for this item. Either `imageBase64` or `imageUrl` must be provided. Supported image formats are `jpg`, `jpeg`, `png`, `gif`. Recommended size of 1:1 aspect ratio and no more than 1500x1500 pixels.
	ImageBase64 *string `json:"imageBase64,omitempty"`
	// An external url that resolves to an image to use for this item. This can also be set to an ipfs:// uri. Recommended size of 1:1 aspect ratio and no more than 1500x1500 pixels.
	ImageUrl *string `json:"imageUrl,omitempty"`
	// An optional URL where players can go to learn more about this item specifically, or your game, or any other link.
	ExternalUrl *string `json:"externalUrl,omitempty"`
	// An array of objects that conform with the [metadata attributes standard that can be found here](https://docs.opensea.io/docs/metadata-standards#attributes)
	Attributes []CreateCollectionItemRequestAttributesInner `json:"attributes,omitempty"`
	// An arbitrary object of data attached to the top level metadata object. This is useful for including data or resource URLs specific to your game. Such as URLs that point to 3D models, music files, bitmaps, or anything else you need to reference.
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewCreateCollectionItemRequest instantiates a new CreateCollectionItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollectionItemRequest(id int32, name string, description string) *CreateCollectionItemRequest {
	this := CreateCollectionItemRequest{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewCreateCollectionItemRequestWithDefaults instantiates a new CreateCollectionItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollectionItemRequestWithDefaults() *CreateCollectionItemRequest {
	this := CreateCollectionItemRequest{}
	return &this
}

// GetId returns the Id field value
func (o *CreateCollectionItemRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCollectionItemRequest) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CreateCollectionItemRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCollectionItemRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateCollectionItemRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateCollectionItemRequest) SetDescription(v string) {
	o.Description = v
}

// GetImageBase64 returns the ImageBase64 field value if set, zero value otherwise.
func (o *CreateCollectionItemRequest) GetImageBase64() string {
	if o == nil || isNil(o.ImageBase64) {
		var ret string
		return ret
	}
	return *o.ImageBase64
}

// GetImageBase64Ok returns a tuple with the ImageBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetImageBase64Ok() (*string, bool) {
	if o == nil || isNil(o.ImageBase64) {
    return nil, false
	}
	return o.ImageBase64, true
}

// HasImageBase64 returns a boolean if a field has been set.
func (o *CreateCollectionItemRequest) HasImageBase64() bool {
	if o != nil && !isNil(o.ImageBase64) {
		return true
	}

	return false
}

// SetImageBase64 gets a reference to the given string and assigns it to the ImageBase64 field.
func (o *CreateCollectionItemRequest) SetImageBase64(v string) {
	o.ImageBase64 = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *CreateCollectionItemRequest) GetImageUrl() string {
	if o == nil || isNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrl) {
    return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CreateCollectionItemRequest) HasImageUrl() bool {
	if o != nil && !isNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *CreateCollectionItemRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *CreateCollectionItemRequest) GetExternalUrl() string {
	if o == nil || isNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetExternalUrlOk() (*string, bool) {
	if o == nil || isNil(o.ExternalUrl) {
    return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *CreateCollectionItemRequest) HasExternalUrl() bool {
	if o != nil && !isNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *CreateCollectionItemRequest) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CreateCollectionItemRequest) GetAttributes() []CreateCollectionItemRequestAttributesInner {
	if o == nil || isNil(o.Attributes) {
		var ret []CreateCollectionItemRequestAttributesInner
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetAttributesOk() ([]CreateCollectionItemRequestAttributesInner, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CreateCollectionItemRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []CreateCollectionItemRequestAttributesInner and assigns it to the Attributes field.
func (o *CreateCollectionItemRequest) SetAttributes(v []CreateCollectionItemRequestAttributesInner) {
	o.Attributes = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateCollectionItemRequest) GetData() map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionItemRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
    return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateCollectionItemRequest) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CreateCollectionItemRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o CreateCollectionItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ImageBase64) {
		toSerialize["imageBase64"] = o.ImageBase64
	}
	if !isNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !isNil(o.ExternalUrl) {
		toSerialize["externalUrl"] = o.ExternalUrl
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollectionItemRequest struct {
	value *CreateCollectionItemRequest
	isSet bool
}

func (v NullableCreateCollectionItemRequest) Get() *CreateCollectionItemRequest {
	return v.value
}

func (v *NullableCreateCollectionItemRequest) Set(val *CreateCollectionItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollectionItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollectionItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollectionItemRequest(val *CreateCollectionItemRequest) *NullableCreateCollectionItemRequest {
	return &NullableCreateCollectionItemRequest{value: val, isSet: true}
}

func (v NullableCreateCollectionItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollectionItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


