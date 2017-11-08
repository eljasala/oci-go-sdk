// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type UpdateDatabaseDetails struct {
	DbBackupConfig *DbBackupConfig `mandatory:"false" json:"dbBackupConfig,omitempty"`
}

func (model UpdateDatabaseDetails) String() string {
	return common.PointerString(model)
}
