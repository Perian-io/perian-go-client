# ProviderProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeBillingGranularity** | Pointer to [**BillingGranularity**](BillingGranularity.md) |  | [optional] [default to BILLINGGRANULARITY_UNDEFINED]

## Methods

### NewProviderProperties

`func NewProviderProperties() *ProviderProperties`

NewProviderProperties instantiates a new ProviderProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderPropertiesWithDefaults

`func NewProviderPropertiesWithDefaults() *ProviderProperties`

NewProviderPropertiesWithDefaults instantiates a new ProviderProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeBillingGranularity

`func (o *ProviderProperties) GetComputeBillingGranularity() BillingGranularity`

GetComputeBillingGranularity returns the ComputeBillingGranularity field if non-nil, zero value otherwise.

### GetComputeBillingGranularityOk

`func (o *ProviderProperties) GetComputeBillingGranularityOk() (*BillingGranularity, bool)`

GetComputeBillingGranularityOk returns a tuple with the ComputeBillingGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBillingGranularity

`func (o *ProviderProperties) SetComputeBillingGranularity(v BillingGranularity)`

SetComputeBillingGranularity sets ComputeBillingGranularity field to given value.

### HasComputeBillingGranularity

`func (o *ProviderProperties) HasComputeBillingGranularity() bool`

HasComputeBillingGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


