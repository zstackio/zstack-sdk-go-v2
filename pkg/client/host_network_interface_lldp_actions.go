// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetHostNetworkInterfaceLldp gets HostNetworkInterfaceLldp by uuid
func (cli *ZSClient) GetHostNetworkInterfaceLldp(uuid string) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp view.HostNetworkInterfaceLldpInventoryView
	if err := cli.Get("v1/hostNetworkInterface/lldp/{interfaceUuid}/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostNetworkInterfaceLldp queries HostNetworkInterfaceLldp list
func (cli *ZSClient) QueryHostNetworkInterfaceLldp(params *param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp []view.HostNetworkInterfaceLldpInventoryView
	return resp, cli.List("v1/hostNetworkInterface/lldp/all", params, &resp)
}
