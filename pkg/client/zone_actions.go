// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryZone queries Zone list
func (cli *ZSClient) QueryZone(params *param.QueryParam) ([]view.ZoneInventoryView, error) {
	var resp []view.ZoneInventoryView
	return resp, cli.List("v1/zones", params, &resp)
}
// GetZone gets Zone by uuid
func (cli *ZSClient) GetZone(uuid string) (*view.ZoneInventoryView, error) {
	var resp view.ZoneInventoryView
	if err := cli.Get("v1/zones/{uuid}/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateZone creates Zone
func (cli *ZSClient) CreateZone(params param.CreateZoneParam) (*view.ZoneInventoryView, error) {
	var resp view.CreateZoneEventView
	if err := cli.Post("v1/zones", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteZone deletes Zone
func (cli *ZSClient) DeleteZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{uuid}", uuid, string(deleteMode))
}
// UpdateZone updates Zone
func (cli *ZSClient) UpdateZone(uuid string, params param.UpdateZoneParam) (*view.ZoneInventoryView, error) {
	var resp view.UpdateZoneEventView
	if err := cli.Put("v1/zones/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
