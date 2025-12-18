// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2VirtualIDAPIPermissionDetailParam GetIAM2VirtualIDPermission detail param
type GetIAM2VirtualIDAPIPermissionDetailParam struct {
	ApisToCheck []string `json:"apisToCheck,omitempty"`
	OnlyCheckParams bool `json:"onlyCheckParams,omitempty"`
}

// GetIAM2VirtualIDAPIPermissionParam GetIAM2VirtualIDPermission request param
type GetIAM2VirtualIDAPIPermissionParam struct {
	BaseParam
	Params GetIAM2VirtualIDAPIPermissionDetailParam `json:"params"`
}
