// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAliyunSnapshot updates AliyunSnapshot
func (cli *ZSClient) UpdateAliyunSnapshot(uuid string, params param.UpdateAliyunSnapshotParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/snapshot", uuid, "", map[string]interface{}{
		"updateAliyunSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunSnapshotFromRemote deletes AliyunSnapshotFromRemote
func (cli *ZSClient) DeleteAliyunSnapshotFromRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot", uuid, string(deleteMode))
}

// GCAliyunSnapshotRemote operates on GCAliyunSnapshotRemote
func (cli *ZSClient) GCAliyunSnapshotRemote(dataCenterUuid string, params param.GCAliyunSnapshotRemoteParam) (*view.GCAliyunSnapshotRemoteEventView, error) {
	resp := view.GCAliyunSnapshotRemoteEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/snapshot/%s/gc", dataCenterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAliyunSnapshotRemote operates on AliyunSnapshotRemote
func (cli *ZSClient) SyncAliyunSnapshotRemote(dataCenterUuid string, params param.SyncAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/snapshot/%s/sync", dataCenterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunSnapshotFromLocal queries AliyunSnapshotFromLocal list
func (cli *ZSClient) QueryAliyunSnapshotFromLocal(params *param.QueryParam) ([]view.AliyunSnapshotInventoryView, error) {
	var resp []view.AliyunSnapshotInventoryView
	return resp, cli.List("v1/hybrid/aliyun/snapshot", params, &resp)
}

func (cli *ZSClient) GetAliyunSnapshotFromLocal(uuid string) (*view.AliyunSnapshotInventoryView, error) {
	var resp view.AliyunSnapshotInventoryView
	if err := cli.Get("v1/hybrid/aliyun/snapshot", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunSnapshotFromLocal Pagination
func (cli *ZSClient) PageAliyunSnapshotFromLocal(params *param.QueryParam) ([]view.AliyunSnapshotInventoryView, int, error) {
	var aliyunSnapshotFromLocals []view.AliyunSnapshotInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/snapshot", params, &aliyunSnapshotFromLocals)
	return aliyunSnapshotFromLocals, total, err
}

// CreateAliyunSnapshotRemote creates AliyunSnapshotRemote
func (cli *ZSClient) CreateAliyunSnapshotRemote(params param.CreateAliyunSnapshotRemoteParam) (*view.AliyunSnapshotInventoryView, error) {
	resp := view.AliyunSnapshotInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunSnapshotFromLocal deletes AliyunSnapshotFromLocal
func (cli *ZSClient) DeleteAliyunSnapshotFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/snapshot", uuid, string(deleteMode))
}
