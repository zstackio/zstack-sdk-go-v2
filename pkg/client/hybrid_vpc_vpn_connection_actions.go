// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetVpcVpnConfigurationFromRemote gets VpcVpnConfigurationFromRemote by uuid
func (cli *ZSClient) GetVpcVpnConfigurationFromRemote(uuid string) (*view.GetVpcVpnConfigurationFromRemoteView, error) {
	var resp view.GetVpcVpnConfigurationFromRemoteView
	if err := cli.GetWithRespKey("v1/hybrid/vpn-conf", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVpcVpnConnectionFromRemote operates on VpcVpnConnectionFromRemote
func (cli *ZSClient) SyncVpcVpnConnectionFromRemote(dataCenterUuid string, params param.SyncVpcVpnConnectionFromRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/vpn-connection", dataCenterUuid, "", map[string]interface{}{
		"syncVpcVpnConnectionFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVpnConnectionRemote creates VpcVpnConnectionRemote
func (cli *ZSClient) CreateVpcVpnConnectionRemote(params param.CreateVpcVpnConnectionRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.Post("v1/hybrid/vpn-connection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnConnectionRemote deletes VpcVpnConnectionRemote
func (cli *ZSClient) DeleteVpcVpnConnectionRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection", uuid, string(deleteMode))
}

// QueryVpcVpnConnectionFromLocal queries VpcVpnConnectionFromLocal list
func (cli *ZSClient) QueryVpcVpnConnectionFromLocal(params *param.QueryParam) ([]view.VpcVpnConnectionInventoryView, error) {
	var resp []view.VpcVpnConnectionInventoryView
	return resp, cli.List("v1/hybrid/vpn-connection", params, &resp)
}

func (cli *ZSClient) GetVpcVpnConnectionFromLocal(uuid string) (*view.VpcVpnConnectionInventoryView, error) {
	var resp view.VpcVpnConnectionInventoryView
	if err := cli.Get("v1/hybrid/vpn-connection", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcVpnConnectionFromLocal Pagination
func (cli *ZSClient) PageVpcVpnConnectionFromLocal(params *param.QueryParam) ([]view.VpcVpnConnectionInventoryView, int, error) {
	var vpcVpnConnectionFromLocals []view.VpcVpnConnectionInventoryView
	total, err := cli.Page("v1/hybrid/vpn-connection", params, &vpcVpnConnectionFromLocals)
	return vpcVpnConnectionFromLocals, total, err
}

// UpdateVpcVpnConnectionRemote updates VpcVpnConnectionRemote
func (cli *ZSClient) UpdateVpcVpnConnectionRemote(uuid string, params param.UpdateVpcVpnConnectionRemoteParam) (*view.VpcVpnConnectionInventoryView, error) {
	resp := view.VpcVpnConnectionInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/vpn-connection", uuid, "", map[string]interface{}{
		"updateVpcVpnConnectionRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpcVpnConnectionLocal deletes VpcVpnConnectionLocal
func (cli *ZSClient) DeleteVpcVpnConnectionLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/vpn-connection", uuid, string(deleteMode))
}
