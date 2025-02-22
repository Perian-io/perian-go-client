# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inbound** | Pointer to [**Bandwidth**](Bandwidth.md) |  | [optional] [default to {speed=0.0, maximum=0.0, minimum=0.0, unit=Undefined, sla=Undefined, limit=Undefined}]
**Outbound** | Pointer to [**Bandwidth**](Bandwidth.md) |  | [optional] [default to {speed=0.0, maximum=0.0, minimum=0.0, unit=Undefined, sla=Undefined, limit=Undefined}]

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInbound

`func (o *Network) GetInbound() Bandwidth`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *Network) GetInboundOk() (*Bandwidth, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *Network) SetInbound(v Bandwidth)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *Network) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *Network) GetOutbound() Bandwidth`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *Network) GetOutboundOk() (*Bandwidth, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *Network) SetOutbound(v Bandwidth)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *Network) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


