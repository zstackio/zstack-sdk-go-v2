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

func (cli *ZSClient) GetHostNetworkBonding(uuid string) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.HostNetworkBondingInventoryView
	if err := cli.Get("v1/hosts/bondings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
