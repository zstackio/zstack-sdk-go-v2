// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2VirtualIDAPIPermissionDetailParam GetIAM2VirtualIDPermission详细参数
type GetIAM2VirtualIDAPIPermissionDetailParam struct {
	rest []string `json:"apisToCheck,omitempty"`
	rest bool `json:"onlyCheckParams,omitempty"`
}

// GetIAM2VirtualIDAPIPermissionParam GetIAM2VirtualIDPermission请求参数
type GetIAM2VirtualIDAPIPermissionParam struct {
	BaseParam
	Params GetIAM2VirtualIDAPIPermissionDetailParam `json:"params"` // 详细参数
}

