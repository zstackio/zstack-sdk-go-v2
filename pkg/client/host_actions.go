// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectHost operates on Host
func (cli *ZSClient) ReconnectHost(uuid string, params param.ReconnectHostParam) (*view.HostInventoryView, error) {
	var resp view.ReconnectHostEventView
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateHost updates Host
func (cli *ZSClient) UpdateHost(uuid string, params param.UpdateHostParam) (*view.HostInventoryView, error) {
	var resp view.UpdateHostEventView
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteHost deletes Host
func (cli *ZSClient) DeleteHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/{uuid}", uuid, string(deleteMode))
}
// QueryHost queries Host list
func (cli *ZSClient) QueryHost(params *param.QueryParam) ([]view.HostInventoryView, error) {
	var resp []view.HostInventoryView
	return resp, cli.List("v1/hosts", params, &resp)
}
