# GetOrganizationAuditLogsSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "Success"]
**Message** | Pointer to **string** |  | [optional] [default to "Operation was successful"]
**Detail** | Pointer to **string** |  | [optional] [default to ""]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 200]
**AuditLogs** | Pointer to [**[]RequestLog**](RequestLog.md) |  | [optional] [default to []]

## Methods

### NewGetOrganizationAuditLogsSuccess

`func NewGetOrganizationAuditLogsSuccess() *GetOrganizationAuditLogsSuccess`

NewGetOrganizationAuditLogsSuccess instantiates a new GetOrganizationAuditLogsSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAuditLogsSuccessWithDefaults

`func NewGetOrganizationAuditLogsSuccessWithDefaults() *GetOrganizationAuditLogsSuccess`

NewGetOrganizationAuditLogsSuccessWithDefaults instantiates a new GetOrganizationAuditLogsSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetOrganizationAuditLogsSuccess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationAuditLogsSuccess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationAuditLogsSuccess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationAuditLogsSuccess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetOrganizationAuditLogsSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetOrganizationAuditLogsSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetOrganizationAuditLogsSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetOrganizationAuditLogsSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *GetOrganizationAuditLogsSuccess) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetOrganizationAuditLogsSuccess) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetOrganizationAuditLogsSuccess) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GetOrganizationAuditLogsSuccess) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetStatusCode

`func (o *GetOrganizationAuditLogsSuccess) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GetOrganizationAuditLogsSuccess) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GetOrganizationAuditLogsSuccess) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GetOrganizationAuditLogsSuccess) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetAuditLogs

`func (o *GetOrganizationAuditLogsSuccess) GetAuditLogs() []RequestLog`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *GetOrganizationAuditLogsSuccess) GetAuditLogsOk() (*[]RequestLog, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *GetOrganizationAuditLogsSuccess) SetAuditLogs(v []RequestLog)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *GetOrganizationAuditLogsSuccess) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


