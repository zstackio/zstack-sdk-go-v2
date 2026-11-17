// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVpcUserVpnGateway updates VpcUserVpnGateway
func (cli *ZSClient) UpdateVpcUserVpnGateway(uuid string, params param.UpdateVpcUserVpnGatewayParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/user-vpn", uuid, "", map[string]interface{}{
		"updateVpcUserVpnGateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcUserVpnGatewayLocal deletes VpcUserVpnGatewayLocal
func (cli *ZSClient) DeleteVpcUserVpnGatewayLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway", uuid, string(deleteMode))
}

// SyncVpcUserVpnGatewayFromRemote operates on VpcUserVpnGatewayFromRemote
func (cli *ZSClient) SyncVpcUserVpnGatewayFromRemote(dataCenterUuid string, params param.SyncVpcUserVpnGatewayFromRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/user-vpn", dataCenterUuid, "", map[string]interface{}{
		"syncVpcUserVpnGatewayFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcUserVpnGatewayRemote deletes VpcUserVpnGatewayRemote
func (cli *ZSClient) DeleteVpcUserVpnGatewayRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/user-gateway", uuid, string(deleteMode))
}

// QueryVpcUserVpnGatewayFromLocal queries VpcUserVpnGatewayFromLocal list
func (cli *ZSClient) QueryVpcUserVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, error) {
	var resp []view.VpcUserVpnGatewayInventoryView
	return resp, cli.List("v1/hybrid/user-vpn", params, &resp)
}

func (cli *ZSClient) GetVpcUserVpnGatewayFromLocal(uuid string) (*view.VpcUserVpnGatewayInventoryView, error) {
	var resp view.VpcUserVpnGatewayInventoryView
	if err := cli.Get("v1/hybrid/user-vpn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcUserVpnGatewayFromLocal Pagination
func (cli *ZSClient) PageVpcUserVpnGatewayFromLocal(params *param.QueryParam) ([]view.VpcUserVpnGatewayInventoryView, int, error) {
	var vpcUserVpnGatewayFromLocals []view.VpcUserVpnGatewayInventoryView
	total, err := cli.Page("v1/hybrid/user-vpn", params, &vpcUserVpnGatewayFromLocals)
	return vpcUserVpnGatewayFromLocals, total, err
}

// CreateVpcUserVpnGatewayRemote creates VpcUserVpnGatewayRemote
func (cli *ZSClient) CreateVpcUserVpnGatewayRemote(params param.CreateVpcUserVpnGatewayRemoteParam) (*view.VpcUserVpnGatewayInventoryView, error) {
	resp := view.VpcUserVpnGatewayInventoryView{}
	if err := cli.Post("v1/hybrid/user-vpn", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
