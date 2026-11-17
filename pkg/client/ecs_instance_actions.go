// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// StartEcsInstance starts EcsInstance
func (cli *ZSClient) StartEcsInstance(uuid string, params param.StartEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/ecs", uuid, "", map[string]interface{}{
		"startEcsInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteEcsInstance deletes EcsInstance
func (cli *ZSClient) DeleteEcsInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/ecs", uuid, string(deleteMode))
}
// StopEcsInstance stops EcsInstance
func (cli *ZSClient) StopEcsInstance(uuid string, params param.StopEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/ecs", uuid, "", map[string]interface{}{
		"stopEcsInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RebootEcsInstance operates on EcsInstance
func (cli *ZSClient) RebootEcsInstance(uuid string, params param.RebootEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/ecs", uuid, "", map[string]interface{}{
		"rebootEcsInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEcsInstance updates EcsInstance
func (cli *ZSClient) UpdateEcsInstance(uuid string, params param.UpdateEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/%s/ecs", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsInstanceLocal deletes EcsInstanceLocal
func (cli *ZSClient) DeleteEcsInstanceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/ecs", uuid, string(deleteMode))
}

// GetEcsInstanceVncUrl gets EcsInstanceVncUrl by uuid
func (cli *ZSClient) GetEcsInstanceVncUrl(uuid string) (*view.GetEcsInstanceVncUrlView, error) {
	var resp view.GetEcsInstanceVncUrlView
	if err := cli.GetWithRespKey("v1/hybrid/aliyun/ecs-vnc", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsInstanceFromLocal queries EcsInstanceFromLocal list
func (cli *ZSClient) QueryEcsInstanceFromLocal(params *param.QueryParam) ([]view.EcsInstanceInventoryView, error) {
	var resp []view.EcsInstanceInventoryView
	return resp, cli.List("v1/hybrid/aliyun/ecs", params, &resp)
}

func (cli *ZSClient) GetEcsInstanceFromLocal(uuid string) (*view.EcsInstanceInventoryView, error) {
	var resp view.EcsInstanceInventoryView
	if err := cli.Get("v1/hybrid/aliyun/ecs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsInstanceFromLocal Pagination
func (cli *ZSClient) PageEcsInstanceFromLocal(params *param.QueryParam) ([]view.EcsInstanceInventoryView, int, error) {
	var ecsInstanceFromLocals []view.EcsInstanceInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/ecs", params, &ecsInstanceFromLocals)
	return ecsInstanceFromLocals, total, err
}

// SyncEcsInstanceFromRemote operates on EcsInstanceFromRemote
func (cli *ZSClient) SyncEcsInstanceFromRemote(dataCenterUuid string, params param.SyncEcsInstanceFromRemoteParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/ecs/%s/sync", dataCenterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEcsInstanceFromEcsImage creates EcsInstanceFromEcsImage
func (cli *ZSClient) CreateEcsInstanceFromEcsImage(params param.CreateEcsInstanceFromEcsImageParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/ecs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEcsInstanceVncPassword updates EcsInstanceVncPassword
func (cli *ZSClient) UpdateEcsInstanceVncPassword(uuid string, params param.UpdateEcsInstanceVncPasswordParam) (*view.UpdateEcsInstanceVncPasswordEventView, error) {
	resp := view.UpdateEcsInstanceVncPasswordEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/%s/ecs-vnc", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEcsInstanceType gets EcsInstanceType by uuid
func (cli *ZSClient) GetEcsInstanceType() (*view.GetEcsInstanceTypeView, error) {
	var resp view.GetEcsInstanceTypeView
	if err := cli.GetWithRespKey("v1/hybrid/ecs/type", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
