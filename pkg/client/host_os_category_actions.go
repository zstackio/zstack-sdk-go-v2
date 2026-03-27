// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryHostOsCategory queries HostOsCategory list
func (cli *ZSClient) QueryHostOsCategory(ctx context.Context, params *param.QueryParam) ([]view.HostOsCategoryInventoryView, error) {
	var resp []view.HostOsCategoryInventoryView
	return resp, cli.List(ctx, "v1/hosts/os/category", params, &resp)
}

func (cli *ZSClient) GetHostOsCategory(ctx context.Context, uuid string) (*view.HostOsCategoryInventoryView, error) {
	var resp view.HostOsCategoryInventoryView
	if err := cli.Get(ctx, "v1/hosts/os/category", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostOsCategory Pagination
func (cli *ZSClient) PageHostOsCategory(ctx context.Context, params *param.QueryParam) ([]view.HostOsCategoryInventoryView, int, error) {
	var hostOsCategories []view.HostOsCategoryInventoryView
	total, err := cli.Page(ctx, "v1/hosts/os/category", params, &hostOsCategories)
	return hostOsCategories, total, err
}
