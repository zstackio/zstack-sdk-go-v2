// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryHostNetworkInterfaceLldp queries HostNetworkInterfaceLldp list
func (cli *ZSClient) QueryHostNetworkInterfaceLldp(params param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp []view.HostNetworkInterfaceLldpInventoryView
	return resp, cli.List("v1/hostNetworkInterface/lldp/all", &params, &resp)
}
