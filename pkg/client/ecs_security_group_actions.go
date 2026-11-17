// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsSecurityGroup updates EcsSecurityGroup
func (cli *ZSClient) UpdateEcsSecurityGroup(uuid string, params param.UpdateEcsSecurityGroupParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/security-group", uuid, "", map[string]interface{}{
		"updateEcsSecurityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsSecurityGroupFromRemote operates on EcsSecurityGroupFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupFromRemote(ecsVpcUuid string, params param.SyncEcsSecurityGroupFromRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/security-group", ecsVpcUuid, "", map[string]interface{}{
		"syncEcsSecurityGroupFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsSecurityGroupFromLocal queries EcsSecurityGroupFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupInventoryView, error) {
	var resp []view.EcsSecurityGroupInventoryView
	return resp, cli.List("v1/hybrid/aliyun/security-group", params, &resp)
}

func (cli *ZSClient) GetEcsSecurityGroupFromLocal(uuid string) (*view.EcsSecurityGroupInventoryView, error) {
	var resp view.EcsSecurityGroupInventoryView
	if err := cli.Get("v1/hybrid/aliyun/security-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsSecurityGroupFromLocal Pagination
func (cli *ZSClient) PageEcsSecurityGroupFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupInventoryView, int, error) {
	var ecsSecurityGroupFromLocals []view.EcsSecurityGroupInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/security-group", params, &ecsSecurityGroupFromLocals)
	return ecsSecurityGroupFromLocals, total, err
}

// CreateEcsSecurityGroupRemote creates EcsSecurityGroupRemote
func (cli *ZSClient) CreateEcsSecurityGroupRemote(params param.CreateEcsSecurityGroupRemoteParam) (*view.EcsSecurityGroupInventoryView, error) {
	resp := view.EcsSecurityGroupInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/security-group/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsSecurityGroupInLocal deletes EcsSecurityGroupInLocal
func (cli *ZSClient) DeleteEcsSecurityGroupInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group", uuid, string(deleteMode))
}

// DeleteEcsSecurityGroupRemote deletes EcsSecurityGroupRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group/remote", uuid, string(deleteMode))
}
