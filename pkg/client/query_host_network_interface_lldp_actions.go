// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHostNetworkInterfaceLldp queries HostNetworkInterfaceLldp list
func (cli *ZSClient) QueryHostNetworkInterfaceLldp(params *param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp []view.HostNetworkInterfaceLldpInventoryView
	return resp, cli.List("v1/hostNetworkInterface/lldp/all", params, &resp)
}
