// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2PlatformDetailParam LoginIAM2Platform详细参数
type LoginIAM2PlatformDetailParam struct {
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2PlatformParam LoginIAM2Platform请求参数
type LoginIAM2PlatformParam struct {
	BaseParam
	Params LoginIAM2PlatformDetailParam `json:"params"` // 详细参数
}

