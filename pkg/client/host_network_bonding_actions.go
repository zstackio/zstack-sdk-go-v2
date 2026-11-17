// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHostNetworkBonding queries HostNetworkBonding list
func (cli *ZSClient) QueryHostNetworkBonding(params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, error) {
	var resp []view.HostNetworkBondingInventoryView
	return resp, cli.List("v1/hosts/bondings", params, &resp)
}

func (cli *ZSClient) GetHostNetworkBonding(uuid string) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.HostNetworkBondingInventoryView
	if err := cli.Get("v1/hosts/bondings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostNetworkBonding Pagination
func (cli *ZSClient) PageHostNetworkBonding(params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, int, error) {
	var hostNetworkBondings []view.HostNetworkBondingInventoryView
	total, err := cli.Page("v1/hosts/bondings", params, &hostNetworkBondings)
	return hostNetworkBondings, total, err
}

// SetServiceTypeOnHostNetworkBonding operates on ServiceTypeOnHostNetworkBonding
func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(params param.SetServiceTypeOnHostNetworkBondingParam) (*view.HostNetworkBondingServiceRefInventoryView, error) {
	resp := view.HostNetworkBondingServiceRefInventoryView{}
	if err := cli.Post("v1/hosts/bondings/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachNicToBonding operates on NicToBonding
func (cli *ZSClient) AttachNicToBonding(uuid string, params param.AttachNicToBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/bondings", uuid, "", map[string]interface{}{
		"attachNicToBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachNicFromBonding operates on NicFromBonding
func (cli *ZSClient) DetachNicFromBonding(uuid string, params param.DetachNicFromBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/bondings", uuid, "", map[string]interface{}{
		"detachNicFromBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBonding updates Bonding
func (cli *ZSClient) UpdateBonding(uuid string, params param.UpdateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/bondings", uuid, "", map[string]interface{}{
		"updateBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIpOnHostNetworkBonding operates on IpOnHostNetworkBonding
func (cli *ZSClient) SetIpOnHostNetworkBonding(bondingUuid string, params param.SetIpOnHostNetworkBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hosts/bondings/%s/ip", bondingUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBonding creates Bonding
func (cli *ZSClient) CreateBonding(params param.CreateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.Post("v1/hosts/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBonding deletes Bonding
func (cli *ZSClient) DeleteBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/bondings", uuid, string(deleteMode))
}
