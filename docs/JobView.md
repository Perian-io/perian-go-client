# JobView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] [default to JOBSTATUS_UNDEFINED]
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**DoneAt** | Pointer to **NullableTime** |  | [optional] 
**Logs** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to [**PriceData**](PriceData.md) |  | [optional] [default to {unit_price=0, currency=Undefined, granularity=UNDEFINED, provider_billing_granularity=UNDEFINED}]
**DockerMetadata** | Pointer to [**NullableDockerMetadata**](DockerMetadata.md) |  | [optional] 
**RequirementMetadata** | Pointer to [**NullableRequirementMetadata**](RequirementMetadata.md) |  | [optional] 
**RuntimeMetadata** | Pointer to [**RuntimeMetadata**](RuntimeMetadata.md) |  | [optional] [default to {auto_failover_instance_type=false}]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewJobView

`func NewJobView() *JobView`

NewJobView instantiates a new JobView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobViewWithDefaults

`func NewJobViewWithDefaults() *JobView`

NewJobViewWithDefaults instantiates a new JobView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *JobView) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobView) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobView) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *JobView) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *JobView) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *JobView) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *JobView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *JobView) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *JobView) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetDoneAt

`func (o *JobView) GetDoneAt() time.Time`

GetDoneAt returns the DoneAt field if non-nil, zero value otherwise.

### GetDoneAtOk

`func (o *JobView) GetDoneAtOk() (*time.Time, bool)`

GetDoneAtOk returns a tuple with the DoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneAt

`func (o *JobView) SetDoneAt(v time.Time)`

SetDoneAt sets DoneAt field to given value.

### HasDoneAt

`func (o *JobView) HasDoneAt() bool`

HasDoneAt returns a boolean if a field has been set.

### SetDoneAtNil

`func (o *JobView) SetDoneAtNil(b bool)`

 SetDoneAtNil sets the value for DoneAt to be an explicit nil

### UnsetDoneAt
`func (o *JobView) UnsetDoneAt()`

UnsetDoneAt ensures that no value is present for DoneAt, not even an explicit nil
### GetLogs

`func (o *JobView) GetLogs() string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *JobView) GetLogsOk() (*string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *JobView) SetLogs(v string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *JobView) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogsNil

`func (o *JobView) SetLogsNil(b bool)`

 SetLogsNil sets the value for Logs to be an explicit nil

### UnsetLogs
`func (o *JobView) UnsetLogs()`

UnsetLogs ensures that no value is present for Logs, not even an explicit nil
### GetErrors

`func (o *JobView) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *JobView) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *JobView) SetErrors(v string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *JobView) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *JobView) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *JobView) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetPrice

`func (o *JobView) GetPrice() PriceData`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *JobView) GetPriceOk() (*PriceData, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *JobView) SetPrice(v PriceData)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *JobView) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDockerMetadata

`func (o *JobView) GetDockerMetadata() DockerMetadata`

GetDockerMetadata returns the DockerMetadata field if non-nil, zero value otherwise.

### GetDockerMetadataOk

`func (o *JobView) GetDockerMetadataOk() (*DockerMetadata, bool)`

GetDockerMetadataOk returns a tuple with the DockerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerMetadata

`func (o *JobView) SetDockerMetadata(v DockerMetadata)`

SetDockerMetadata sets DockerMetadata field to given value.

### HasDockerMetadata

`func (o *JobView) HasDockerMetadata() bool`

HasDockerMetadata returns a boolean if a field has been set.

### SetDockerMetadataNil

`func (o *JobView) SetDockerMetadataNil(b bool)`

 SetDockerMetadataNil sets the value for DockerMetadata to be an explicit nil

### UnsetDockerMetadata
`func (o *JobView) UnsetDockerMetadata()`

UnsetDockerMetadata ensures that no value is present for DockerMetadata, not even an explicit nil
### GetRequirementMetadata

`func (o *JobView) GetRequirementMetadata() RequirementMetadata`

GetRequirementMetadata returns the RequirementMetadata field if non-nil, zero value otherwise.

### GetRequirementMetadataOk

`func (o *JobView) GetRequirementMetadataOk() (*RequirementMetadata, bool)`

GetRequirementMetadataOk returns a tuple with the RequirementMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementMetadata

`func (o *JobView) SetRequirementMetadata(v RequirementMetadata)`

SetRequirementMetadata sets RequirementMetadata field to given value.

### HasRequirementMetadata

`func (o *JobView) HasRequirementMetadata() bool`

HasRequirementMetadata returns a boolean if a field has been set.

### SetRequirementMetadataNil

`func (o *JobView) SetRequirementMetadataNil(b bool)`

 SetRequirementMetadataNil sets the value for RequirementMetadata to be an explicit nil

### UnsetRequirementMetadata
`func (o *JobView) UnsetRequirementMetadata()`

UnsetRequirementMetadata ensures that no value is present for RequirementMetadata, not even an explicit nil
### GetRuntimeMetadata

`func (o *JobView) GetRuntimeMetadata() RuntimeMetadata`

GetRuntimeMetadata returns the RuntimeMetadata field if non-nil, zero value otherwise.

### GetRuntimeMetadataOk

`func (o *JobView) GetRuntimeMetadataOk() (*RuntimeMetadata, bool)`

GetRuntimeMetadataOk returns a tuple with the RuntimeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeMetadata

`func (o *JobView) SetRuntimeMetadata(v RuntimeMetadata)`

SetRuntimeMetadata sets RuntimeMetadata field to given value.

### HasRuntimeMetadata

`func (o *JobView) HasRuntimeMetadata() bool`

HasRuntimeMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *JobView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JobView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *JobView) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JobView) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JobView) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *JobView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


