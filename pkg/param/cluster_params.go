// Copyright (c) ZStack.io, Inc.

package param

// UpdateClusterDetailParam UpdateCluster详细参数
type UpdateClusterDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
}

// UpdateClusterParam UpdateCluster请求参数
type UpdateClusterParam struct {
	BaseParam
	Params UpdateClusterDetailParam `json:"params"` // 详细参数
}

