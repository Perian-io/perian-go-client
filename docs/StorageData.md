# StorageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**No** | Pointer to **int32** |  | [optional] [default to 0]
**Size** | Pointer to [**Memory**](Memory.md) |  | [optional] [default to {size=0.0, unit=Gb, bandwidth={limit=Undefined, maximum=0.0, minimum=0.0, sla=Undefined, speed=0.0, unit=Undefined}, interface=Undefined}]
**Included** | Pointer to [**StorageIncluded**](StorageIncluded.md) |  | [optional] [default to STORAGEINCLUDED_UNDEFINED]
**Storages** | Pointer to [**[]Storage**](Storage.md) |  | [optional] [default to []]

## Methods

### NewStorageData

`func NewStorageData() *StorageData`

NewStorageData instantiates a new StorageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageDataWithDefaults

`func NewStorageDataWithDefaults() *StorageData`

NewStorageDataWithDefaults instantiates a new StorageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNo

`func (o *StorageData) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *StorageData) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *StorageData) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *StorageData) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetSize

`func (o *StorageData) GetSize() Memory`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageData) GetSizeOk() (*Memory, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageData) SetSize(v Memory)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageData) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIncluded

`func (o *StorageData) GetIncluded() StorageIncluded`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *StorageData) GetIncludedOk() (*StorageIncluded, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *StorageData) SetIncluded(v StorageIncluded)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *StorageData) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetStorages

`func (o *StorageData) GetStorages() []Storage`

GetStorages returns the Storages field if non-nil, zero value otherwise.

### GetStoragesOk

`func (o *StorageData) GetStoragesOk() (*[]Storage, bool)`

GetStoragesOk returns a tuple with the Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorages

`func (o *StorageData) SetStorages(v []Storage)`

SetStorages sets Storages field to given value.

### HasStorages

`func (o *StorageData) HasStorages() bool`

HasStorages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


