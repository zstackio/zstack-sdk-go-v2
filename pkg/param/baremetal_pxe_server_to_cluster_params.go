// Copyright (c) ZStack.io, Inc.

package param

// AttachBaremetalPxeServerToClusterDetailParam AttachBaremetalPxeServerToCluster详细参数
type AttachBaremetalPxeServerToClusterDetailParam struct {
	rest string `json:"pxeServerUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// AttachBaremetalPxeServerToClusterParam AttachBaremetalPxeServerToCluster请求参数
type AttachBaremetalPxeServerToClusterParam struct {
	BaseParam
	Params AttachBaremetalPxeServerToClusterDetailParam `json:"params"` // 详细参数
}

