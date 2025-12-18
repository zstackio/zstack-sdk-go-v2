// Copyright (c) ZStack.io, Inc.

package param

// GetPrimaryStorageCapacityDetailParam GetPrimaryStorageCapacity detail param
type GetPrimaryStorageCapacityDetailParam struct {
	ZoneUuids []string `json:"zoneUuids,omitempty"`
	ClusterUuids []string `json:"clusterUuids,omitempty"`
	PrimaryStorageUuids []string `json:"primaryStorageUuids,omitempty"`
	All bool `json:"all,omitempty"`
}

// GetPrimaryStorageCapacityParam GetPrimaryStorageCapacity request param
type GetPrimaryStorageCapacityParam struct {
	BaseParam
	Params GetPrimaryStorageCapacityDetailParam `json:"params"`
}
