// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNetworkServiceProvider queries NetworkServiceProvider list
func (cli *ZSClient) QueryNetworkServiceProvider(params *param.QueryParam) ([]view.NetworkServiceProviderInventoryView, error) {
	var resp []view.NetworkServiceProviderInventoryView
	return resp, cli.List("v1/network-services/providers", params, &resp)
}
