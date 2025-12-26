// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryHostNetworkBonding queries HostNetworkBonding list
func (cli *ZSClient) QueryHostNetworkBonding(params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, error) {
	var resp []view.HostNetworkBondingInventoryView
	return resp, cli.List("v1/hosts/bondings", params, &resp)
}
