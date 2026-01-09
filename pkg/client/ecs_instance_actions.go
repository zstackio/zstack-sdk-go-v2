// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// StartEcsInstance starts EcsInstance
func (cli *ZSClient) StartEcsInstance(uuid string, params param.StartEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/ecs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteEcsInstance deletes EcsInstance
func (cli *ZSClient) DeleteEcsInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/ecs/remote", uuid, string(deleteMode))
}
// StopEcsInstance stops EcsInstance
func (cli *ZSClient) StopEcsInstance(uuid string, params param.StopEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/ecs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RebootEcsInstance operates on EcsInstance
func (cli *ZSClient) RebootEcsInstance(uuid string, params param.RebootEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Put("v1/hybrid/aliyun/ecs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEcsInstance updates EcsInstance
func (cli *ZSClient) UpdateEcsInstance(uuid string, params param.UpdateEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	var resp view.UpdateEcsInstanceEventView
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/ecs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
