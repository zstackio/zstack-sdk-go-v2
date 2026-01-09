// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAliyunProxyVpc creates AliyunProxyVpc
func (cli *ZSClient) CreateAliyunProxyVpc(params param.CreateAliyunProxyVpcParam) (*view.AliyunProxyVpcInventoryView, error) {
	var resp view.CreateAliyunProxyVpcEventView
	if err := cli.Post("v1/aliyun-proxy/vpcs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAliyunProxyVpc updates AliyunProxyVpc
func (cli *ZSClient) UpdateAliyunProxyVpc(uuid string, params param.UpdateAliyunProxyVpcParam) (*view.AliyunProxyVpcInventoryView, error) {
	var resp view.UpdateAliyunProxyVpcEventView
	if err := cli.Put("v1/aliyun-proxy/vpcs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAliyunProxyVpc queries AliyunProxyVpc list
func (cli *ZSClient) QueryAliyunProxyVpc(params *param.QueryParam) ([]view.AliyunProxyVpcInventoryView, error) {
	var resp []view.AliyunProxyVpcInventoryView
	return resp, cli.List("v1/aliyun-proxy/vpcs", params, &resp)
}

func (cli *ZSClient) GetAliyunProxyVpc(uuid string) (*view.AliyunProxyVpcInventoryView, error) {
	var resp view.AliyunProxyVpcInventoryView
	if err := cli.Get("v1/aliyun-proxy/vpcs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAliyunProxyVpc deletes AliyunProxyVpc
func (cli *ZSClient) DeleteAliyunProxyVpc(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun-proxy/vpcs", uuid, string(deleteMode))
}
