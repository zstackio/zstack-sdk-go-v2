// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVip queries Vip list
func (cli *ZSClient) QueryVip(params *param.QueryParam) ([]view.VipInventoryView, error) {
	var resp []view.VipInventoryView
	return resp, cli.List("v1/vips", params, &resp)
}
// DeleteVip deletes Vip
func (cli *ZSClient) DeleteVip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/{uuid}", uuid, string(deleteMode))
}
// UpdateVip updates Vip
func (cli *ZSClient) UpdateVip(uuid string, params param.UpdateVipParam) (*view.VipInventoryView, error) {
	var resp view.UpdateVipEventView
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
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
