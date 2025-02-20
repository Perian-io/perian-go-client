# RequestLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Outcome** | Pointer to **NullableString** |  | [optional] 
**Ip** | Pointer to **NullableString** |  | [optional] 
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**RequestUrl** | Pointer to **NullableString** |  | [optional] 
**RequestMethod** | Pointer to **NullableString** |  | [optional] 
**RequestPayload** | Pointer to **NullableString** |  | [optional] 
**ResponseStatusCode** | Pointer to **NullableInt32** |  | [optional] 
**ResponseBody** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRequestLog

`func NewRequestLog(organizationId string, timestamp time.Time, ) *RequestLog`

NewRequestLog instantiates a new RequestLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestLogWithDefaults

`func NewRequestLogWithDefaults() *RequestLog`

NewRequestLogWithDefaults instantiates a new RequestLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestLog) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RequestLog) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RequestLog) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOrganizationId

`func (o *RequestLog) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RequestLog) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RequestLog) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetTimestamp

`func (o *RequestLog) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RequestLog) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RequestLog) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOutcome

`func (o *RequestLog) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *RequestLog) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *RequestLog) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *RequestLog) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *RequestLog) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *RequestLog) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetIp

`func (o *RequestLog) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RequestLog) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RequestLog) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RequestLog) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *RequestLog) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *RequestLog) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetUserAgent

`func (o *RequestLog) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RequestLog) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RequestLog) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *RequestLog) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *RequestLog) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *RequestLog) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetRequestUrl

`func (o *RequestLog) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *RequestLog) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *RequestLog) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *RequestLog) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### SetRequestUrlNil

`func (o *RequestLog) SetRequestUrlNil(b bool)`

 SetRequestUrlNil sets the value for RequestUrl to be an explicit nil

### UnsetRequestUrl
`func (o *RequestLog) UnsetRequestUrl()`

UnsetRequestUrl ensures that no value is present for RequestUrl, not even an explicit nil
### GetRequestMethod

`func (o *RequestLog) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *RequestLog) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *RequestLog) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *RequestLog) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### SetRequestMethodNil

`func (o *RequestLog) SetRequestMethodNil(b bool)`

 SetRequestMethodNil sets the value for RequestMethod to be an explicit nil

### UnsetRequestMethod
`func (o *RequestLog) UnsetRequestMethod()`

UnsetRequestMethod ensures that no value is present for RequestMethod, not even an explicit nil
### GetRequestPayload

`func (o *RequestLog) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *RequestLog) GetRequestPayloadOk() (*string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *RequestLog) SetRequestPayload(v string)`

SetRequestPayload sets RequestPayload field to given value.

### HasRequestPayload

`func (o *RequestLog) HasRequestPayload() bool`

HasRequestPayload returns a boolean if a field has been set.

### SetRequestPayloadNil

`func (o *RequestLog) SetRequestPayloadNil(b bool)`

 SetRequestPayloadNil sets the value for RequestPayload to be an explicit nil

### UnsetRequestPayload
`func (o *RequestLog) UnsetRequestPayload()`

UnsetRequestPayload ensures that no value is present for RequestPayload, not even an explicit nil
### GetResponseStatusCode

`func (o *RequestLog) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *RequestLog) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *RequestLog) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *RequestLog) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *RequestLog) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *RequestLog) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetResponseBody

`func (o *RequestLog) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *RequestLog) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *RequestLog) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *RequestLog) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *RequestLog) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *RequestLog) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


