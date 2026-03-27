// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAliyunProxyVpc creates AliyunProxyVpc
func (cli *ZSClient) CreateAliyunProxyVpc(ctx context.Context, params param.CreateAliyunProxyVpcParam) (*view.AliyunProxyVpcInventoryView, error) {
	resp := view.AliyunProxyVpcInventoryView{}
	if err := cli.Post(ctx, "v1/aliyun-proxy/vpcs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAliyunProxyVpc updates AliyunProxyVpc
func (cli *ZSClient) UpdateAliyunProxyVpc(ctx context.Context, uuid string, params param.UpdateAliyunProxyVpcParam) (*view.AliyunProxyVpcInventoryView, error) {
	resp := view.AliyunProxyVpcInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/aliyun-proxy/vpcs", uuid, "", map[string]interface{}{
		"updateAliyunProxyVpc": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAliyunProxyVpc queries AliyunProxyVpc list
func (cli *ZSClient) QueryAliyunProxyVpc(ctx context.Context, params *param.QueryParam) ([]view.AliyunProxyVpcInventoryView, error) {
	var resp []view.AliyunProxyVpcInventoryView
	return resp, cli.List(ctx, "v1/aliyun-proxy/vpcs", params, &resp)
}

func (cli *ZSClient) GetAliyunProxyVpc(ctx context.Context, uuid string) (*view.AliyunProxyVpcInventoryView, error) {
	var resp view.AliyunProxyVpcInventoryView
	if err := cli.Get(ctx, "v1/aliyun-proxy/vpcs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunProxyVpc Pagination
func (cli *ZSClient) PageAliyunProxyVpc(ctx context.Context, params *param.QueryParam) ([]view.AliyunProxyVpcInventoryView, int, error) {
	var aliyunProxyVpcs []view.AliyunProxyVpcInventoryView
	total, err := cli.Page(ctx, "v1/aliyun-proxy/vpcs", params, &aliyunProxyVpcs)
	return aliyunProxyVpcs, total, err
}
// DeleteAliyunProxyVpc deletes AliyunProxyVpc
func (cli *ZSClient) DeleteAliyunProxyVpc(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/aliyun-proxy/vpcs", uuid, string(deleteMode))
}
