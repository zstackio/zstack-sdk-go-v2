// Copyright (c) ZStack.io, Inc.

package param

// CreateVxlanPoolRemoteVtepDetailParam CreateVxlanPoolRemoteVtep详细参数
type CreateVxlanPoolRemoteVtepDetailParam struct {
	rest string `json:"l2NetworkUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"remoteVtepIp" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateVxlanPoolRemoteVtepParam CreateVxlanPoolRemoteVtep请求参数
type CreateVxlanPoolRemoteVtepParam struct {
	BaseParam
	Params CreateVxlanPoolRemoteVtepDetailParam `json:"params"` // 详细参数
}

