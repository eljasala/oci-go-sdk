// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to view indicators of compromise and related items. For more information, see Overview of Threat Intelligence (https://docs.cloud.oracle.com/Content/ThreatIntelligence/Concepts/threatintelligenceoverview.htm).
//

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IndicatorSummary Summary of a data signature observed on a network or host that indicates a potential security threat.
type IndicatorSummary struct {

	// The OCID of the indicator.
	Id *string `mandatory:"true" json:"id"`

	// Type of indicator
	Type IndicatorTypeEnum `mandatory:"true" json:"type"`

	// The value of indicator.
	Value *string `mandatory:"true" json:"value"`

	// Characteristics of the threat indicator based on previous observations or behavior. May include related tactics, techniques, and procedures.
	ThreatTypes []ThreatType `mandatory:"true" json:"threatTypes"`

	// The time the data was first seen for this indicator. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The last time this indicator was updated. It starts with the same value as timeCreated and is never empty. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Confidence is an integer from 0 to 100 that provides a measure of our certainty in the maliciousness of the indicator.  This confidence value is aggregated from the confidence in the threat types, attributes, and relationships to create an overall value for the indicator.
	Confidence *int `mandatory:"false" json:"confidence"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The state of the indicator.  It will always be ACTIVE.  This field is added for consistency.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m IndicatorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndicatorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIndicatorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIndicatorTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
