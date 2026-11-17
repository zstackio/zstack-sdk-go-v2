// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateVpnIpsecConfig creates VpnIpsecConfig
func (cli *ZSClient) CreateVpnIpsecConfig(params param.CreateVpnIpsecConfigParam) (*view.VpcVpnIpSecConfigInventoryView, error) {
	resp := view.VpcVpnIpSecConfigInventoryView{}
	if err := cli.Post("v1/hybrid/vpn-connection/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcIpSecConfigFromLocal queries VpcIpSecConfigFromLocal list
func (cli *ZSClient) QueryVpcIpSecConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, error) {
	var resp []view.VpcVpnIpSecConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ipsec", params, &resp)
}

func (cli *ZSClient) GetVpcIpSecConfigFromLocal(uuid string) (*view.VpcVpnIpSecConfigInventoryView, error) {
	var resp view.VpcVpnIpSecConfigInventoryView
	if err := cli.Get("v1/hybrid/vpn-connection/ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcIpSecConfigFromLocal Pagination
func (cli *ZSClient) PageVpcIpSecConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIpSecConfigInventoryView, int, error) {
	var vpcIpSecConfigFromLocals []view.VpcVpnIpSecConfigInventoryView
	total, err := cli.Page("v1/hybrid/vpn-connection/ipsec", params, &vpcIpSecConfigFromLocals)
	return vpcIpSecConfigFromLocals, total, err
}

// DeleteVpcIpSecConfigLocal deletes VpcIpSecConfigLocal
func (cli *ZSClient) DeleteVpcIpSecConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ipsec", uuid, string(deleteMode))
}
