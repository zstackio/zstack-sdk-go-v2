// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNativeCluster queries NativeCluster list
func (cli *ZSClient) QueryNativeCluster(ctx context.Context, params *param.QueryParam) ([]view.NativeClusterInventoryView, error) {
	var resp []view.NativeClusterInventoryView
	return resp, cli.List(ctx, "v1/container/native/cluster", params, &resp)
}

func (cli *ZSClient) GetNativeCluster(ctx context.Context, uuid string) (*view.NativeClusterInventoryView, error) {
	var resp view.NativeClusterInventoryView
	if err := cli.Get(ctx, "v1/container/native/cluster", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNativeCluster Pagination
func (cli *ZSClient) PageNativeCluster(ctx context.Context, params *param.QueryParam) ([]view.NativeClusterInventoryView, int, error) {
	var nativeClusters []view.NativeClusterInventoryView
	total, err := cli.Page(ctx, "v1/container/native/cluster", params, &nativeClusters)
	return nativeClusters, total, err
}
