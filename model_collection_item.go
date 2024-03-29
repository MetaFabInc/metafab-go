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

// CollectionItem struct for CollectionItem
type CollectionItem struct {
	// This field has not had a description added.
	Id *string `json:"id,omitempty"`
	// This field has not had a description added.
	Image *string `json:"image,omitempty"`
	// This field has not had a description added.
	Name *string `json:"name,omitempty"`
	// This field has not had a description added.
	Description *string `json:"description,omitempty"`
	// This field has not had a description added.
	ExternalUrl *string `json:"externalUrl,omitempty"`
	// This field has not had a description added.
	Attributes []CollectionItemAttributesInner `json:"attributes,omitempty"`
	// This field has not had a description added.
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewCollectionItem instantiates a new CollectionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionItem() *CollectionItem {
	this := CollectionItem{}
	return &this
}

// NewCollectionItemWithDefaults instantiates a new CollectionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionItemWithDefaults() *CollectionItem {
	this := CollectionItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CollectionItem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CollectionItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CollectionItem) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CollectionItem) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CollectionItem) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CollectionItem) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CollectionItem) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CollectionItem) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CollectionItem) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CollectionItem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CollectionItem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CollectionItem) SetDescription(v string) {
	o.Description = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *CollectionItem) GetExternalUrl() string {
	if o == nil || isNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetExternalUrlOk() (*string, bool) {
	if o == nil || isNil(o.ExternalUrl) {
    return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *CollectionItem) HasExternalUrl() bool {
	if o != nil && !isNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *CollectionItem) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CollectionItem) GetAttributes() []CollectionItemAttributesInner {
	if o == nil || isNil(o.Attributes) {
		var ret []CollectionItemAttributesInner
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetAttributesOk() ([]CollectionItemAttributesInner, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CollectionItem) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []CollectionItemAttributesInner and assigns it to the Attributes field.
func (o *CollectionItem) SetAttributes(v []CollectionItemAttributesInner) {
	o.Attributes = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CollectionItem) GetData() map[string]interface{} {
	if o == nil || isNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionItem) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Data) {
    return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CollectionItem) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CollectionItem) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o CollectionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
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

type NullableCollectionItem struct {
	value *CollectionItem
	isSet bool
}

func (v NullableCollectionItem) Get() *CollectionItem {
	return v.value
}

func (v *NullableCollectionItem) Set(val *CollectionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionItem(val *CollectionItem) *NullableCollectionItem {
	return &NullableCollectionItem{value: val, isSet: true}
}

func (v NullableCollectionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


