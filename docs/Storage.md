# Storage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**StorageType**](StorageType.md) |  | [optional] [default to UNDEFINED]
**Size** | Pointer to [**Memory**](Memory.md) |  | [optional] [default to {size=0.0, unit=Gb, bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined}]
**Speed** | Pointer to [**Bandwidth**](Bandwidth.md) |  | [optional] [default to {speed=0.0, maximum=0.0, minimum=0.0, unit=Undefined, sla=Undefined, limit=Undefined}]

## Methods

### NewStorage

`func NewStorage() *Storage`

NewStorage instantiates a new Storage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageWithDefaults

`func NewStorageWithDefaults() *Storage`

NewStorageWithDefaults instantiates a new Storage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Storage) GetType() StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Storage) GetTypeOk() (*StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Storage) SetType(v StorageType)`

SetType sets Type field to given value.

### HasType

`func (o *Storage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *Storage) GetSize() Memory`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Storage) GetSizeOk() (*Memory, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Storage) SetSize(v Memory)`

SetSize sets Size field to given value.

### HasSize

`func (o *Storage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSpeed

`func (o *Storage) GetSpeed() Bandwidth`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *Storage) GetSpeedOk() (*Bandwidth, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *Storage) SetSpeed(v Bandwidth)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *Storage) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


