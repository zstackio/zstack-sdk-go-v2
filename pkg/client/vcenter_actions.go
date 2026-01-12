// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddVCenter adds VCenter
func (cli *ZSClient) AddVCenter(params param.AddVCenterParam) (*view.VCenterInventoryView, error) {
	var resp view.AddVCenterEventView
	if err := cli.Post("v1/vcenters", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// SyncVCenter operates on VCenter
func (cli *ZSClient) SyncVCenter(uuid string, params param.SyncVCenterParam) (*view.VCenterInventoryView, error) {
	resp := view.VCenterInventoryView{}
	err := cli.PutWithSpec("v1/vcenters", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVCenter queries VCenter list
func (cli *ZSClient) QueryVCenter(params *param.QueryParam) ([]view.VCenterInventoryView, error) {
	var resp []view.VCenterInventoryView
	return resp, cli.List("v1/vcenters", params, &resp)
}

func (cli *ZSClient) GetVCenter(uuid string) (*view.VCenterInventoryView, error) {
	var resp view.VCenterInventoryView
	if err := cli.Get("v1/vcenters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVCenter updates VCenter
func (cli *ZSClient) UpdateVCenter(uuid string, params param.UpdateVCenterParam) (*view.VCenterInventoryView, error) {
	var resp view.UpdateVCenterEventView
	err := cli.PutWithSpec("v1/vcenters", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVCenter deletes VCenter
func (cli *ZSClient) DeleteVCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vcenters", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
