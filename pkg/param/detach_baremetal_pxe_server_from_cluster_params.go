// Copyright (c) ZStack.io, Inc.

package param

// DetachBaremetalPxeServerFromClusterDetailParam DetachBaremetalPxeServerFromCluster detail param
type DetachBaremetalPxeServerFromClusterDetailParam struct {
	PxeServerUuid string `json:"pxeServerUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// DetachBaremetalPxeServerFromClusterParam DetachBaremetalPxeServerFromCluster request param
type DetachBaremetalPxeServerFromClusterParam struct {
	BaseParam
	Params DetachBaremetalPxeServerFromClusterDetailParam `json:"params"`
}
