// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostNetworkInterface updates HostNetworkInterface
func (cli *ZSClient) UpdateHostNetworkInterface(uuid string, params param.UpdateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.UpdateHostNetworkInterfaceEventView
	if err := cli.Put("v1/hosts/nics/{interfaceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryHostNetworkInterface queries HostNetworkInterface list
func (cli *ZSClient) QueryHostNetworkInterface(params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, error) {
	var resp []view.HostNetworkInterfaceInventoryView
	return resp, cli.List("v1/hosts/nics", params, &resp)
}
// LocateHostNetworkInterface operates on HostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(uuid string, params param.LocateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.Put("v1/hosts/{hostUuid}/locate/network-interface", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
