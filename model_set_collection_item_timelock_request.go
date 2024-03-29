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

// SetCollectionItemTimelockRequest struct for SetCollectionItemTimelockRequest
type SetCollectionItemTimelockRequest struct {
	// A unix timestamp (in seconds) defining when the set timelock expires.
	Timelock int32 `json:"timelock"`
}

// NewSetCollectionItemTimelockRequest instantiates a new SetCollectionItemTimelockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCollectionItemTimelockRequest(timelock int32) *SetCollectionItemTimelockRequest {
	this := SetCollectionItemTimelockRequest{}
	this.Timelock = timelock
	return &this
}

// NewSetCollectionItemTimelockRequestWithDefaults instantiates a new SetCollectionItemTimelockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCollectionItemTimelockRequestWithDefaults() *SetCollectionItemTimelockRequest {
	this := SetCollectionItemTimelockRequest{}
	return &this
}

// GetTimelock returns the Timelock field value
func (o *SetCollectionItemTimelockRequest) GetTimelock() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timelock
}

// GetTimelockOk returns a tuple with the Timelock field value
// and a boolean to check if the value has been set.
func (o *SetCollectionItemTimelockRequest) GetTimelockOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timelock, true
}

// SetTimelock sets field value
func (o *SetCollectionItemTimelockRequest) SetTimelock(v int32) {
	o.Timelock = v
}

func (o SetCollectionItemTimelockRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timelock"] = o.Timelock
	}
	return json.Marshal(toSerialize)
}

type NullableSetCollectionItemTimelockRequest struct {
	value *SetCollectionItemTimelockRequest
	isSet bool
}

func (v NullableSetCollectionItemTimelockRequest) Get() *SetCollectionItemTimelockRequest {
	return v.value
}

func (v *NullableSetCollectionItemTimelockRequest) Set(val *SetCollectionItemTimelockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCollectionItemTimelockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCollectionItemTimelockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCollectionItemTimelockRequest(val *SetCollectionItemTimelockRequest) *NullableSetCollectionItemTimelockRequest {
	return &NullableSetCollectionItemTimelockRequest{value: val, isSet: true}
}

func (v NullableSetCollectionItemTimelockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCollectionItemTimelockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


