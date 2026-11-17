// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// SyncEcsSecurityGroupRuleFromRemote operates on EcsSecurityGroupRuleFromRemote
func (cli *ZSClient) SyncEcsSecurityGroupRuleFromRemote(uuid string, params param.SyncEcsSecurityGroupRuleFromRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	resp := view.EcsSecurityGroupRuleInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/security-group-rule", uuid, "", map[string]interface{}{
		"syncEcsSecurityGroupRuleFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsSecurityGroupRuleFromLocal queries EcsSecurityGroupRuleFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupRuleFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupRuleInventoryView, error) {
	var resp []view.EcsSecurityGroupRuleInventoryView
	return resp, cli.List("v1/hybrid/aliyun/security-group-rule", params, &resp)
}

func (cli *ZSClient) GetEcsSecurityGroupRuleFromLocal(uuid string) (*view.EcsSecurityGroupRuleInventoryView, error) {
	var resp view.EcsSecurityGroupRuleInventoryView
	if err := cli.Get("v1/hybrid/aliyun/security-group-rule", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsSecurityGroupRuleFromLocal Pagination
func (cli *ZSClient) PageEcsSecurityGroupRuleFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupRuleInventoryView, int, error) {
	var ecsSecurityGroupRuleFromLocals []view.EcsSecurityGroupRuleInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/security-group-rule", params, &ecsSecurityGroupRuleFromLocals)
	return ecsSecurityGroupRuleFromLocals, total, err
}

// DeleteEcsSecurityGroupRuleRemote deletes EcsSecurityGroupRuleRemote
func (cli *ZSClient) DeleteEcsSecurityGroupRuleRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/security-group-rule/remote", uuid, string(deleteMode))
}

// CreateEcsSecurityGroupRuleRemote creates EcsSecurityGroupRuleRemote
func (cli *ZSClient) CreateEcsSecurityGroupRuleRemote(params param.CreateEcsSecurityGroupRuleRemoteParam) (*view.EcsSecurityGroupRuleInventoryView, error) {
	resp := view.EcsSecurityGroupRuleInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/security-group-rule", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
