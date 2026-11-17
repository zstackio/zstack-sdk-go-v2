// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsVpc updates EcsVpc
func (cli *ZSClient) UpdateEcsVpc(uuid string, params param.UpdateEcsVpcParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/vpc", uuid, "", map[string]interface{}{
		"updateEcsVpc": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsVpcFromLocal queries EcsVpcFromLocal list
func (cli *ZSClient) QueryEcsVpcFromLocal(params *param.QueryParam) ([]view.EcsVpcInventoryView, error) {
	var resp []view.EcsVpcInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vpc", params, &resp)
}

func (cli *ZSClient) GetEcsVpcFromLocal(uuid string) (*view.EcsVpcInventoryView, error) {
	var resp view.EcsVpcInventoryView
	if err := cli.Get("v1/hybrid/aliyun/vpc", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsVpcFromLocal Pagination
func (cli *ZSClient) PageEcsVpcFromLocal(params *param.QueryParam) ([]view.EcsVpcInventoryView, int, error) {
	var ecsVpcFromLocals []view.EcsVpcInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/vpc", params, &ecsVpcFromLocals)
	return ecsVpcFromLocals, total, err
}

// DeleteEcsVpcRemote deletes EcsVpcRemote
func (cli *ZSClient) DeleteEcsVpcRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc/remote", uuid, string(deleteMode))
}

// CreateEcsVpcRemote creates EcsVpcRemote
func (cli *ZSClient) CreateEcsVpcRemote(params param.CreateEcsVpcRemoteParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsVpcFromRemote operates on EcsVpcFromRemote
func (cli *ZSClient) SyncEcsVpcFromRemote(dataCenterUuid string, params param.SyncEcsVpcFromRemoteParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/vpc/%s/sync", dataCenterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsVpcInLocal deletes EcsVpcInLocal
func (cli *ZSClient) DeleteEcsVpcInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vpc", uuid, string(deleteMode))
}
