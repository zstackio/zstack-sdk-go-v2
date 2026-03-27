// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedResource queries SharedResource list
func (cli *ZSClient) QuerySharedResource(ctx context.Context, params *param.QueryParam) ([]view.SharedResourceInventoryView, error) {
	var resp []view.SharedResourceInventoryView
	return resp, cli.List(ctx, "v1/accounts/resources", params, &resp)
}

func (cli *ZSClient) GetSharedResource(ctx context.Context, uuid string) (*view.SharedResourceInventoryView, error) {
	var resp view.SharedResourceInventoryView
	if err := cli.Get(ctx, "v1/accounts/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSharedResource Pagination
func (cli *ZSClient) PageSharedResource(ctx context.Context, params *param.QueryParam) ([]view.SharedResourceInventoryView, int, error) {
	var sharedResources []view.SharedResourceInventoryView
	total, err := cli.Page(ctx, "v1/accounts/resources", params, &sharedResources)
	return sharedResources, total, err
}
