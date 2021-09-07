// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// CreateExternalPublicationDetails Properties used to publish an Oracle Cloud Infrastructure Data Flow object.
type CreateExternalPublicationDetails struct {

	// The OCID of the compartment where the application is created in the Oracle Cloud Infrastructure Data Flow Service.
	ApplicationCompartmentId *string `mandatory:"true" json:"applicationCompartmentId"`

	// The name of the application.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique OCID of the identifier that is returned after creating the Oracle Cloud Infrastructure Data Flow application.
	ApplicationId *string `mandatory:"false" json:"applicationId"`

	// The details of the data flow or the application.
	Description *string `mandatory:"false" json:"description"`

	ResourceConfiguration *ResourceConfiguration `mandatory:"false" json:"resourceConfiguration"`

	ConfigurationDetails *ConfigurationDetails `mandatory:"false" json:"configurationDetails"`
}

func (m CreateExternalPublicationDetails) String() string {
	return common.PointerString(m)
}
