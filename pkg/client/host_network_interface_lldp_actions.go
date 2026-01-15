// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetHostNetworkInterfaceLldp gets HostNetworkInterfaceLldp by uuid
func (cli *ZSClient) GetHostNetworkInterfaceLldp(uuid string) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp view.HostNetworkInterfaceLldpInventoryView
	if err := cli.Get("v1/hostNetworkInterface/lldp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostNetworkInterfaceLldp queries HostNetworkInterfaceLldp list
func (cli *ZSClient) QueryHostNetworkInterfaceLldp(params *param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp []view.HostNetworkInterfaceLldpInventoryView
	return resp, cli.List("v1/hostNetworkInterface/lldp/all", params, &resp)
}

// PageHostNetworkInterfaceLldp Pagination
func (cli *ZSClient) PageHostNetworkInterfaceLldp(params *param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, int, error) {
	var hostNetworkInterfaceLldps []view.HostNetworkInterfaceLldpInventoryView
	total, err := cli.Page("v1/hostNetworkInterface/lldp/all", params, &hostNetworkInterfaceLldps)
	return hostNetworkInterfaceLldps, total, err
}
