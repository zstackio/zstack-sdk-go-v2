// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBaremetalBonding queries BaremetalBonding list
func (cli *ZSClient) QueryBaremetalBonding(params param.QueryParam) ([]view.BaremetalBondingInventoryView, error) {
	var resp []view.BaremetalBondingInventoryView
	return resp, cli.List("v1/baremetal/network/bondings", &params, &resp)
}
