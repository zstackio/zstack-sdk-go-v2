// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2ProvisionNetwork queries BareMetal2ProvisionNetwork list
func (cli *ZSClient) QueryBareMetal2ProvisionNetwork(params param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp []view.BareMetal2ProvisionNetworkInventoryView
	return resp, cli.List("v1/baremetal2/provision-networks", &params, &resp)
}
