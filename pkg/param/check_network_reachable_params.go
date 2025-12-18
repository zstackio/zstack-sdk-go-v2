// Copyright (c) ZStack.io, Inc.

package param

// CheckNetworkReachableDetailParam CheckNetworkReachable detail param
type CheckNetworkReachableDetailParam struct {
	SourceHostnames []string `json:"sourceHostnames,omitempty"`
	TargetHostnames []string `json:"targetHostnames" validate:"required"`
}

// CheckNetworkReachableParam CheckNetworkReachable request param
type CheckNetworkReachableParam struct {
	BaseParam
	Params CheckNetworkReachableDetailParam `json:"params"`
}
