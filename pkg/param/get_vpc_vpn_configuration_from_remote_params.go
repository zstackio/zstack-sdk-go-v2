// Copyright (c) ZStack.io, Inc.

package param

// GetVpcVpnConfigurationFromRemoteDetailParam GetVpcVpnConfigurationFromRemote detail param
type GetVpcVpnConfigurationFromRemoteDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVpcVpnConfigurationFromRemoteParam GetVpcVpnConfigurationFromRemote request param
type GetVpcVpnConfigurationFromRemoteParam struct {
	BaseParam
	Params GetVpcVpnConfigurationFromRemoteDetailParam `json:"params"`
}
