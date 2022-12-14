/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.4.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// SetLootboxManagerLootboxRequest struct for SetLootboxManagerLootboxRequest
type SetLootboxManagerLootboxRequest struct {
	// A unique lootbox id to use for this lootbox for the lootbox manager. If an existing lootbox id is used, the current lootbox will be updated but the existing number of opens will be kept. If you want to reset the number of opens for an existing lootbox, first remove it using the remove lootbox endpoint, then set it.
	Id int32 `json:"id"`
	// A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for input items required by this lootbox. `inputCollectionAddress` or `inputCollectionId` can optionally be provided.
	InputCollectionAddress *string `json:"inputCollectionAddress,omitempty"`
	// A valid MetaFab collection id that represents the collection for input items required by this lootbox. `inputCollectionAddress` or `inputCollectionId` can optionally be provided.
	InputCollectionId *string `json:"inputCollectionId,omitempty"`
	// An array of item ids from the provided input collection that are required to open this lootbox. Input items are burn upon opening a lootbox.
	InputCollectionItemIds []int32 `json:"inputCollectionItemIds,omitempty"`
	// An array of amounts for each item id from the provided input collection that are required to open this lootbox. Item amounts array indices are reflective of the amount required for a given item id at the same index.
	InputCollectionItemAmounts []int32 `json:"inputCollectionItemAmounts,omitempty"`
	// A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for possible output items given by this lootbox. `outputCollectionAddress` or `outputCollectionId` can optionally be provided.
	OutputCollectionAddress *string `json:"outputCollectionAddress,omitempty"`
	// A valid MetaFab collection id that represents the collection for possible output items given by this lootbox. `outputCollectionAddress` or `outputCollectionId` can optionally be provided.
	OutputCollectionId *string `json:"outputCollectionId,omitempty"`
	// An array of item ids from the provided output collection that are possibly given by this lootbox. Randomly selected output items are automatically minted if the lootbox manager contract has the `minter` role for the output collection contract. Otherwise, they are transferred from the item balance held by the lootbox manager contract.
	OutputCollectionItemIds []int32 `json:"outputCollectionItemIds,omitempty"`
	// An array of amounts for each item id that can be randomly selected from the provided output collection that are given by this lootbox. Item amounts array indices are reflective of the amount required for a given item id at the same index.
	OutputCollectionItemAmounts []int32 `json:"outputCollectionItemAmounts,omitempty"`
	// An array of weights for each item id that can be randomly selected from the provided output collection that are given by this lootbox. Any positive integer for an item's weight can be provided. The weight for an item relative to the sum of all possible item weights determines the probability that an item will be picked upon a lootbox being opened. Item weights array indices are reflective of the probability weight for a given item id at the same index.
	OutputCollectionItemWeights []int32 `json:"outputCollectionItemWeights,omitempty"`
	// The number of items randomly selected from the possible output items when this lootbox is open. If you provide a value greater than 1, it is possible for the same item to be selected more than once, giving the opener more than one of that item's output from the lootbox.
	OutputTotalItems *int32 `json:"outputTotalItems,omitempty"`
}

// NewSetLootboxManagerLootboxRequest instantiates a new SetLootboxManagerLootboxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetLootboxManagerLootboxRequest(id int32) *SetLootboxManagerLootboxRequest {
	this := SetLootboxManagerLootboxRequest{}
	this.Id = id
	return &this
}

// NewSetLootboxManagerLootboxRequestWithDefaults instantiates a new SetLootboxManagerLootboxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetLootboxManagerLootboxRequestWithDefaults() *SetLootboxManagerLootboxRequest {
	this := SetLootboxManagerLootboxRequest{}
	return &this
}

// GetId returns the Id field value
func (o *SetLootboxManagerLootboxRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SetLootboxManagerLootboxRequest) SetId(v int32) {
	o.Id = v
}

// GetInputCollectionAddress returns the InputCollectionAddress field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionAddress() string {
	if o == nil || o.InputCollectionAddress == nil {
		var ret string
		return ret
	}
	return *o.InputCollectionAddress
}

// GetInputCollectionAddressOk returns a tuple with the InputCollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionAddressOk() (*string, bool) {
	if o == nil || o.InputCollectionAddress == nil {
		return nil, false
	}
	return o.InputCollectionAddress, true
}

// HasInputCollectionAddress returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasInputCollectionAddress() bool {
	if o != nil && o.InputCollectionAddress != nil {
		return true
	}

	return false
}

// SetInputCollectionAddress gets a reference to the given string and assigns it to the InputCollectionAddress field.
func (o *SetLootboxManagerLootboxRequest) SetInputCollectionAddress(v string) {
	o.InputCollectionAddress = &v
}

// GetInputCollectionId returns the InputCollectionId field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionId() string {
	if o == nil || o.InputCollectionId == nil {
		var ret string
		return ret
	}
	return *o.InputCollectionId
}

// GetInputCollectionIdOk returns a tuple with the InputCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionIdOk() (*string, bool) {
	if o == nil || o.InputCollectionId == nil {
		return nil, false
	}
	return o.InputCollectionId, true
}

// HasInputCollectionId returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasInputCollectionId() bool {
	if o != nil && o.InputCollectionId != nil {
		return true
	}

	return false
}

// SetInputCollectionId gets a reference to the given string and assigns it to the InputCollectionId field.
func (o *SetLootboxManagerLootboxRequest) SetInputCollectionId(v string) {
	o.InputCollectionId = &v
}

// GetInputCollectionItemIds returns the InputCollectionItemIds field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemIds() []int32 {
	if o == nil || o.InputCollectionItemIds == nil {
		var ret []int32
		return ret
	}
	return o.InputCollectionItemIds
}

// GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemIdsOk() ([]int32, bool) {
	if o == nil || o.InputCollectionItemIds == nil {
		return nil, false
	}
	return o.InputCollectionItemIds, true
}

// HasInputCollectionItemIds returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasInputCollectionItemIds() bool {
	if o != nil && o.InputCollectionItemIds != nil {
		return true
	}

	return false
}

// SetInputCollectionItemIds gets a reference to the given []int32 and assigns it to the InputCollectionItemIds field.
func (o *SetLootboxManagerLootboxRequest) SetInputCollectionItemIds(v []int32) {
	o.InputCollectionItemIds = v
}

// GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemAmounts() []int32 {
	if o == nil || o.InputCollectionItemAmounts == nil {
		var ret []int32
		return ret
	}
	return o.InputCollectionItemAmounts
}

// GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetInputCollectionItemAmountsOk() ([]int32, bool) {
	if o == nil || o.InputCollectionItemAmounts == nil {
		return nil, false
	}
	return o.InputCollectionItemAmounts, true
}

// HasInputCollectionItemAmounts returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasInputCollectionItemAmounts() bool {
	if o != nil && o.InputCollectionItemAmounts != nil {
		return true
	}

	return false
}

// SetInputCollectionItemAmounts gets a reference to the given []int32 and assigns it to the InputCollectionItemAmounts field.
func (o *SetLootboxManagerLootboxRequest) SetInputCollectionItemAmounts(v []int32) {
	o.InputCollectionItemAmounts = v
}

// GetOutputCollectionAddress returns the OutputCollectionAddress field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionAddress() string {
	if o == nil || o.OutputCollectionAddress == nil {
		var ret string
		return ret
	}
	return *o.OutputCollectionAddress
}

// GetOutputCollectionAddressOk returns a tuple with the OutputCollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionAddressOk() (*string, bool) {
	if o == nil || o.OutputCollectionAddress == nil {
		return nil, false
	}
	return o.OutputCollectionAddress, true
}

// HasOutputCollectionAddress returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionAddress() bool {
	if o != nil && o.OutputCollectionAddress != nil {
		return true
	}

	return false
}

// SetOutputCollectionAddress gets a reference to the given string and assigns it to the OutputCollectionAddress field.
func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionAddress(v string) {
	o.OutputCollectionAddress = &v
}

// GetOutputCollectionId returns the OutputCollectionId field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionId() string {
	if o == nil || o.OutputCollectionId == nil {
		var ret string
		return ret
	}
	return *o.OutputCollectionId
}

// GetOutputCollectionIdOk returns a tuple with the OutputCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionIdOk() (*string, bool) {
	if o == nil || o.OutputCollectionId == nil {
		return nil, false
	}
	return o.OutputCollectionId, true
}

// HasOutputCollectionId returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionId() bool {
	if o != nil && o.OutputCollectionId != nil {
		return true
	}

	return false
}

// SetOutputCollectionId gets a reference to the given string and assigns it to the OutputCollectionId field.
func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionId(v string) {
	o.OutputCollectionId = &v
}

// GetOutputCollectionItemIds returns the OutputCollectionItemIds field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemIds() []int32 {
	if o == nil || o.OutputCollectionItemIds == nil {
		var ret []int32
		return ret
	}
	return o.OutputCollectionItemIds
}

// GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemIdsOk() ([]int32, bool) {
	if o == nil || o.OutputCollectionItemIds == nil {
		return nil, false
	}
	return o.OutputCollectionItemIds, true
}

// HasOutputCollectionItemIds returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionItemIds() bool {
	if o != nil && o.OutputCollectionItemIds != nil {
		return true
	}

	return false
}

// SetOutputCollectionItemIds gets a reference to the given []int32 and assigns it to the OutputCollectionItemIds field.
func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionItemIds(v []int32) {
	o.OutputCollectionItemIds = v
}

// GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemAmounts() []int32 {
	if o == nil || o.OutputCollectionItemAmounts == nil {
		var ret []int32
		return ret
	}
	return o.OutputCollectionItemAmounts
}

// GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemAmountsOk() ([]int32, bool) {
	if o == nil || o.OutputCollectionItemAmounts == nil {
		return nil, false
	}
	return o.OutputCollectionItemAmounts, true
}

// HasOutputCollectionItemAmounts returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionItemAmounts() bool {
	if o != nil && o.OutputCollectionItemAmounts != nil {
		return true
	}

	return false
}

// SetOutputCollectionItemAmounts gets a reference to the given []int32 and assigns it to the OutputCollectionItemAmounts field.
func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionItemAmounts(v []int32) {
	o.OutputCollectionItemAmounts = v
}

// GetOutputCollectionItemWeights returns the OutputCollectionItemWeights field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemWeights() []int32 {
	if o == nil || o.OutputCollectionItemWeights == nil {
		var ret []int32
		return ret
	}
	return o.OutputCollectionItemWeights
}

// GetOutputCollectionItemWeightsOk returns a tuple with the OutputCollectionItemWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetOutputCollectionItemWeightsOk() ([]int32, bool) {
	if o == nil || o.OutputCollectionItemWeights == nil {
		return nil, false
	}
	return o.OutputCollectionItemWeights, true
}

// HasOutputCollectionItemWeights returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasOutputCollectionItemWeights() bool {
	if o != nil && o.OutputCollectionItemWeights != nil {
		return true
	}

	return false
}

// SetOutputCollectionItemWeights gets a reference to the given []int32 and assigns it to the OutputCollectionItemWeights field.
func (o *SetLootboxManagerLootboxRequest) SetOutputCollectionItemWeights(v []int32) {
	o.OutputCollectionItemWeights = v
}

// GetOutputTotalItems returns the OutputTotalItems field value if set, zero value otherwise.
func (o *SetLootboxManagerLootboxRequest) GetOutputTotalItems() int32 {
	if o == nil || o.OutputTotalItems == nil {
		var ret int32
		return ret
	}
	return *o.OutputTotalItems
}

// GetOutputTotalItemsOk returns a tuple with the OutputTotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetLootboxManagerLootboxRequest) GetOutputTotalItemsOk() (*int32, bool) {
	if o == nil || o.OutputTotalItems == nil {
		return nil, false
	}
	return o.OutputTotalItems, true
}

// HasOutputTotalItems returns a boolean if a field has been set.
func (o *SetLootboxManagerLootboxRequest) HasOutputTotalItems() bool {
	if o != nil && o.OutputTotalItems != nil {
		return true
	}

	return false
}

// SetOutputTotalItems gets a reference to the given int32 and assigns it to the OutputTotalItems field.
func (o *SetLootboxManagerLootboxRequest) SetOutputTotalItems(v int32) {
	o.OutputTotalItems = &v
}

func (o SetLootboxManagerLootboxRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.InputCollectionAddress != nil {
		toSerialize["inputCollectionAddress"] = o.InputCollectionAddress
	}
	if o.InputCollectionId != nil {
		toSerialize["inputCollectionId"] = o.InputCollectionId
	}
	if o.InputCollectionItemIds != nil {
		toSerialize["inputCollectionItemIds"] = o.InputCollectionItemIds
	}
	if o.InputCollectionItemAmounts != nil {
		toSerialize["inputCollectionItemAmounts"] = o.InputCollectionItemAmounts
	}
	if o.OutputCollectionAddress != nil {
		toSerialize["outputCollectionAddress"] = o.OutputCollectionAddress
	}
	if o.OutputCollectionId != nil {
		toSerialize["outputCollectionId"] = o.OutputCollectionId
	}
	if o.OutputCollectionItemIds != nil {
		toSerialize["outputCollectionItemIds"] = o.OutputCollectionItemIds
	}
	if o.OutputCollectionItemAmounts != nil {
		toSerialize["outputCollectionItemAmounts"] = o.OutputCollectionItemAmounts
	}
	if o.OutputCollectionItemWeights != nil {
		toSerialize["outputCollectionItemWeights"] = o.OutputCollectionItemWeights
	}
	if o.OutputTotalItems != nil {
		toSerialize["outputTotalItems"] = o.OutputTotalItems
	}
	return json.Marshal(toSerialize)
}

type NullableSetLootboxManagerLootboxRequest struct {
	value *SetLootboxManagerLootboxRequest
	isSet bool
}

func (v NullableSetLootboxManagerLootboxRequest) Get() *SetLootboxManagerLootboxRequest {
	return v.value
}

func (v *NullableSetLootboxManagerLootboxRequest) Set(val *SetLootboxManagerLootboxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLootboxManagerLootboxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLootboxManagerLootboxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLootboxManagerLootboxRequest(val *SetLootboxManagerLootboxRequest) *NullableSetLootboxManagerLootboxRequest {
	return &NullableSetLootboxManagerLootboxRequest{value: val, isSet: true}
}

func (v NullableSetLootboxManagerLootboxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLootboxManagerLootboxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


