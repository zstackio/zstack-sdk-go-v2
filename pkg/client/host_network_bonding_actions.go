// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHostNetworkBonding queries HostNetworkBonding list
func (cli *ZSClient) QueryHostNetworkBonding(params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, error) {
	var resp []view.HostNetworkBondingInventoryView
	return resp, cli.List("v1/hosts/bondings", params, &resp)
}

// PageHostNetworkBonding Pagination
func (cli *ZSClient) PageHostNetworkBonding(params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, int, error) {
	var hostNetworkBondings []view.HostNetworkBondingInventoryView
	total, err := cli.Page("v1/hosts/bondings", params, &hostNetworkBondings)
	return hostNetworkBondings, total, err
}
