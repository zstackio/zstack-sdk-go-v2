// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// StartEcsInstance starts EcsInstance
func (cli *ZSClient) StartEcsInstance(ctx context.Context, uuid string, params param.StartEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/ecs", uuid, "", map[string]interface{}{
		"startEcsInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteEcsInstance deletes EcsInstance
func (cli *ZSClient) DeleteEcsInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hybrid/aliyun/ecs", uuid, string(deleteMode))
}
// StopEcsInstance stops EcsInstance
func (cli *ZSClient) StopEcsInstance(ctx context.Context, uuid string, params param.StopEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/ecs", uuid, "", map[string]interface{}{
		"stopEcsInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RebootEcsInstance operates on EcsInstance
func (cli *ZSClient) RebootEcsInstance(ctx context.Context, uuid string, params param.RebootEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/ecs", uuid, "", map[string]interface{}{
		"rebootEcsInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEcsInstance updates EcsInstance
func (cli *ZSClient) UpdateEcsInstance(ctx context.Context, params param.UpdateEcsInstanceParam) (*view.EcsInstanceInventoryView, error) {
	resp := view.EcsInstanceInventoryView{}
	if err := cli.Post(ctx, "v1/hybrid/aliyun/{uuid}/ecs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
