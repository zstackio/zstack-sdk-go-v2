// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSPluginEndpoint queries SNSPluginEndpoint list
func (cli *ZSClient) QuerySNSPluginEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSPluginEndpointInventoryView, error) {
	var resp []view.SNSPluginEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/plugin", params, &resp)
}

func (cli *ZSClient) GetSNSPluginEndpoint(ctx context.Context, uuid string) (*view.SNSPluginEndpointInventoryView, error) {
	var resp view.SNSPluginEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/plugin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSPluginEndpoint Pagination
func (cli *ZSClient) PageSNSPluginEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSPluginEndpointInventoryView, int, error) {
	var sNSPluginEndpoints []view.SNSPluginEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/plugin", params, &sNSPluginEndpoints)
	return sNSPluginEndpoints, total, err
}
// CreateSNSPluginEndpoint creates SNSPluginEndpoint
func (cli *ZSClient) CreateSNSPluginEndpoint(ctx context.Context, params param.CreateSNSPluginEndpointParam) (*view.SNSPluginEndpointInventoryView, error) {
	resp := view.SNSPluginEndpointInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/plugin", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
