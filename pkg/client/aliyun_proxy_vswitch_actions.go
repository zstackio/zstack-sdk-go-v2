// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunProxyVSwitch updates AliyunProxyVSwitch
func (cli *ZSClient) UpdateAliyunProxyVSwitch(ctx context.Context, uuid string, params param.UpdateAliyunProxyVSwitchParam) (*view.AliyunProxyVSwitchInventoryView, error) {
	resp := view.AliyunProxyVSwitchInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/aliyun-proxy/vswitches", uuid, "", map[string]interface{}{
		"updateAliyunProxyVSwitch": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAliyunProxyVSwitch queries AliyunProxyVSwitch list
func (cli *ZSClient) QueryAliyunProxyVSwitch(ctx context.Context, params *param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, error) {
	var resp []view.AliyunProxyVSwitchInventoryView
	return resp, cli.List(ctx, "v1/aliyun-proxy/vpcs/vswitches", params, &resp)
}

func (cli *ZSClient) GetAliyunProxyVSwitch(ctx context.Context, uuid string) (*view.AliyunProxyVSwitchInventoryView, error) {
	var resp view.AliyunProxyVSwitchInventoryView
	if err := cli.Get(ctx, "v1/aliyun-proxy/vpcs/vswitches", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunProxyVSwitch Pagination
func (cli *ZSClient) PageAliyunProxyVSwitch(ctx context.Context, params *param.QueryParam) ([]view.AliyunProxyVSwitchInventoryView, int, error) {
	var aliyunProxyVSwitchs []view.AliyunProxyVSwitchInventoryView
	total, err := cli.Page(ctx, "v1/aliyun-proxy/vpcs/vswitches", params, &aliyunProxyVSwitchs)
	return aliyunProxyVSwitchs, total, err
}
// CreateAliyunProxyVSwitch creates AliyunProxyVSwitch
func (cli *ZSClient) CreateAliyunProxyVSwitch(ctx context.Context, params param.CreateAliyunProxyVSwitchParam) (*view.AliyunProxyVSwitchInventoryView, error) {
	resp := view.AliyunProxyVSwitchInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/aliyun-proxy/vpcs/vswitches", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAliyunProxyVSwitch deletes AliyunProxyVSwitch
func (cli *ZSClient) DeleteAliyunProxyVSwitch(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/aliyun-proxy/vpcs/vswitches", uuid, string(deleteMode))
}
