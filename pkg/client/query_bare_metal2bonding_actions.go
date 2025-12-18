// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2Bonding queries BareMetal2Bonding list
func (cli *ZSClient) QueryBareMetal2Bonding(params param.QueryParam) ([]view.BareMetal2BondingInventoryView, error) {
	var resp []view.BareMetal2BondingInventoryView
	return resp, cli.List("v1/baremetal2/bonding", &params, &resp)
}
