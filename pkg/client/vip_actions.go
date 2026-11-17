// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVip queries Vip list
func (cli *ZSClient) QueryVip(params *param.QueryParam) ([]view.VipInventoryView, error) {
	var resp []view.VipInventoryView
	return resp, cli.List("v1/vips", params, &resp)
}

func (cli *ZSClient) GetVip(uuid string) (*view.VipInventoryView, error) {
	var resp view.VipInventoryView
	if err := cli.Get("v1/vips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVip Pagination
func (cli *ZSClient) PageVip(params *param.QueryParam) ([]view.VipInventoryView, int, error) {
	var vips []view.VipInventoryView
	total, err := cli.Page("v1/vips", params, &vips)
	return vips, total, err
}
// DeleteVip deletes Vip
func (cli *ZSClient) DeleteVip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips", uuid, string(deleteMode))
}
// UpdateVip updates Vip
func (cli *ZSClient) UpdateVip(uuid string, params param.UpdateVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PutWithRespKey("v1/vips", uuid, "", map[string]interface{}{
		"updateVip": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVip creates Vip
func (cli *ZSClient) CreateVip(params param.CreateVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.Post("v1/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVipPortAvailability operates on VipPortAvailability
func (cli *ZSClient) CheckVipPortAvailability(uuid string) (*view.CheckVipPortAvailabilityView, error) {
	var resp view.CheckVipPortAvailabilityView
	if err := cli.GetWithRespKey("v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVipState changes VipState
func (cli *ZSClient) ChangeVipState(uuid string, params param.ChangeVipStateParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PutWithRespKey("v1/vips", uuid, "", map[string]interface{}{
		"changeVipState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedVip gets VpcAttachedVip by uuid
func (cli *ZSClient) GetVpcAttachedVip(uuid string, params param.GetVpcAttachedVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-vip", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVipUsedPorts gets VipUsedPorts by uuid
func (cli *ZSClient) GetVipUsedPorts(uuid string) (*view.VipPortRangeInventoryView, error) {
	var resp view.VipPortRangeInventoryView
	if err := cli.GetWithRespKey("v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVipAvailablePort gets VipAvailablePort by uuid
func (cli *ZSClient) GetVipAvailablePort(uuid string) (*view.GetVipAvailablePortView, error) {
	var resp view.GetVipAvailablePortView
	if err := cli.GetWithRespKey("v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
