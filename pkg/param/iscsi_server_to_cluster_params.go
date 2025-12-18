// Copyright (c) ZStack.io, Inc.

package param

// AttachIscsiServerToClusterDetailParam AttachIscsiServerToCluster详细参数
type AttachIscsiServerToClusterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"clusterUuid" validate:"required"` // 必填
}

// AttachIscsiServerToClusterParam AttachIscsiServerToCluster请求参数
type AttachIscsiServerToClusterParam struct {
	BaseParam
	Params AttachIscsiServerToClusterDetailParam `json:"params"` // 详细参数
}

