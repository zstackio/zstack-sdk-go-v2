// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVpnConfigurationFromRemoteDetailParam GetVpcVpnConfigurationFromRemote详细参数
type GetVpcVpnConfigurationFromRemoteDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVpcVpnConfigurationFromRemoteParam GetVpcVpnConfigurationFromRemote请求参数
type GetVpcVpnConfigurationFromRemoteParam struct {
	BaseParam
	Params GetVpcVpnConfigurationFromRemoteDetailParam `json:"params"` // 详细参数
}

