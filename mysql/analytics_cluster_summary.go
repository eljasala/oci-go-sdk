// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v48/common"
)

// AnalyticsClusterSummary DEPRECATED -- please use HeatWave API instead.
// A summary of an Analytics Cluster.
type AnalyticsClusterSummary struct {

	// The shape determines resources to allocate to the Analytics
	// Cluster nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing compute instances, of the
	// specified shape, in the Analytics Cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	// The current state of the MySQL Analytics Cluster.
	LifecycleState AnalyticsClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Analytics Cluster was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Analytics Cluster was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m AnalyticsClusterSummary) String() string {
	return common.PointerString(m)
}
