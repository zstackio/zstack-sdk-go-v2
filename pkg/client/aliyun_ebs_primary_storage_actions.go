// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunEbsPrimaryStorage updates AliyunEbsPrimaryStorage
func (cli *ZSClient) UpdateAliyunEbsPrimaryStorage(uuid string, params param.UpdateAliyunEbsPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.UpdatePrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/aliyun/ebs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAliyunEbsPrimaryStorage queries AliyunEbsPrimaryStorage list
func (cli *ZSClient) QueryAliyunEbsPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/aliyun/ebs", params, &resp)
}

func (cli *ZSClient) GetAliyunEbsPrimaryStorage(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage/aliyun/ebs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddAliyunEbsPrimaryStorage adds AliyunEbsPrimaryStorage
func (cli *ZSClient) AddAliyunEbsPrimaryStorage(params param.AddAliyunEbsPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
