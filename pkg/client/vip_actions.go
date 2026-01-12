// Copyright (c) ZStack.io, Inc.

package client

import (
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
// DeleteVip deletes Vip
func (cli *ZSClient) DeleteVip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips", uuid, string(deleteMode))
}
// UpdateVip updates Vip
func (cli *ZSClient) UpdateVip(uuid string, params param.UpdateVipParam) (*view.VipInventoryView, error) {
	var resp view.UpdateVipEventView
	if err := cli.Put("v1/vips", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateVip creates Vip
func (cli *ZSClient) CreateVip(params param.CreateVipParam) (*view.VipInventoryView, error) {
	var resp view.CreateVipEventView
	if err := cli.Post("v1/vips", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
