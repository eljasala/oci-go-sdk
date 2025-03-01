// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"strings"
)

// LifecycleDetailsEnum Enum with underlying type: string
type LifecycleDetailsEnum string

// Set of constants representing the allowable values for LifecycleDetailsEnum
const (
	LifecycleDetailsPendingWithOracle   LifecycleDetailsEnum = "PENDING_WITH_ORACLE"
	LifecycleDetailsPendingWithCustomer LifecycleDetailsEnum = "PENDING_WITH_CUSTOMER"
	LifecycleDetailsCloseRequested      LifecycleDetailsEnum = "CLOSE_REQUESTED"
	LifecycleDetailsClosed              LifecycleDetailsEnum = "CLOSED"
)

var mappingLifecycleDetailsEnum = map[string]LifecycleDetailsEnum{
	"PENDING_WITH_ORACLE":   LifecycleDetailsPendingWithOracle,
	"PENDING_WITH_CUSTOMER": LifecycleDetailsPendingWithCustomer,
	"CLOSE_REQUESTED":       LifecycleDetailsCloseRequested,
	"CLOSED":                LifecycleDetailsClosed,
}

var mappingLifecycleDetailsEnumLowerCase = map[string]LifecycleDetailsEnum{
	"pending_with_oracle":   LifecycleDetailsPendingWithOracle,
	"pending_with_customer": LifecycleDetailsPendingWithCustomer,
	"close_requested":       LifecycleDetailsCloseRequested,
	"closed":                LifecycleDetailsClosed,
}

// GetLifecycleDetailsEnumValues Enumerates the set of values for LifecycleDetailsEnum
func GetLifecycleDetailsEnumValues() []LifecycleDetailsEnum {
	values := make([]LifecycleDetailsEnum, 0)
	for _, v := range mappingLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetLifecycleDetailsEnumStringValues Enumerates the set of values in String for LifecycleDetailsEnum
func GetLifecycleDetailsEnumStringValues() []string {
	return []string{
		"PENDING_WITH_ORACLE",
		"PENDING_WITH_CUSTOMER",
		"CLOSE_REQUESTED",
		"CLOSED",
	}
}

// GetMappingLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLifecycleDetailsEnum(val string) (LifecycleDetailsEnum, bool) {
	enum, ok := mappingLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
