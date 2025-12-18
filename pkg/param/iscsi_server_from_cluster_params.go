// Copyright (c) ZStack.io, Inc.

package param

// DetachIscsiServerFromClusterDetailParam DetachIscsiServerFromCluster详细参数
type DetachIscsiServerFromClusterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// DetachIscsiServerFromClusterParam DetachIscsiServerFromCluster请求参数
type DetachIscsiServerFromClusterParam struct {
	BaseParam
	Params DetachIscsiServerFromClusterDetailParam `json:"params"` // 详细参数
}

