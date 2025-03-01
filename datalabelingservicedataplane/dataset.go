// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Dataset A dataset is a logical collection of records. The dataset contains all the information necessary to describe a record's source, format, the type of annotations allowed for the record, and the labels allowed on annotations.
type Dataset struct {

	// The OCID of the dataset.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment of the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The state of a dataset.
	// CREATING - The dataset is being created.  It transitions to ACTIVE when it is ready for labeling.
	// ACTIVE   - The dataset is ready for labeling.
	// UPDATING - The dataset is being updated.  It, and its related resources, might be unavailable for other updates until it returns to ACTIVE.
	// NEEDS_ATTENTION - A dataset updaten operation has failed due to validation or other errors, and needs attention.
	// DELETING - The dataset and its related resources are being deleted.
	// DELETED  - The dataset has been deleted and is no longer available.
	// FAILED   - The dataset has failed due to validation or other errors.
	LifecycleState DatasetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The annotation format name required for labeling records.
	AnnotationFormat *string `mandatory:"true" json:"annotationFormat"`

	DatasetSourceDetails DatasetSourceDetails `mandatory:"true" json:"datasetSourceDetails"`

	DatasetFormatDetails DatasetFormatDetails `mandatory:"true" json:"datasetFormatDetails"`

	LabelSet *LabelSet `mandatory:"true" json:"labelSet"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-provided description of the dataset
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in FAILED or NEEDS_ATTENTION state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	InitialRecordGenerationConfiguration *InitialRecordGenerationConfiguration `mandatory:"false" json:"initialRecordGenerationConfiguration"`

	// The labeling instructions for human labelers in rich text format
	LabelingInstructions *string `mandatory:"false" json:"labelingInstructions"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The usage of system tag keys. These predefined keys are scoped to namespaces.
	// For example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Dataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Dataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatasetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatasetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Dataset) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                          *string                               `json:"displayName"`
		Description                          *string                               `json:"description"`
		LifecycleDetails                     *string                               `json:"lifecycleDetails"`
		InitialRecordGenerationConfiguration *InitialRecordGenerationConfiguration `json:"initialRecordGenerationConfiguration"`
		LabelingInstructions                 *string                               `json:"labelingInstructions"`
		FreeformTags                         map[string]string                     `json:"freeformTags"`
		DefinedTags                          map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags                           map[string]map[string]interface{}     `json:"systemTags"`
		Id                                   *string                               `json:"id"`
		CompartmentId                        *string                               `json:"compartmentId"`
		TimeCreated                          *common.SDKTime                       `json:"timeCreated"`
		TimeUpdated                          *common.SDKTime                       `json:"timeUpdated"`
		LifecycleState                       DatasetLifecycleStateEnum             `json:"lifecycleState"`
		AnnotationFormat                     *string                               `json:"annotationFormat"`
		DatasetSourceDetails                 datasetsourcedetails                  `json:"datasetSourceDetails"`
		DatasetFormatDetails                 datasetformatdetails                  `json:"datasetFormatDetails"`
		LabelSet                             *LabelSet                             `json:"labelSet"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.InitialRecordGenerationConfiguration = model.InitialRecordGenerationConfiguration

	m.LabelingInstructions = model.LabelingInstructions

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.AnnotationFormat = model.AnnotationFormat

	nn, e = model.DatasetSourceDetails.UnmarshalPolymorphicJSON(model.DatasetSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetSourceDetails = nn.(DatasetSourceDetails)
	} else {
		m.DatasetSourceDetails = nil
	}

	nn, e = model.DatasetFormatDetails.UnmarshalPolymorphicJSON(model.DatasetFormatDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetFormatDetails = nn.(DatasetFormatDetails)
	} else {
		m.DatasetFormatDetails = nil
	}

	m.LabelSet = model.LabelSet

	return
}

// DatasetLifecycleStateEnum Enum with underlying type: string
type DatasetLifecycleStateEnum string

// Set of constants representing the allowable values for DatasetLifecycleStateEnum
const (
	DatasetLifecycleStateCreating       DatasetLifecycleStateEnum = "CREATING"
	DatasetLifecycleStateUpdating       DatasetLifecycleStateEnum = "UPDATING"
	DatasetLifecycleStateActive         DatasetLifecycleStateEnum = "ACTIVE"
	DatasetLifecycleStateNeedsAttention DatasetLifecycleStateEnum = "NEEDS_ATTENTION"
	DatasetLifecycleStateDeleting       DatasetLifecycleStateEnum = "DELETING"
	DatasetLifecycleStateDeleted        DatasetLifecycleStateEnum = "DELETED"
	DatasetLifecycleStateFailed         DatasetLifecycleStateEnum = "FAILED"
)

var mappingDatasetLifecycleStateEnum = map[string]DatasetLifecycleStateEnum{
	"CREATING":        DatasetLifecycleStateCreating,
	"UPDATING":        DatasetLifecycleStateUpdating,
	"ACTIVE":          DatasetLifecycleStateActive,
	"NEEDS_ATTENTION": DatasetLifecycleStateNeedsAttention,
	"DELETING":        DatasetLifecycleStateDeleting,
	"DELETED":         DatasetLifecycleStateDeleted,
	"FAILED":          DatasetLifecycleStateFailed,
}

var mappingDatasetLifecycleStateEnumLowerCase = map[string]DatasetLifecycleStateEnum{
	"creating":        DatasetLifecycleStateCreating,
	"updating":        DatasetLifecycleStateUpdating,
	"active":          DatasetLifecycleStateActive,
	"needs_attention": DatasetLifecycleStateNeedsAttention,
	"deleting":        DatasetLifecycleStateDeleting,
	"deleted":         DatasetLifecycleStateDeleted,
	"failed":          DatasetLifecycleStateFailed,
}

// GetDatasetLifecycleStateEnumValues Enumerates the set of values for DatasetLifecycleStateEnum
func GetDatasetLifecycleStateEnumValues() []DatasetLifecycleStateEnum {
	values := make([]DatasetLifecycleStateEnum, 0)
	for _, v := range mappingDatasetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatasetLifecycleStateEnumStringValues Enumerates the set of values in String for DatasetLifecycleStateEnum
func GetDatasetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDatasetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatasetLifecycleStateEnum(val string) (DatasetLifecycleStateEnum, bool) {
	enum, ok := mappingDatasetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
