// Copyright (c) ZStack.io, Inc.

package view

// GetVpcVpnConfigurationFromRemoteView GetVpcVpnConfigurationFromRemote
type GetVpcVpnConfigurationFromRemoteView struct {
	IkeConf VpcVpnIkeConfigStructView `json:"ikeConf,omitempty"`
	IpSecConf VpcVpnIpSecConfigStructView `json:"ipSecConf,omitempty"`
}

