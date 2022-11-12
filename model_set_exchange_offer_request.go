/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.2.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"encoding/json"
)

// SetExchangeOfferRequest struct for SetExchangeOfferRequest
type SetExchangeOfferRequest struct {
	// A unique offer id to use for this offer for the exchange. If an existing offer id is used, the current offer will be updated but the existing number of uses will be kept. If you want to reset the number of uses for an existing offer, first remove it using the remove offer endpoint, then set it.
	Id float32 `json:"id"`
	// A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for input items required by this offer. `inputCollectionAddress` or `inputCollectionId` can optionally be provided.
	InputCollectionAddress *string `json:"inputCollectionAddress,omitempty"`
	// A valid MetaFab collection id that represents the collection for input items required by this offer. `inputCollectionAddress` or `inputCollectionId` can optionally be provided.
	InputCollectionId *string `json:"inputCollectionId,omitempty"`
	// An array of item ids from the provided input collection that are required to use this offer. Input items are transferred from the wallet to the exchange contract upon using an offer.
	InputCollectionItemIds []float32 `json:"inputCollectionItemIds,omitempty"`
	// An array of amounts for each item id from the provided input collection that are required to use this offer. Item amounts array indices are reflective of the amount required for a given item id at the same index.
	InputCollectionItemAmounts []float32 `json:"inputCollectionItemAmounts,omitempty"`
	// A valid EVM based ERC20 or MetaFab game currency contract address that for the currency required by this offer.
	InputCurrencyAddress *string `json:"inputCurrencyAddress,omitempty"`
	// A valid MetaFab currency id that represents the currency required by this offer.
	InputCurrencyId *string `json:"inputCurrencyId,omitempty"`
	// The amount of currency required by this offer. If an inputCurrencyAmount is provided without in input currency address or id, the native chain currency is used as the required currency.
	InputCurrencyAmount *float32 `json:"inputCurrencyAmount,omitempty"`
	// A valid EVM based ERC1155 or MetaFab game items contract address that represents the collection for output items given by this offer. `outputCollectionAddress` or `outputCollectionId` can optionally be provided.
	OutputCollectionAddress *string `json:"outputCollectionAddress,omitempty"`
	// A valid MetaFab collection id that represents the collection for output items given by this offer. `outputCollectionAddress` or `outputCollectionId` can optionally be provided.
	OutputCollectionId *string `json:"outputCollectionId,omitempty"`
	// An array of item ids from the provided output collection that are given by this offer. Output items are automatically minted if the exchange contract has the `minter` role for the output collection contract. Otherwise, they are transferred from the item balance held by the exchange contract.
	OutputCollectionItemIds []float32 `json:"outputCollectionItemIds,omitempty"`
	// An array of amounts for each item id from the provided output collection that are given by this offer. Item amounts array indices are reflective of the amount required for a given item id at the same index.
	OutputCollectionItemAmounts []float32 `json:"outputCollectionItemAmounts,omitempty"`
	// A valid EVM based ERC20 or MetaFab game currency contract address that for the currency given by this offer. The output currency amount is automatically minted if the exchange contract has the `minter` role for the output currency contract. Otherwise, they are transferred from the currency balance held by the exchange contract.
	OutputCurrencyAddress *string `json:"outputCurrencyAddress,omitempty"`
	// A valid MetaFab currency id for the currency given by this offer.
	OutputCurrencyId *string `json:"outputCurrencyId,omitempty"`
	// The amount of currency given by this offer. If an outputCurrencyAmount is provided without an output currency address or id, the native chain currency is used as the given currency.
	OutputCurrencyAmount *float32 `json:"outputCurrencyAmount,omitempty"`
	// The maximum number of times this offer can be used in total. maxUses is collective across all uses of the offer. If 5 unique players use an offer, that counts as 5 offer uses. Exclude this or use 0 to allow unlimited uses.
	MaxUses *float32 `json:"maxUses,omitempty"`
}

// NewSetExchangeOfferRequest instantiates a new SetExchangeOfferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetExchangeOfferRequest(id float32) *SetExchangeOfferRequest {
	this := SetExchangeOfferRequest{}
	this.Id = id
	return &this
}

// NewSetExchangeOfferRequestWithDefaults instantiates a new SetExchangeOfferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetExchangeOfferRequestWithDefaults() *SetExchangeOfferRequest {
	this := SetExchangeOfferRequest{}
	return &this
}

// GetId returns the Id field value
func (o *SetExchangeOfferRequest) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SetExchangeOfferRequest) SetId(v float32) {
	o.Id = v
}

// GetInputCollectionAddress returns the InputCollectionAddress field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCollectionAddress() string {
	if o == nil || o.InputCollectionAddress == nil {
		var ret string
		return ret
	}
	return *o.InputCollectionAddress
}

// GetInputCollectionAddressOk returns a tuple with the InputCollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCollectionAddressOk() (*string, bool) {
	if o == nil || o.InputCollectionAddress == nil {
		return nil, false
	}
	return o.InputCollectionAddress, true
}

// HasInputCollectionAddress returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCollectionAddress() bool {
	if o != nil && o.InputCollectionAddress != nil {
		return true
	}

	return false
}

// SetInputCollectionAddress gets a reference to the given string and assigns it to the InputCollectionAddress field.
func (o *SetExchangeOfferRequest) SetInputCollectionAddress(v string) {
	o.InputCollectionAddress = &v
}

// GetInputCollectionId returns the InputCollectionId field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCollectionId() string {
	if o == nil || o.InputCollectionId == nil {
		var ret string
		return ret
	}
	return *o.InputCollectionId
}

// GetInputCollectionIdOk returns a tuple with the InputCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCollectionIdOk() (*string, bool) {
	if o == nil || o.InputCollectionId == nil {
		return nil, false
	}
	return o.InputCollectionId, true
}

// HasInputCollectionId returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCollectionId() bool {
	if o != nil && o.InputCollectionId != nil {
		return true
	}

	return false
}

// SetInputCollectionId gets a reference to the given string and assigns it to the InputCollectionId field.
func (o *SetExchangeOfferRequest) SetInputCollectionId(v string) {
	o.InputCollectionId = &v
}

// GetInputCollectionItemIds returns the InputCollectionItemIds field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCollectionItemIds() []float32 {
	if o == nil || o.InputCollectionItemIds == nil {
		var ret []float32
		return ret
	}
	return o.InputCollectionItemIds
}

// GetInputCollectionItemIdsOk returns a tuple with the InputCollectionItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCollectionItemIdsOk() ([]float32, bool) {
	if o == nil || o.InputCollectionItemIds == nil {
		return nil, false
	}
	return o.InputCollectionItemIds, true
}

// HasInputCollectionItemIds returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCollectionItemIds() bool {
	if o != nil && o.InputCollectionItemIds != nil {
		return true
	}

	return false
}

// SetInputCollectionItemIds gets a reference to the given []float32 and assigns it to the InputCollectionItemIds field.
func (o *SetExchangeOfferRequest) SetInputCollectionItemIds(v []float32) {
	o.InputCollectionItemIds = v
}

// GetInputCollectionItemAmounts returns the InputCollectionItemAmounts field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCollectionItemAmounts() []float32 {
	if o == nil || o.InputCollectionItemAmounts == nil {
		var ret []float32
		return ret
	}
	return o.InputCollectionItemAmounts
}

// GetInputCollectionItemAmountsOk returns a tuple with the InputCollectionItemAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCollectionItemAmountsOk() ([]float32, bool) {
	if o == nil || o.InputCollectionItemAmounts == nil {
		return nil, false
	}
	return o.InputCollectionItemAmounts, true
}

// HasInputCollectionItemAmounts returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCollectionItemAmounts() bool {
	if o != nil && o.InputCollectionItemAmounts != nil {
		return true
	}

	return false
}

// SetInputCollectionItemAmounts gets a reference to the given []float32 and assigns it to the InputCollectionItemAmounts field.
func (o *SetExchangeOfferRequest) SetInputCollectionItemAmounts(v []float32) {
	o.InputCollectionItemAmounts = v
}

// GetInputCurrencyAddress returns the InputCurrencyAddress field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCurrencyAddress() string {
	if o == nil || o.InputCurrencyAddress == nil {
		var ret string
		return ret
	}
	return *o.InputCurrencyAddress
}

// GetInputCurrencyAddressOk returns a tuple with the InputCurrencyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCurrencyAddressOk() (*string, bool) {
	if o == nil || o.InputCurrencyAddress == nil {
		return nil, false
	}
	return o.InputCurrencyAddress, true
}

// HasInputCurrencyAddress returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCurrencyAddress() bool {
	if o != nil && o.InputCurrencyAddress != nil {
		return true
	}

	return false
}

// SetInputCurrencyAddress gets a reference to the given string and assigns it to the InputCurrencyAddress field.
func (o *SetExchangeOfferRequest) SetInputCurrencyAddress(v string) {
	o.InputCurrencyAddress = &v
}

// GetInputCurrencyId returns the InputCurrencyId field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCurrencyId() string {
	if o == nil || o.InputCurrencyId == nil {
		var ret string
		return ret
	}
	return *o.InputCurrencyId
}

// GetInputCurrencyIdOk returns a tuple with the InputCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCurrencyIdOk() (*string, bool) {
	if o == nil || o.InputCurrencyId == nil {
		return nil, false
	}
	return o.InputCurrencyId, true
}

// HasInputCurrencyId returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCurrencyId() bool {
	if o != nil && o.InputCurrencyId != nil {
		return true
	}

	return false
}

// SetInputCurrencyId gets a reference to the given string and assigns it to the InputCurrencyId field.
func (o *SetExchangeOfferRequest) SetInputCurrencyId(v string) {
	o.InputCurrencyId = &v
}

// GetInputCurrencyAmount returns the InputCurrencyAmount field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetInputCurrencyAmount() float32 {
	if o == nil || o.InputCurrencyAmount == nil {
		var ret float32
		return ret
	}
	return *o.InputCurrencyAmount
}

// GetInputCurrencyAmountOk returns a tuple with the InputCurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetInputCurrencyAmountOk() (*float32, bool) {
	if o == nil || o.InputCurrencyAmount == nil {
		return nil, false
	}
	return o.InputCurrencyAmount, true
}

// HasInputCurrencyAmount returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasInputCurrencyAmount() bool {
	if o != nil && o.InputCurrencyAmount != nil {
		return true
	}

	return false
}

// SetInputCurrencyAmount gets a reference to the given float32 and assigns it to the InputCurrencyAmount field.
func (o *SetExchangeOfferRequest) SetInputCurrencyAmount(v float32) {
	o.InputCurrencyAmount = &v
}

// GetOutputCollectionAddress returns the OutputCollectionAddress field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCollectionAddress() string {
	if o == nil || o.OutputCollectionAddress == nil {
		var ret string
		return ret
	}
	return *o.OutputCollectionAddress
}

// GetOutputCollectionAddressOk returns a tuple with the OutputCollectionAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCollectionAddressOk() (*string, bool) {
	if o == nil || o.OutputCollectionAddress == nil {
		return nil, false
	}
	return o.OutputCollectionAddress, true
}

// HasOutputCollectionAddress returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCollectionAddress() bool {
	if o != nil && o.OutputCollectionAddress != nil {
		return true
	}

	return false
}

// SetOutputCollectionAddress gets a reference to the given string and assigns it to the OutputCollectionAddress field.
func (o *SetExchangeOfferRequest) SetOutputCollectionAddress(v string) {
	o.OutputCollectionAddress = &v
}

// GetOutputCollectionId returns the OutputCollectionId field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCollectionId() string {
	if o == nil || o.OutputCollectionId == nil {
		var ret string
		return ret
	}
	return *o.OutputCollectionId
}

// GetOutputCollectionIdOk returns a tuple with the OutputCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCollectionIdOk() (*string, bool) {
	if o == nil || o.OutputCollectionId == nil {
		return nil, false
	}
	return o.OutputCollectionId, true
}

// HasOutputCollectionId returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCollectionId() bool {
	if o != nil && o.OutputCollectionId != nil {
		return true
	}

	return false
}

// SetOutputCollectionId gets a reference to the given string and assigns it to the OutputCollectionId field.
func (o *SetExchangeOfferRequest) SetOutputCollectionId(v string) {
	o.OutputCollectionId = &v
}

// GetOutputCollectionItemIds returns the OutputCollectionItemIds field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCollectionItemIds() []float32 {
	if o == nil || o.OutputCollectionItemIds == nil {
		var ret []float32
		return ret
	}
	return o.OutputCollectionItemIds
}

// GetOutputCollectionItemIdsOk returns a tuple with the OutputCollectionItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCollectionItemIdsOk() ([]float32, bool) {
	if o == nil || o.OutputCollectionItemIds == nil {
		return nil, false
	}
	return o.OutputCollectionItemIds, true
}

// HasOutputCollectionItemIds returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCollectionItemIds() bool {
	if o != nil && o.OutputCollectionItemIds != nil {
		return true
	}

	return false
}

// SetOutputCollectionItemIds gets a reference to the given []float32 and assigns it to the OutputCollectionItemIds field.
func (o *SetExchangeOfferRequest) SetOutputCollectionItemIds(v []float32) {
	o.OutputCollectionItemIds = v
}

// GetOutputCollectionItemAmounts returns the OutputCollectionItemAmounts field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCollectionItemAmounts() []float32 {
	if o == nil || o.OutputCollectionItemAmounts == nil {
		var ret []float32
		return ret
	}
	return o.OutputCollectionItemAmounts
}

// GetOutputCollectionItemAmountsOk returns a tuple with the OutputCollectionItemAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCollectionItemAmountsOk() ([]float32, bool) {
	if o == nil || o.OutputCollectionItemAmounts == nil {
		return nil, false
	}
	return o.OutputCollectionItemAmounts, true
}

// HasOutputCollectionItemAmounts returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCollectionItemAmounts() bool {
	if o != nil && o.OutputCollectionItemAmounts != nil {
		return true
	}

	return false
}

// SetOutputCollectionItemAmounts gets a reference to the given []float32 and assigns it to the OutputCollectionItemAmounts field.
func (o *SetExchangeOfferRequest) SetOutputCollectionItemAmounts(v []float32) {
	o.OutputCollectionItemAmounts = v
}

// GetOutputCurrencyAddress returns the OutputCurrencyAddress field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCurrencyAddress() string {
	if o == nil || o.OutputCurrencyAddress == nil {
		var ret string
		return ret
	}
	return *o.OutputCurrencyAddress
}

// GetOutputCurrencyAddressOk returns a tuple with the OutputCurrencyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCurrencyAddressOk() (*string, bool) {
	if o == nil || o.OutputCurrencyAddress == nil {
		return nil, false
	}
	return o.OutputCurrencyAddress, true
}

// HasOutputCurrencyAddress returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCurrencyAddress() bool {
	if o != nil && o.OutputCurrencyAddress != nil {
		return true
	}

	return false
}

// SetOutputCurrencyAddress gets a reference to the given string and assigns it to the OutputCurrencyAddress field.
func (o *SetExchangeOfferRequest) SetOutputCurrencyAddress(v string) {
	o.OutputCurrencyAddress = &v
}

// GetOutputCurrencyId returns the OutputCurrencyId field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCurrencyId() string {
	if o == nil || o.OutputCurrencyId == nil {
		var ret string
		return ret
	}
	return *o.OutputCurrencyId
}

// GetOutputCurrencyIdOk returns a tuple with the OutputCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCurrencyIdOk() (*string, bool) {
	if o == nil || o.OutputCurrencyId == nil {
		return nil, false
	}
	return o.OutputCurrencyId, true
}

// HasOutputCurrencyId returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCurrencyId() bool {
	if o != nil && o.OutputCurrencyId != nil {
		return true
	}

	return false
}

// SetOutputCurrencyId gets a reference to the given string and assigns it to the OutputCurrencyId field.
func (o *SetExchangeOfferRequest) SetOutputCurrencyId(v string) {
	o.OutputCurrencyId = &v
}

// GetOutputCurrencyAmount returns the OutputCurrencyAmount field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetOutputCurrencyAmount() float32 {
	if o == nil || o.OutputCurrencyAmount == nil {
		var ret float32
		return ret
	}
	return *o.OutputCurrencyAmount
}

// GetOutputCurrencyAmountOk returns a tuple with the OutputCurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetOutputCurrencyAmountOk() (*float32, bool) {
	if o == nil || o.OutputCurrencyAmount == nil {
		return nil, false
	}
	return o.OutputCurrencyAmount, true
}

// HasOutputCurrencyAmount returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasOutputCurrencyAmount() bool {
	if o != nil && o.OutputCurrencyAmount != nil {
		return true
	}

	return false
}

// SetOutputCurrencyAmount gets a reference to the given float32 and assigns it to the OutputCurrencyAmount field.
func (o *SetExchangeOfferRequest) SetOutputCurrencyAmount(v float32) {
	o.OutputCurrencyAmount = &v
}

// GetMaxUses returns the MaxUses field value if set, zero value otherwise.
func (o *SetExchangeOfferRequest) GetMaxUses() float32 {
	if o == nil || o.MaxUses == nil {
		var ret float32
		return ret
	}
	return *o.MaxUses
}

// GetMaxUsesOk returns a tuple with the MaxUses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetExchangeOfferRequest) GetMaxUsesOk() (*float32, bool) {
	if o == nil || o.MaxUses == nil {
		return nil, false
	}
	return o.MaxUses, true
}

// HasMaxUses returns a boolean if a field has been set.
func (o *SetExchangeOfferRequest) HasMaxUses() bool {
	if o != nil && o.MaxUses != nil {
		return true
	}

	return false
}

// SetMaxUses gets a reference to the given float32 and assigns it to the MaxUses field.
func (o *SetExchangeOfferRequest) SetMaxUses(v float32) {
	o.MaxUses = &v
}

func (o SetExchangeOfferRequest) MarshalJSON() ([]byte, error) {
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
	if o.InputCurrencyAddress != nil {
		toSerialize["inputCurrencyAddress"] = o.InputCurrencyAddress
	}
	if o.InputCurrencyId != nil {
		toSerialize["inputCurrencyId"] = o.InputCurrencyId
	}
	if o.InputCurrencyAmount != nil {
		toSerialize["inputCurrencyAmount"] = o.InputCurrencyAmount
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
	if o.OutputCurrencyAddress != nil {
		toSerialize["outputCurrencyAddress"] = o.OutputCurrencyAddress
	}
	if o.OutputCurrencyId != nil {
		toSerialize["outputCurrencyId"] = o.OutputCurrencyId
	}
	if o.OutputCurrencyAmount != nil {
		toSerialize["outputCurrencyAmount"] = o.OutputCurrencyAmount
	}
	if o.MaxUses != nil {
		toSerialize["maxUses"] = o.MaxUses
	}
	return json.Marshal(toSerialize)
}

type NullableSetExchangeOfferRequest struct {
	value *SetExchangeOfferRequest
	isSet bool
}

func (v NullableSetExchangeOfferRequest) Get() *SetExchangeOfferRequest {
	return v.value
}

func (v *NullableSetExchangeOfferRequest) Set(val *SetExchangeOfferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetExchangeOfferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetExchangeOfferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetExchangeOfferRequest(val *SetExchangeOfferRequest) *NullableSetExchangeOfferRequest {
	return &NullableSetExchangeOfferRequest{value: val, isSet: true}
}

func (v NullableSetExchangeOfferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetExchangeOfferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

