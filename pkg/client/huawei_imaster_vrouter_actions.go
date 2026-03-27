// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateHuaweiIMasterVRouter creates HuaweiIMasterVRouter
func (cli *ZSClient) CreateHuaweiIMasterVRouter(ctx context.Context, params param.CreateHuaweiIMasterVRouterParam) (*view.HuaweiIMasterVRouterInventoryView, error) {
	resp := view.HuaweiIMasterVRouterInventoryView{}
	if err := cli.Post(ctx, "v1/sdn-controller/huawei-imaster/vrouters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHuaweiIMasterVRouter queries HuaweiIMasterVRouter list
func (cli *ZSClient) QueryHuaweiIMasterVRouter(ctx context.Context, params *param.QueryParam) ([]view.HuaweiIMasterVRouterInventoryView, error) {
	var resp []view.HuaweiIMasterVRouterInventoryView
	return resp, cli.List(ctx, "v1/sdn-controller/huawei-imaster/vrouters", params, &resp)
}

func (cli *ZSClient) GetHuaweiIMasterVRouter(ctx context.Context, uuid string) (*view.HuaweiIMasterVRouterInventoryView, error) {
	var resp view.HuaweiIMasterVRouterInventoryView
	if err := cli.Get(ctx, "v1/sdn-controller/huawei-imaster/vrouters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHuaweiIMasterVRouter Pagination
func (cli *ZSClient) PageHuaweiIMasterVRouter(ctx context.Context, params *param.QueryParam) ([]view.HuaweiIMasterVRouterInventoryView, int, error) {
	var huaweiIMasterVRouters []view.HuaweiIMasterVRouterInventoryView
	total, err := cli.Page(ctx, "v1/sdn-controller/huawei-imaster/vrouters", params, &huaweiIMasterVRouters)
	return huaweiIMasterVRouters, total, err
}
// DeleteHuaweiIMasterVRouter deletes HuaweiIMasterVRouter
func (cli *ZSClient) DeleteHuaweiIMasterVRouter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sdn-controller/huawei-imaster/vrouters", uuid, string(deleteMode))
}
