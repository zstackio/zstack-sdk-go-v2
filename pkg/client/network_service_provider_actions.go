// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNetworkServiceProvider queries NetworkServiceProvider list
func (cli *ZSClient) QueryNetworkServiceProvider(params *param.QueryParam) ([]view.NetworkServiceProviderInventoryView, error) {
	var resp []view.NetworkServiceProviderInventoryView
	return resp, cli.List("v1/network-services/providers", params, &resp)
}

func (cli *ZSClient) GetNetworkServiceProvider(uuid string) (*view.NetworkServiceProviderInventoryView, error) {
	var resp view.NetworkServiceProviderInventoryView
	if err := cli.Get("v1/network-services/providers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNetworkServiceProvider Pagination
func (cli *ZSClient) PageNetworkServiceProvider(params *param.QueryParam) ([]view.NetworkServiceProviderInventoryView, int, error) {
	var networkServiceProviders []view.NetworkServiceProviderInventoryView
	total, err := cli.Page("v1/network-services/providers", params, &networkServiceProviders)
	return networkServiceProviders, total, err
}
