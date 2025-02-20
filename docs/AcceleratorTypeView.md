# AcceleratorTypeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**AcceleratorName**](AcceleratorName.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to [**AcceleratorVendor**](AcceleratorVendor.md) |  | [optional] [default to UNDEFINED]
**Memory** | Pointer to [**Memory**](Memory.md) |  | [optional] [default to {size=0.0, unit=Gb, bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined}]

## Methods

### NewAcceleratorTypeView

`func NewAcceleratorTypeView() *AcceleratorTypeView`

NewAcceleratorTypeView instantiates a new AcceleratorTypeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceleratorTypeViewWithDefaults

`func NewAcceleratorTypeViewWithDefaults() *AcceleratorTypeView`

NewAcceleratorTypeViewWithDefaults instantiates a new AcceleratorTypeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AcceleratorTypeView) GetName() AcceleratorName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcceleratorTypeView) GetNameOk() (*AcceleratorName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcceleratorTypeView) SetName(v AcceleratorName)`

SetName sets Name field to given value.

### HasName

`func (o *AcceleratorTypeView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AcceleratorTypeView) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AcceleratorTypeView) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AcceleratorTypeView) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AcceleratorTypeView) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *AcceleratorTypeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcceleratorTypeView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcceleratorTypeView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AcceleratorTypeView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVendor

`func (o *AcceleratorTypeView) GetVendor() AcceleratorVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AcceleratorTypeView) GetVendorOk() (*AcceleratorVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AcceleratorTypeView) SetVendor(v AcceleratorVendor)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AcceleratorTypeView) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetMemory

`func (o *AcceleratorTypeView) GetMemory() Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AcceleratorTypeView) GetMemoryOk() (*Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AcceleratorTypeView) SetMemory(v Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AcceleratorTypeView) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


