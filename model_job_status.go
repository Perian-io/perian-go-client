/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package perian

import (
	"encoding/json"
	"fmt"
)

// JobStatus Enum class representing the status of a job.
type JobStatus string

// List of JobStatus
const (
	JOBSTATUS_QUEUED JobStatus = "Queued"
	JOBSTATUS_INITIALIZING JobStatus = "Initializing"
	JOBSTATUS_RUNNING JobStatus = "Running"
	JOBSTATUS_DONE JobStatus = "Done"
	JOBSTATUS_SERVER_ERROR JobStatus = "ServerError"
	JOBSTATUS_USER_ERROR JobStatus = "UserError"
	JOBSTATUS_CANCELLED JobStatus = "Cancelled"
	JOBSTATUS_TIMEOUT JobStatus = "Timeout"
	JOBSTATUS_UNDEFINED JobStatus = "Undefined"
)

// All allowed values of JobStatus enum
var AllowedJobStatusEnumValues = []JobStatus{
	"Queued",
	"Initializing",
	"Running",
	"Done",
	"ServerError",
	"UserError",
	"Cancelled",
	"Timeout",
	"Undefined",
}

func (v *JobStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobStatus(value)
	for _, existing := range AllowedJobStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobStatus", value)
}

// NewJobStatusFromValue returns a pointer to a valid JobStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobStatusFromValue(v string) (*JobStatus, error) {
	ev := JobStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobStatus: valid values are %v", v, AllowedJobStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobStatus) IsValid() bool {
	for _, existing := range AllowedJobStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobStatus value
func (v JobStatus) Ptr() *JobStatus {
	return &v
}

type NullableJobStatus struct {
	value *JobStatus
	isSet bool
}

func (v NullableJobStatus) Get() *JobStatus {
	return v.value
}

func (v *NullableJobStatus) Set(val *JobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStatus(val *JobStatus) *NullableJobStatus {
	return &NullableJobStatus{value: val, isSet: true}
}

func (v NullableJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

