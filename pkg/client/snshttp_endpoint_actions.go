// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSHttpEndpoint creates SNSHttpEndpoint
func (cli *ZSClient) CreateSNSHttpEndpoint(ctx context.Context, params param.CreateSNSHttpEndpointParam) (*view.SNSHttpEndpointInventoryView, error) {
	resp := view.SNSHttpEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/sns/application-endpoints/http", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSNSHttpEndpoint updates SNSHttpEndpoint
func (cli *ZSClient) UpdateSNSHttpEndpoint(ctx context.Context, uuid string, params param.UpdateSNSHttpEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/http", uuid, "", map[string]interface{}{
		"updateSNSHttpEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSHttpEndpoint queries SNSHttpEndpoint list
func (cli *ZSClient) QuerySNSHttpEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSHttpEndpointInventoryView, error) {
	var resp []view.SNSHttpEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/http", params, &resp)
}

func (cli *ZSClient) GetSNSHttpEndpoint(ctx context.Context, uuid string) (*view.SNSHttpEndpointInventoryView, error) {
	var resp view.SNSHttpEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/http", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSHttpEndpoint Pagination
func (cli *ZSClient) PageSNSHttpEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSHttpEndpointInventoryView, int, error) {
	var sNSHttpEndpoints []view.SNSHttpEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/http", params, &sNSHttpEndpoints)
	return sNSHttpEndpoints, total, err
}
