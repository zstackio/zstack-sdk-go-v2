// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteVpcIkeConfigLocal deletes VpcIkeConfigLocal
func (cli *ZSClient) DeleteVpcIkeConfigLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection/ike", uuid, string(deleteMode))
}

// CreateVpnIkeConfig creates VpnIkeConfig
func (cli *ZSClient) CreateVpnIkeConfig(params param.CreateVpnIkeConfigParam) (*view.VpcVpnIkeConfigInventoryView, error) {
	resp := view.VpcVpnIkeConfigInventoryView{}
	if err := cli.Post("v1/hybrid/vpn-connection/ike", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcIkeConfigFromLocal queries VpcIkeConfigFromLocal list
func (cli *ZSClient) QueryVpcIkeConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, error) {
	var resp []view.VpcVpnIkeConfigInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection/ike", params, &resp)
}

func (cli *ZSClient) GetVpcIkeConfigFromLocal(uuid string) (*view.VpcVpnIkeConfigInventoryView, error) {
	var resp view.VpcVpnIkeConfigInventoryView
	if err := cli.Get("v1/hybrid/vpn-connection/ike", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcIkeConfigFromLocal Pagination
func (cli *ZSClient) PageVpcIkeConfigFromLocal(params *param.QueryParam) ([]view.VpcVpnIkeConfigInventoryView, int, error) {
	var vpcIkeConfigFromLocals []view.VpcVpnIkeConfigInventoryView
	total, err := cli.Page("v1/hybrid/vpn-connection/ike", params, &vpcIkeConfigFromLocals)
	return vpcIkeConfigFromLocals, total, err
}
