// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IncidentResourceType Details about the resource associated with the support request.
type IncidentResourceType struct {

	// The label associated with the resource.
	Label *string `mandatory:"true" json:"label"`

	// Unique identifier of the resource.
	ResourceTypeKey *string `mandatory:"false" json:"resourceTypeKey"`

	// The display name of the resource.
	Name *string `mandatory:"false" json:"name"`

	// The description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The service category list.
	ServiceCategoryList []ServiceCategory `mandatory:"false" json:"serviceCategoryList"`
}

func (m IncidentResourceType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IncidentResourceType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
