// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v48/common"
)

// UpdateEmManagedExternalDatabaseInsightDetails The information to be updated.
type UpdateEmManagedExternalDatabaseInsightDetails struct {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

//GetFreeformTags returns FreeformTags
func (m UpdateEmManagedExternalDatabaseInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateEmManagedExternalDatabaseInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateEmManagedExternalDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateEmManagedExternalDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateEmManagedExternalDatabaseInsightDetails UpdateEmManagedExternalDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeUpdateEmManagedExternalDatabaseInsightDetails
	}{
		"EM_MANAGED_EXTERNAL_DATABASE",
		(MarshalTypeUpdateEmManagedExternalDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}
