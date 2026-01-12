// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryZone queries Zone list
func (cli *ZSClient) QueryZone(params *param.QueryParam) ([]view.ZoneInventoryView, error) {
	var resp []view.ZoneInventoryView
	return resp, cli.List("v1/zones", params, &resp)
}

func (cli *ZSClient) GetZone(uuid string) (*view.ZoneInventoryView, error) {
	var resp view.ZoneInventoryView
	if err := cli.Get("v1/zones", uuid, nil, &resp); err != nil {
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
	return cli.DeleteWithSpec("v1/zones", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateZone updates Zone
func (cli *ZSClient) UpdateZone(uuid string, params param.UpdateZoneParam) (*view.ZoneInventoryView, error) {
	var resp view.UpdateZoneEventView
	err := cli.PutWithSpec("v1/zones", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
