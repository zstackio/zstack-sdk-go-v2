// Copyright (c) ZStack.io, Inc.

package param

// AttachBaremetalPxeServerToClusterDetailParam AttachBaremetalPxeServerToCluster detail param
type AttachBaremetalPxeServerToClusterDetailParam struct {
	PxeServerUuid string `json:"pxeServerUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// AttachBaremetalPxeServerToClusterParam AttachBaremetalPxeServerToCluster request param
type AttachBaremetalPxeServerToClusterParam struct {
	BaseParam
	Params AttachBaremetalPxeServerToClusterDetailParam `json:"params"`
}
