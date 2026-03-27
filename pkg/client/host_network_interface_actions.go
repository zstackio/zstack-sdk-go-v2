// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostNetworkInterface updates HostNetworkInterface
func (cli *ZSClient) UpdateHostNetworkInterface(ctx context.Context, params param.UpdateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.Post(ctx, "v1/hosts/nics/{interfaceUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostNetworkInterface queries HostNetworkInterface list
func (cli *ZSClient) QueryHostNetworkInterface(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, error) {
	var resp []view.HostNetworkInterfaceInventoryView
	return resp, cli.List(ctx, "v1/hosts/nics", params, &resp)
}

func (cli *ZSClient) GetHostNetworkInterface(ctx context.Context, uuid string) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.HostNetworkInterfaceInventoryView
	if err := cli.Get(ctx, "v1/hosts/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostNetworkInterface Pagination
func (cli *ZSClient) PageHostNetworkInterface(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, int, error) {
	var hostNetworkInterfaces []view.HostNetworkInterfaceInventoryView
	total, err := cli.Page(ctx, "v1/hosts/nics", params, &hostNetworkInterfaces)
	return hostNetworkInterfaces, total, err
}
// LocateHostNetworkInterface operates on HostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(ctx context.Context, hostUuid string, params param.LocateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts", hostUuid, "", map[string]interface{}{
		"locateHostNetworkInterface": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
