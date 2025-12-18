// Copyright (c) ZStack.io, Inc.

package param

// UpdateChronyServersDetailParam UpdateChronyServers详细参数
type UpdateChronyServersDetailParam struct {
	rest []string `json:"internalHostnames,omitempty"`
	rest []string `json:"externalHostnames,omitempty"`
}

// UpdateChronyServersParam UpdateChronyServers请求参数
type UpdateChronyServersParam struct {
	BaseParam
	Params UpdateChronyServersDetailParam `json:"params"` // 详细参数
}

