// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunPanguPartition adds AliyunPanguPartition
func (cli *ZSClient) AddAliyunPanguPartition(params param.AddAliyunPanguPartitionParam) (*view.AliyunPanguPartitionInventoryView, error) {
	var resp view.AddAliyunPanguPartitionEventView
	if err := cli.Post("v1/aliyun/pangu", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAliyunPanguPartition deletes AliyunPanguPartition
func (cli *ZSClient) DeleteAliyunPanguPartition(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/aliyun/pangu", uuid, string(deleteMode))
}
// QueryAliyunPanguPartition queries AliyunPanguPartition list
func (cli *ZSClient) QueryAliyunPanguPartition(params *param.QueryParam) ([]view.AliyunPanguPartitionInventoryView, error) {
	var resp []view.AliyunPanguPartitionInventoryView
	return resp, cli.List("v1/aliyun/pangu", params, &resp)
}

func (cli *ZSClient) GetAliyunPanguPartition(uuid string) (*view.AliyunPanguPartitionInventoryView, error) {
	var resp view.AliyunPanguPartitionInventoryView
	if err := cli.Get("v1/aliyun/pangu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAliyunPanguPartition updates AliyunPanguPartition
func (cli *ZSClient) UpdateAliyunPanguPartition(uuid string, params param.UpdateAliyunPanguPartitionParam) (*view.AliyunPanguPartitionInventoryView, error) {
	var resp view.UpdateAliyunPanguPartitionEventView
	if err := cli.Put("v1/aliyun/pangu", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
