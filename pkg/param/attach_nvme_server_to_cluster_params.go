// Copyright (c) ZStack.io, Inc.

package param

// AttachNvmeServerToClusterDetailParam AttachNvmeServerToCluster detail param
type AttachNvmeServerToClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// AttachNvmeServerToClusterParam AttachNvmeServerToCluster request param
type AttachNvmeServerToClusterParam struct {
	BaseParam
	Params AttachNvmeServerToClusterDetailParam `json:"params"`
}
