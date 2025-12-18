// Copyright (c) ZStack.io, Inc.

package param

// DeleteVxlanPoolRemoteVtepDetailParam DeleteVxlanPoolRemoteVtep detail param
type DeleteVxlanPoolRemoteVtepDetailParam struct {
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	RemoteVtepIp string `json:"remoteVtepIp" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteVxlanPoolRemoteVtepParam DeleteVxlanPoolRemoteVtep request param
type DeleteVxlanPoolRemoteVtepParam struct {
	BaseParam
	Params DeleteVxlanPoolRemoteVtepDetailParam `json:"params"`
}
