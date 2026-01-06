// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePortMirror creates PortMirror
func (cli *ZSClient) CreatePortMirror(params param.CreatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	var resp view.CreatePortMirrorEventView
	if err := cli.Post("v1/port-mirrors", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryPortMirror queries PortMirror list
func (cli *ZSClient) QueryPortMirror(params *param.QueryParam) ([]view.PortMirrorInventoryView, error) {
	var resp []view.PortMirrorInventoryView
	return resp, cli.List("v1/port-mirrors", params, &resp)
}
// DeletePortMirror deletes PortMirror
func (cli *ZSClient) DeletePortMirror(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/port-mirrors/{uuid}", uuid, string(deleteMode))
}
// UpdatePortMirror updates PortMirror
func (cli *ZSClient) UpdatePortMirror(uuid string, params param.UpdatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	var resp view.UpdatePortMirrorEventView
	if err := cli.Put("v1/port-mirrors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
