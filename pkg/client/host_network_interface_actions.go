// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostNetworkInterface updates HostNetworkInterface
func (cli *ZSClient) UpdateHostNetworkInterface(interfaceUuid string, params param.UpdateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.UpdateHostNetworkInterfaceEventView
	err := cli.PutWithSpec("v1/hosts/nics", fmt.Sprintf(\"%s/actions\", interfaceUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryHostNetworkInterface queries HostNetworkInterface list
func (cli *ZSClient) QueryHostNetworkInterface(params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, error) {
	var resp []view.HostNetworkInterfaceInventoryView
	return resp, cli.List("v1/hosts/nics", params, &resp)
}

func (cli *ZSClient) GetHostNetworkInterface(uuid string) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.HostNetworkInterfaceInventoryView
	if err := cli.Get("v1/hosts/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// LocateHostNetworkInterface operates on HostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(hostUuid string, params param.LocateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	err := cli.PutWithSpec("v1/hosts", fmt.Sprintf(\"%s/locate/network-interface\", hostUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
