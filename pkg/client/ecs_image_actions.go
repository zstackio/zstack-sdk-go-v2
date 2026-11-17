// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsImage updates EcsImage
func (cli *ZSClient) UpdateEcsImage(uuid string, params param.UpdateEcsImageParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/image", uuid, "", map[string]interface{}{
		"updateEcsImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCreateEcsImageProgress gets CreateEcsImageProgress by uuid
func (cli *ZSClient) GetCreateEcsImageProgress(dataCenterUuid string, imageUuid string) (*view.GetCreateEcsImageProgressView, error) {
	var resp view.GetCreateEcsImageProgressView
	err := cli.GetWithSpec("v1/hybrid/aliyun/image", dataCenterUuid, fmt.Sprintf("%s/progress", imageUuid), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsImageFromRemote operates on EcsImageFromRemote
func (cli *ZSClient) SyncEcsImageFromRemote(dataCenterUuid string, params param.SyncEcsImageFromRemoteParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/image/%s/sync", dataCenterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsImageRemote deletes EcsImageRemote
func (cli *ZSClient) DeleteEcsImageRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/image/remote", uuid, string(deleteMode))
}

// CreateEcsImageFromEcsSnapshot creates EcsImageFromEcsSnapshot
func (cli *ZSClient) CreateEcsImageFromEcsSnapshot(params param.CreateEcsImageFromEcsSnapshotParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/image/snapshot", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsImageFromLocalImage creates EcsImageFromLocalImage
func (cli *ZSClient) CreateEcsImageFromLocalImage(params param.CreateEcsImageFromLocalImageParam) (*view.EcsImageInventoryView, error) {
	resp := view.EcsImageInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/image", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsImageLocal deletes EcsImageLocal
func (cli *ZSClient) DeleteEcsImageLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/image", uuid, string(deleteMode))
}

// QueryEcsImageFromLocal queries EcsImageFromLocal list
func (cli *ZSClient) QueryEcsImageFromLocal(params *param.QueryParam) ([]view.EcsImageInventoryView, error) {
	var resp []view.EcsImageInventoryView
	return resp, cli.List("v1/hybrid/aliyun/image", params, &resp)
}

func (cli *ZSClient) GetEcsImageFromLocal(uuid string) (*view.EcsImageInventoryView, error) {
	var resp view.EcsImageInventoryView
	if err := cli.Get("v1/hybrid/aliyun/image", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsImageFromLocal Pagination
func (cli *ZSClient) PageEcsImageFromLocal(params *param.QueryParam) ([]view.EcsImageInventoryView, int, error) {
	var ecsImageFromLocals []view.EcsImageInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/image", params, &ecsImageFromLocals)
	return ecsImageFromLocals, total, err
}
