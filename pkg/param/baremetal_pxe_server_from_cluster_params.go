// Copyright (c) ZStack.io, Inc.

package param

// DetachBaremetalPxeServerFromClusterDetailParam DetachBaremetalPxeServerFromCluster详细参数
type DetachBaremetalPxeServerFromClusterDetailParam struct {
	rest string `json:"pxeServerUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// DetachBaremetalPxeServerFromClusterParam DetachBaremetalPxeServerFromCluster请求参数
type DetachBaremetalPxeServerFromClusterParam struct {
	BaseParam
	Params DetachBaremetalPxeServerFromClusterDetailParam `json:"params"` // 详细参数
}

