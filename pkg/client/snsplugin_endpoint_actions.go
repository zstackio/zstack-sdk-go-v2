// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSPluginEndpoint queries SNSPluginEndpoint list
func (cli *ZSClient) QuerySNSPluginEndpoint(params *param.QueryParam) ([]view.SNSPluginEndpointInventoryView, error) {
	var resp []view.SNSPluginEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/plugin", params, &resp)
}
// CreateSNSPluginEndpoint creates SNSPluginEndpoint
func (cli *ZSClient) CreateSNSPluginEndpoint(params param.CreateSNSPluginEndpointParam) (*view.SNSPluginEndpointInventoryView, error) {
	var resp view.CreateSNSPluginEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
