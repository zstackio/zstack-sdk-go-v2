// Copyright (c) ZStack.io, Inc.

package param

// CreateVxlanPoolRemoteVtepDetailParam CreateVxlanPoolRemoteVtep detail param
type CreateVxlanPoolRemoteVtepDetailParam struct {
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	RemoteVtepIp string `json:"remoteVtepIp" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVxlanPoolRemoteVtepParam CreateVxlanPoolRemoteVtep request param
type CreateVxlanPoolRemoteVtepParam struct {
	BaseParam
	Params CreateVxlanPoolRemoteVtepDetailParam `json:"params"`
}
