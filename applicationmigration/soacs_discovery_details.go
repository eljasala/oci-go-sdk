// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoacsDiscoveryDetails Credentials to access the Oracle SOA Cloud Service application in the source environment. When you create and update a migration,
// Application Migration connects to the application in the source environment with the supplied credentials and exports the domain
// configuration.
type SoacsDiscoveryDetails struct {

	// WebLogic administrator username for the Oracle SOA Cloud Service application in the source environment.
	WeblogicUser *string `mandatory:"true" json:"weblogicUser"`

	// Password for this user.
	WeblogicPassword *string `mandatory:"true" json:"weblogicPassword"`
}

func (m SoacsDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SoacsDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SoacsDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSoacsDiscoveryDetails SoacsDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSoacsDiscoveryDetails
	}{
		"SOACS",
		(MarshalTypeSoacsDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
