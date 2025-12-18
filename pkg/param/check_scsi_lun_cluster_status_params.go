// Copyright (c) ZStack.io, Inc.

package param

// CheckScsiLunClusterStatusDetailParam CheckScsiLunClusterStatus detail param
type CheckScsiLunClusterStatusDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// CheckScsiLunClusterStatusParam CheckScsiLunClusterStatus request param
type CheckScsiLunClusterStatusParam struct {
	BaseParam
	Params CheckScsiLunClusterStatusDetailParam `json:"params"`
}
