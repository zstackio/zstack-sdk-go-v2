// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcVpnGateway updates VpcVpnGateway
func (cli *ZSClient) UpdateVpcVpnGateway(uuid string, params param.UpdateVpcVpnGatewayParam) (*view.VpcVpnGatewayInventoryView, error) {
	resp := view.VpcVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/vpc-vpn", uuid, "", map[string]interface{}{
		"updateVpcVpnGateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnGatewayLocal deletes VpcVpnGatewayLocal
func (cli *ZSClient) DeleteVpcVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-gateway", uuid, string(deleteMode))
}

// QueryVpcVpnGatewayFromLocal queries VpcVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcVpnGatewayInventoryView, error) {
	var resp []view.VpcVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/vpc-vpn", params, &resp)
}

func (cli *ZSClient) GetVpcVpnGatewayFromLocal(uuid string) (*view.VpcVpnGatewayInventoryView, error) {
	var resp view.VpcVpnGatewayInventoryView
	if err := cli.Get("v1/hybrid/vpc-vpn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcVpnGatewayFromLocal Pagination
func (cli *ZSClient) PageVpcVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcVpnGatewayInventoryView, int, error) {
	var vpcVpnGatewayFromLocals []view.VpcVpnGatewayInventoryView
	total, err := cli.Page("v1/hybrid/vpc-vpn", params, &vpcVpnGatewayFromLocals)
	return vpcVpnGatewayFromLocals, total, err
}

// SyncVpcVpnGatewayFromRemote operates on VpcVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcVpnGatewayFromRemote(dataCenterUuid string, params param.SyncVpcVpnGatewayFromRemoteParam) (*view.VpcVpnGatewayInventoryView, error) {
	resp := view.VpcVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/vpc-vpn", dataCenterUuid, "", map[string]interface{}{
		"syncVpcVpnGatewayFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
