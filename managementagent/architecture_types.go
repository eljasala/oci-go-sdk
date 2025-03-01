// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"strings"
)

// ArchitectureTypesEnum Enum with underlying type: string
type ArchitectureTypesEnum string

// Set of constants representing the allowable values for ArchitectureTypesEnum
const (
	ArchitectureTypesX8664 ArchitectureTypesEnum = "X86_64"
	ArchitectureTypesSparc ArchitectureTypesEnum = "SPARC"
)

var mappingArchitectureTypesEnum = map[string]ArchitectureTypesEnum{
	"X86_64": ArchitectureTypesX8664,
	"SPARC":  ArchitectureTypesSparc,
}

var mappingArchitectureTypesEnumLowerCase = map[string]ArchitectureTypesEnum{
	"x86_64": ArchitectureTypesX8664,
	"sparc":  ArchitectureTypesSparc,
}

// GetArchitectureTypesEnumValues Enumerates the set of values for ArchitectureTypesEnum
func GetArchitectureTypesEnumValues() []ArchitectureTypesEnum {
	values := make([]ArchitectureTypesEnum, 0)
	for _, v := range mappingArchitectureTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetArchitectureTypesEnumStringValues Enumerates the set of values in String for ArchitectureTypesEnum
func GetArchitectureTypesEnumStringValues() []string {
	return []string{
		"X86_64",
		"SPARC",
	}
}

// GetMappingArchitectureTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArchitectureTypesEnum(val string) (ArchitectureTypesEnum, bool) {
	enum, ok := mappingArchitectureTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
