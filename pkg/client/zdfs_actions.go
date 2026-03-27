// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectZdfs operates on Zdfs
func (cli *ZSClient) ReconnectZdfs(ctx context.Context, uuid string, params param.ReconnectZdfsParam) (*view.ZdfsInventoryView, error) {
	resp := view.ZdfsInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zdfs", uuid, "", map[string]interface{}{
		"reconnectZdfs": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryZdfs queries Zdfs list
func (cli *ZSClient) QueryZdfs(ctx context.Context, params *param.QueryParam) ([]view.ZdfsInventoryView, error) {
	var resp []view.ZdfsInventoryView
	return resp, cli.List(ctx, "v1/zdfs", params, &resp)
}

func (cli *ZSClient) GetZdfs(ctx context.Context, uuid string) (*view.ZdfsInventoryView, error) {
	var resp view.ZdfsInventoryView
	if err := cli.Get(ctx, "v1/zdfs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZdfs Pagination
func (cli *ZSClient) PageZdfs(ctx context.Context, params *param.QueryParam) ([]view.ZdfsInventoryView, int, error) {
	var zdfs []view.ZdfsInventoryView
	total, err := cli.Page(ctx, "v1/zdfs", params, &zdfs)
	return zdfs, total, err
}
