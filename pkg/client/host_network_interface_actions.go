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
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.Put("v1/hosts/nics", interfaceUuid, map[string]interface{}{
		"params": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageHostNetworkInterface Pagination
func (cli *ZSClient) PageHostNetworkInterface(params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, int, error) {
	var hostNetworkInterfaces []view.HostNetworkInterfaceInventoryView
	total, err := cli.Page("v1/hosts/nics", params, &hostNetworkInterfaces)
	return hostNetworkInterfaces, total, err
}
// LocateHostNetworkInterface operates on HostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(hostUuid string, params param.LocateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.Put("v1/hosts", hostUuid, map[string]interface{}{
		"locateHostNetworkInterface": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
