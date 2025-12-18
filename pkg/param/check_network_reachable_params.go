// Copyright (c) ZStack.io, Inc.

package param

// CheckNetworkReachableDetailParam CheckNetworkReachable详细参数
type CheckNetworkReachableDetailParam struct {
	rest []string `json:"sourceHostnames,omitempty"`
	rest []string `json:"targetHostnames" validate:"required"` // 必填
}

// CheckNetworkReachableParam CheckNetworkReachable请求参数
type CheckNetworkReachableParam struct {
	BaseParam
	Params CheckNetworkReachableDetailParam `json:"params"` // 详细参数
}

