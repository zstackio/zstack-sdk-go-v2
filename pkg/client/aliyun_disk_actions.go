// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunDisk updates AliyunDisk
func (cli *ZSClient) UpdateAliyunDisk(uuid string, params param.UpdateAliyunDiskParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/disk", uuid, "", map[string]interface{}{
		"updateAliyunDisk": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunDiskFromLocal queries AliyunDiskFromLocal list
func (cli *ZSClient) QueryAliyunDiskFromLocal(params *param.QueryParam) ([]view.AliyunDiskInventoryView, error) {
	var resp []view.AliyunDiskInventoryView
	return resp, cli.List("v1/hybrid/aliyun/disk", params, &resp)
}

func (cli *ZSClient) GetAliyunDiskFromLocal(uuid string) (*view.AliyunDiskInventoryView, error) {
	var resp view.AliyunDiskInventoryView
	if err := cli.Get("v1/hybrid/aliyun/disk", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunDiskFromLocal Pagination
func (cli *ZSClient) PageAliyunDiskFromLocal(params *param.QueryParam) ([]view.AliyunDiskInventoryView, int, error) {
	var aliyunDiskFromLocals []view.AliyunDiskInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/disk", params, &aliyunDiskFromLocals)
	return aliyunDiskFromLocals, total, err
}

// DeleteAliyunDiskFromRemote deletes AliyunDiskFromRemote
func (cli *ZSClient) DeleteAliyunDiskFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk", uuid, string(deleteMode))
}

// DeleteAliyunDiskFromLocal deletes AliyunDiskFromLocal
func (cli *ZSClient) DeleteAliyunDiskFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/disk", uuid, string(deleteMode))
}

// CreateAliyunDiskFromRemote creates AliyunDiskFromRemote
func (cli *ZSClient) CreateAliyunDiskFromRemote(params param.CreateAliyunDiskFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/disk", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAliyunDiskFromEcs operates on AliyunDiskFromEcs
func (cli *ZSClient) DetachAliyunDiskFromEcs(uuid string, params param.DetachAliyunDiskFromEcsParam) (*view.DetachAliyunDiskFromEcsEventView, error) {
	resp := view.DetachAliyunDiskFromEcsEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/disk/%s/detach", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachAliyunDiskToEcs operates on AliyunDiskToEcs
func (cli *ZSClient) AttachAliyunDiskToEcs(diskUuid string, params param.AttachAliyunDiskToEcsParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/disk/%s/attach", diskUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncDiskFromAliyunFromRemote operates on DiskFromAliyunFromRemote
func (cli *ZSClient) SyncDiskFromAliyunFromRemote(identityUuid string, params param.SyncDiskFromAliyunFromRemoteParam) (*view.AliyunDiskInventoryView, error) {
	resp := view.AliyunDiskInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/disk/%s/sync", identityUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
