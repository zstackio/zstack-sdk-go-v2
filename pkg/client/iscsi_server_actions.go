// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddIscsiServer adds IscsiServer
func (cli *ZSClient) AddIscsiServer(params param.AddIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	var resp view.AddIscsiServerEventView
	if err := cli.Post("v1/storage-devices/iscsi/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIscsiServer queries IscsiServer list
func (cli *ZSClient) QueryIscsiServer(params *param.QueryParam) ([]view.IscsiServerInventoryView, error) {
	var resp []view.IscsiServerInventoryView
	return resp, cli.List("v1/storage-devices/iscsi/servers", params, &resp)
}
// DeleteIscsiServer deletes IscsiServer
func (cli *ZSClient) DeleteIscsiServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/storage-devices/iscsi/servers/{uuid}", uuid, string(deleteMode))
}
// RefreshIscsiServer operates on IscsiServer
func (cli *ZSClient) RefreshIscsiServer(params param.RefreshIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	var resp view.RefreshIscsiServerEventView
	if err := cli.Post("v1/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateIscsiServer updates IscsiServer
func (cli *ZSClient) UpdateIscsiServer(uuid string, params param.UpdateIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	var resp view.UpdateIscsiServerEventView
	if err := cli.Put("v1/storage-devices/iscsi/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
