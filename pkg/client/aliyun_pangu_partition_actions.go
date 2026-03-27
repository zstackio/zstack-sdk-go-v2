// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunPanguPartition adds AliyunPanguPartition
func (cli *ZSClient) AddAliyunPanguPartition(ctx context.Context, params param.AddAliyunPanguPartitionParam) (*view.AliyunPanguPartitionInventoryView, error) {
	resp := view.AliyunPanguPartitionInventoryView{}
	if err := cli.Post(ctx, "v1/aliyun/pangu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAliyunPanguPartition deletes AliyunPanguPartition
func (cli *ZSClient) DeleteAliyunPanguPartition(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/aliyun/pangu", uuid, string(deleteMode))
}
// QueryAliyunPanguPartition queries AliyunPanguPartition list
func (cli *ZSClient) QueryAliyunPanguPartition(ctx context.Context, params *param.QueryParam) ([]view.AliyunPanguPartitionInventoryView, error) {
	var resp []view.AliyunPanguPartitionInventoryView
	return resp, cli.List(ctx, "v1/aliyun/pangu", params, &resp)
}

func (cli *ZSClient) GetAliyunPanguPartition(ctx context.Context, uuid string) (*view.AliyunPanguPartitionInventoryView, error) {
	var resp view.AliyunPanguPartitionInventoryView
	if err := cli.Get(ctx, "v1/aliyun/pangu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunPanguPartition Pagination
func (cli *ZSClient) PageAliyunPanguPartition(ctx context.Context, params *param.QueryParam) ([]view.AliyunPanguPartitionInventoryView, int, error) {
	var aliyunPanguPartitions []view.AliyunPanguPartitionInventoryView
	total, err := cli.Page(ctx, "v1/aliyun/pangu", params, &aliyunPanguPartitions)
	return aliyunPanguPartitions, total, err
}
// UpdateAliyunPanguPartition updates AliyunPanguPartition
func (cli *ZSClient) UpdateAliyunPanguPartition(ctx context.Context, uuid string, params param.UpdateAliyunPanguPartitionParam) (*view.AliyunPanguPartitionInventoryView, error) {
	resp := view.AliyunPanguPartitionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/aliyun/pangu", uuid, "", map[string]interface{}{
		"updateAliyunPanguPartition": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
