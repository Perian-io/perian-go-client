# AcceleratorDataView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**No** | Pointer to **int32** |  | [optional] [default to 0]
**Memory** | Pointer to [**Memory**](Memory.md) |  | [optional] [default to {size=0.0, unit=Gb, bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined}]
**AcceleratorTypes** | Pointer to [**[]AcceleratorTypeView**](AcceleratorTypeView.md) |  | [optional] [default to []]

## Methods

### NewAcceleratorDataView

`func NewAcceleratorDataView() *AcceleratorDataView`

NewAcceleratorDataView instantiates a new AcceleratorDataView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceleratorDataViewWithDefaults

`func NewAcceleratorDataViewWithDefaults() *AcceleratorDataView`

NewAcceleratorDataViewWithDefaults instantiates a new AcceleratorDataView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNo

`func (o *AcceleratorDataView) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *AcceleratorDataView) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *AcceleratorDataView) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *AcceleratorDataView) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetMemory

`func (o *AcceleratorDataView) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AcceleratorDataView) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AcceleratorDataView) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AcceleratorDataView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetAcceleratorTypes

`func (o *AcceleratorDataView) GetAcceleratorTypes() []AcceleratorTypeView`

GetAcceleratorTypes returns the AcceleratorTypes field if non-nil, zero value otherwise.

### GetAcceleratorTypesOk

`func (o *AcceleratorDataView) GetAcceleratorTypesOk() (*[]AcceleratorTypeView, bool)`

GetAcceleratorTypesOk returns a tuple with the AcceleratorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleratorTypes

`func (o *AcceleratorDataView) SetAcceleratorTypes(v []AcceleratorTypeView)`

SetAcceleratorTypes sets AcceleratorTypes field to given value.

### HasAcceleratorTypes

`func (o *AcceleratorDataView) HasAcceleratorTypes() bool`

HasAcceleratorTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


