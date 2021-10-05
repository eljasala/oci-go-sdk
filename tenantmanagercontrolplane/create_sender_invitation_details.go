// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// CreateSenderInvitationDetails The parameters for creating a sender invitation.
type CreateSenderInvitationDetails struct {

	// OCID of the sender tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the recipient tenancy.
	RecipientTenancyId *string `mandatory:"true" json:"recipientTenancyId"`

	// Email address of the recipient.
	RecipientEmailAddress *string `mandatory:"false" json:"recipientEmailAddress"`

	// A user-created name to describe the invitation. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSenderInvitationDetails) String() string {
	return common.PointerString(m)
}
