// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// SyncPrimaryStorageCapacityParamDetail SyncPrimaryStorageCapacity detail param
type SyncPrimaryStorageCapacityParamDetail struct {
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
}

// SyncPrimaryStorageCapacityParam SyncPrimaryStorageCapacity request param
type SyncPrimaryStorageCapacityParam struct {
	BaseParam
	SyncPrimaryStorageCapacity SyncPrimaryStorageCapacityParamDetail `json:"syncPrimaryStorageCapacity"`
}
// GetPrimaryStorageCapacityParamDetail GetPrimaryStorageCapacity detail param
type GetPrimaryStorageCapacityParamDetail struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	PrimaryStorageUuids []string `json:"primaryStorageUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetPrimaryStorageCapacityParam GetPrimaryStorageCapacity request param
type GetPrimaryStorageCapacityParam struct {
	BaseParam
	GetPrimaryStorageCapacity GetPrimaryStorageCapacityParamDetail `json:"getPrimaryStorageCapacity"`
}
