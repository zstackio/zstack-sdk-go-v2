// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunEbsPrimaryStorage updates AliyunEbsPrimaryStorage
func (cli *ZSClient) UpdateAliyunEbsPrimaryStorage(uuid string, params param.UpdateAliyunEbsPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.UpdatePrimaryStorageEventView
	if err := cli.Put("v1/primary-storage/aliyun/ebs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAliyunEbsPrimaryStorage queries AliyunEbsPrimaryStorage list
func (cli *ZSClient) QueryAliyunEbsPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/aliyun/ebs", params, &resp)
}
// AddAliyunEbsPrimaryStorage adds AliyunEbsPrimaryStorage
func (cli *ZSClient) AddAliyunEbsPrimaryStorage(params param.AddAliyunEbsPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.AddPrimaryStorageEventView
	if err := cli.Post("v1/primary-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
