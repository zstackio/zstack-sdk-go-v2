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

// PageZone Pagination
func (cli *ZSClient) PageZone(params *param.QueryParam) ([]view.ZoneInventoryView, int, error) {
	var zones []view.ZoneInventoryView
	total, err := cli.Page("v1/zones", params, &zones)
	return zones, total, err
}
// CreateZone creates Zone
func (cli *ZSClient) CreateZone(params param.CreateZoneParam) (*view.ZoneInventoryView, error) {
	resp := view.ZoneInventoryView{}
	if err := cli.Post("v1/zones", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteZone deletes Zone
func (cli *ZSClient) DeleteZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones", uuid, string(deleteMode))
}
// UpdateZone updates Zone
func (cli *ZSClient) UpdateZone(uuid string, params param.UpdateZoneParam) (*view.ZoneInventoryView, error) {
	resp := view.ZoneInventoryView{}
	if err := cli.Put("v1/zones", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
