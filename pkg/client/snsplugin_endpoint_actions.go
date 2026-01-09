// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSPluginEndpoint queries SNSPluginEndpoint list
func (cli *ZSClient) QuerySNSPluginEndpoint(params *param.QueryParam) ([]view.SNSPluginEndpointInventoryView, error) {
	var resp []view.SNSPluginEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/plugin", params, &resp)
}

func (cli *ZSClient) GetSNSPluginEndpoint(uuid string) (*view.SNSPluginEndpointInventoryView, error) {
	var resp view.SNSPluginEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/plugin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSPluginEndpoint creates SNSPluginEndpoint
func (cli *ZSClient) CreateSNSPluginEndpoint(params param.CreateSNSPluginEndpointParam) (*view.SNSPluginEndpointInventoryView, error) {
	var resp view.CreateSNSPluginEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
