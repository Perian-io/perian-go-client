/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package perian

import (
	"encoding/json"
)

// checks if the InstanceTypeQueryInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceTypeQueryInput{}

// InstanceTypeQueryInput Query for the instance type attribute.
type InstanceTypeQueryInput struct {
	Operator NullableOperator `json:"operator,omitempty"`
	Options NullableQueryOptions `json:"options,omitempty"`
	All *bool `json:"all,omitempty"`
	Id NullableId `json:"id,omitempty"`
	Region NullableRegionQueryInput `json:"region,omitempty"`
	Zone NullableZoneQueryInput `json:"zone,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Cpu NullableCpuQueryInput `json:"cpu,omitempty"`
	Accelerator NullableAcceleratorQueryInput `json:"accelerator,omitempty"`
	Ram NullableMemoryQueryInput `json:"ram,omitempty"`
	Storage NullableStorageQueryInput `json:"storage,omitempty"`
	Network NullableNetworkQueryInput `json:"network,omitempty"`
	Price NullablePriceQueryInput `json:"price,omitempty"`
	Availability NullableAvailabilityQueryInput `json:"availability,omitempty"`
	Type NullableString `json:"type,omitempty"`
	BillingGranularity NullableString `json:"billing_granularity,omitempty"`
	Provider NullableProviderQueryInput `json:"provider,omitempty"`
	ReferenceId NullableReferenceIdQueryInput `json:"reference_id,omitempty"`
	AttributesHash NullableAttributesHash `json:"attributes_hash,omitempty"`
}

// NewInstanceTypeQueryInput instantiates a new InstanceTypeQueryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeQueryInput() *InstanceTypeQueryInput {
	this := InstanceTypeQueryInput{}
	var all bool = false
	this.All = &all
	return &this
}

// NewInstanceTypeQueryInputWithDefaults instantiates a new InstanceTypeQueryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeQueryInputWithDefaults() *InstanceTypeQueryInput {
	this := InstanceTypeQueryInput{}
	var all bool = false
	this.All = &all
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetOperator() Operator {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret Operator
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetOperatorOk() (*Operator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableOperator and assigns it to the Operator field.
func (o *InstanceTypeQueryInput) SetOperator(v Operator) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *InstanceTypeQueryInput) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetOperator() {
	o.Operator.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetOptions() QueryOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret QueryOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetOptionsOk() (*QueryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableQueryOptions and assigns it to the Options field.
func (o *InstanceTypeQueryInput) SetOptions(v QueryOptions) {
	o.Options.Set(&v)
}
// SetOptionsNil sets the value for Options to be an explicit nil
func (o *InstanceTypeQueryInput) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetOptions() {
	o.Options.Unset()
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *InstanceTypeQueryInput) GetAll() bool {
	if o == nil || IsNil(o.All) {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeQueryInput) GetAllOk() (*bool, bool) {
	if o == nil || IsNil(o.All) {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasAll() bool {
	if o != nil && !IsNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *InstanceTypeQueryInput) SetAll(v bool) {
	o.All = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetId() Id {
	if o == nil || IsNil(o.Id.Get()) {
		var ret Id
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetIdOk() (*Id, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableId and assigns it to the Id field.
func (o *InstanceTypeQueryInput) SetId(v Id) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InstanceTypeQueryInput) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetId() {
	o.Id.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetRegion() RegionQueryInput {
	if o == nil || IsNil(o.Region.Get()) {
		var ret RegionQueryInput
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetRegionOk() (*RegionQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableRegionQueryInput and assigns it to the Region field.
func (o *InstanceTypeQueryInput) SetRegion(v RegionQueryInput) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *InstanceTypeQueryInput) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetRegion() {
	o.Region.Unset()
}

// GetZone returns the Zone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetZone() ZoneQueryInput {
	if o == nil || IsNil(o.Zone.Get()) {
		var ret ZoneQueryInput
		return ret
	}
	return *o.Zone.Get()
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetZoneOk() (*ZoneQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zone.Get(), o.Zone.IsSet()
}

// HasZone returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasZone() bool {
	if o != nil && o.Zone.IsSet() {
		return true
	}

	return false
}

// SetZone gets a reference to the given NullableZoneQueryInput and assigns it to the Zone field.
func (o *InstanceTypeQueryInput) SetZone(v ZoneQueryInput) {
	o.Zone.Set(&v)
}
// SetZoneNil sets the value for Zone to be an explicit nil
func (o *InstanceTypeQueryInput) SetZoneNil() {
	o.Zone.Set(nil)
}

// UnsetZone ensures that no value is present for Zone, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetZone() {
	o.Zone.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InstanceTypeQueryInput) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InstanceTypeQueryInput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetCpu() CpuQueryInput {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret CpuQueryInput
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetCpuOk() (*CpuQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableCpuQueryInput and assigns it to the Cpu field.
func (o *InstanceTypeQueryInput) SetCpu(v CpuQueryInput) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *InstanceTypeQueryInput) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetCpu() {
	o.Cpu.Unset()
}

// GetAccelerator returns the Accelerator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetAccelerator() AcceleratorQueryInput {
	if o == nil || IsNil(o.Accelerator.Get()) {
		var ret AcceleratorQueryInput
		return ret
	}
	return *o.Accelerator.Get()
}

// GetAcceleratorOk returns a tuple with the Accelerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetAcceleratorOk() (*AcceleratorQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accelerator.Get(), o.Accelerator.IsSet()
}

// HasAccelerator returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasAccelerator() bool {
	if o != nil && o.Accelerator.IsSet() {
		return true
	}

	return false
}

// SetAccelerator gets a reference to the given NullableAcceleratorQueryInput and assigns it to the Accelerator field.
func (o *InstanceTypeQueryInput) SetAccelerator(v AcceleratorQueryInput) {
	o.Accelerator.Set(&v)
}
// SetAcceleratorNil sets the value for Accelerator to be an explicit nil
func (o *InstanceTypeQueryInput) SetAcceleratorNil() {
	o.Accelerator.Set(nil)
}

// UnsetAccelerator ensures that no value is present for Accelerator, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetAccelerator() {
	o.Accelerator.Unset()
}

// GetRam returns the Ram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetRam() MemoryQueryInput {
	if o == nil || IsNil(o.Ram.Get()) {
		var ret MemoryQueryInput
		return ret
	}
	return *o.Ram.Get()
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetRamOk() (*MemoryQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ram.Get(), o.Ram.IsSet()
}

// HasRam returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasRam() bool {
	if o != nil && o.Ram.IsSet() {
		return true
	}

	return false
}

// SetRam gets a reference to the given NullableMemoryQueryInput and assigns it to the Ram field.
func (o *InstanceTypeQueryInput) SetRam(v MemoryQueryInput) {
	o.Ram.Set(&v)
}
// SetRamNil sets the value for Ram to be an explicit nil
func (o *InstanceTypeQueryInput) SetRamNil() {
	o.Ram.Set(nil)
}

// UnsetRam ensures that no value is present for Ram, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetRam() {
	o.Ram.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetStorage() StorageQueryInput {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret StorageQueryInput
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetStorageOk() (*StorageQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableStorageQueryInput and assigns it to the Storage field.
func (o *InstanceTypeQueryInput) SetStorage(v StorageQueryInput) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *InstanceTypeQueryInput) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetStorage() {
	o.Storage.Unset()
}

// GetNetwork returns the Network field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetNetwork() NetworkQueryInput {
	if o == nil || IsNil(o.Network.Get()) {
		var ret NetworkQueryInput
		return ret
	}
	return *o.Network.Get()
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetNetworkOk() (*NetworkQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Network.Get(), o.Network.IsSet()
}

// HasNetwork returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasNetwork() bool {
	if o != nil && o.Network.IsSet() {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NullableNetworkQueryInput and assigns it to the Network field.
func (o *InstanceTypeQueryInput) SetNetwork(v NetworkQueryInput) {
	o.Network.Set(&v)
}
// SetNetworkNil sets the value for Network to be an explicit nil
func (o *InstanceTypeQueryInput) SetNetworkNil() {
	o.Network.Set(nil)
}

// UnsetNetwork ensures that no value is present for Network, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetNetwork() {
	o.Network.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetPrice() PriceQueryInput {
	if o == nil || IsNil(o.Price.Get()) {
		var ret PriceQueryInput
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetPriceOk() (*PriceQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullablePriceQueryInput and assigns it to the Price field.
func (o *InstanceTypeQueryInput) SetPrice(v PriceQueryInput) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *InstanceTypeQueryInput) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetPrice() {
	o.Price.Unset()
}

// GetAvailability returns the Availability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetAvailability() AvailabilityQueryInput {
	if o == nil || IsNil(o.Availability.Get()) {
		var ret AvailabilityQueryInput
		return ret
	}
	return *o.Availability.Get()
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetAvailabilityOk() (*AvailabilityQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Availability.Get(), o.Availability.IsSet()
}

// HasAvailability returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasAvailability() bool {
	if o != nil && o.Availability.IsSet() {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given NullableAvailabilityQueryInput and assigns it to the Availability field.
func (o *InstanceTypeQueryInput) SetAvailability(v AvailabilityQueryInput) {
	o.Availability.Set(&v)
}
// SetAvailabilityNil sets the value for Availability to be an explicit nil
func (o *InstanceTypeQueryInput) SetAvailabilityNil() {
	o.Availability.Set(nil)
}

// UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetAvailability() {
	o.Availability.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *InstanceTypeQueryInput) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *InstanceTypeQueryInput) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetType() {
	o.Type.Unset()
}

// GetBillingGranularity returns the BillingGranularity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetBillingGranularity() string {
	if o == nil || IsNil(o.BillingGranularity.Get()) {
		var ret string
		return ret
	}
	return *o.BillingGranularity.Get()
}

// GetBillingGranularityOk returns a tuple with the BillingGranularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetBillingGranularityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingGranularity.Get(), o.BillingGranularity.IsSet()
}

// HasBillingGranularity returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasBillingGranularity() bool {
	if o != nil && o.BillingGranularity.IsSet() {
		return true
	}

	return false
}

// SetBillingGranularity gets a reference to the given NullableString and assigns it to the BillingGranularity field.
func (o *InstanceTypeQueryInput) SetBillingGranularity(v string) {
	o.BillingGranularity.Set(&v)
}
// SetBillingGranularityNil sets the value for BillingGranularity to be an explicit nil
func (o *InstanceTypeQueryInput) SetBillingGranularityNil() {
	o.BillingGranularity.Set(nil)
}

// UnsetBillingGranularity ensures that no value is present for BillingGranularity, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetBillingGranularity() {
	o.BillingGranularity.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetProvider() ProviderQueryInput {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret ProviderQueryInput
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetProviderOk() (*ProviderQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableProviderQueryInput and assigns it to the Provider field.
func (o *InstanceTypeQueryInput) SetProvider(v ProviderQueryInput) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *InstanceTypeQueryInput) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetProvider() {
	o.Provider.Unset()
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetReferenceId() ReferenceIdQueryInput {
	if o == nil || IsNil(o.ReferenceId.Get()) {
		var ret ReferenceIdQueryInput
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetReferenceIdOk() (*ReferenceIdQueryInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableReferenceIdQueryInput and assigns it to the ReferenceId field.
func (o *InstanceTypeQueryInput) SetReferenceId(v ReferenceIdQueryInput) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *InstanceTypeQueryInput) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetAttributesHash returns the AttributesHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeQueryInput) GetAttributesHash() AttributesHash {
	if o == nil || IsNil(o.AttributesHash.Get()) {
		var ret AttributesHash
		return ret
	}
	return *o.AttributesHash.Get()
}

// GetAttributesHashOk returns a tuple with the AttributesHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeQueryInput) GetAttributesHashOk() (*AttributesHash, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributesHash.Get(), o.AttributesHash.IsSet()
}

// HasAttributesHash returns a boolean if a field has been set.
func (o *InstanceTypeQueryInput) HasAttributesHash() bool {
	if o != nil && o.AttributesHash.IsSet() {
		return true
	}

	return false
}

// SetAttributesHash gets a reference to the given NullableAttributesHash and assigns it to the AttributesHash field.
func (o *InstanceTypeQueryInput) SetAttributesHash(v AttributesHash) {
	o.AttributesHash.Set(&v)
}
// SetAttributesHashNil sets the value for AttributesHash to be an explicit nil
func (o *InstanceTypeQueryInput) SetAttributesHashNil() {
	o.AttributesHash.Set(nil)
}

// UnsetAttributesHash ensures that no value is present for AttributesHash, not even an explicit nil
func (o *InstanceTypeQueryInput) UnsetAttributesHash() {
	o.AttributesHash.Unset()
}

func (o InstanceTypeQueryInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceTypeQueryInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if !IsNil(o.All) {
		toSerialize["all"] = o.All
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Zone.IsSet() {
		toSerialize["zone"] = o.Zone.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.Accelerator.IsSet() {
		toSerialize["accelerator"] = o.Accelerator.Get()
	}
	if o.Ram.IsSet() {
		toSerialize["ram"] = o.Ram.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.Network.IsSet() {
		toSerialize["network"] = o.Network.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.Availability.IsSet() {
		toSerialize["availability"] = o.Availability.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.BillingGranularity.IsSet() {
		toSerialize["billing_granularity"] = o.BillingGranularity.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
	}
	if o.AttributesHash.IsSet() {
		toSerialize["attributes_hash"] = o.AttributesHash.Get()
	}
	return toSerialize, nil
}

type NullableInstanceTypeQueryInput struct {
	value *InstanceTypeQueryInput
	isSet bool
}

func (v NullableInstanceTypeQueryInput) Get() *InstanceTypeQueryInput {
	return v.value
}

func (v *NullableInstanceTypeQueryInput) Set(val *InstanceTypeQueryInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeQueryInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeQueryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeQueryInput(val *InstanceTypeQueryInput) *NullableInstanceTypeQueryInput {
	return &NullableInstanceTypeQueryInput{value: val, isSet: true}
}

func (v NullableInstanceTypeQueryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeQueryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


