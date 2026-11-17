// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsVSwitch updates EcsVSwitch
func (cli *ZSClient) UpdateEcsVSwitch(uuid string, params param.UpdateEcsVSwitchParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/vswitch", uuid, "", map[string]interface{}{
		"updateEcsVSwitch": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncEcsVSwitchFromRemote operates on EcsVSwitchFromRemote
func (cli *ZSClient) SyncEcsVSwitchFromRemote(dataCenterUuid string, params param.SyncEcsVSwitchFromRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/aliyun/vswitch/%s/sync", dataCenterUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEcsVSwitchInLocal deletes EcsVSwitchInLocal
func (cli *ZSClient) DeleteEcsVSwitchInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch", uuid, string(deleteMode))
}

// DeleteEcsVSwitchRemote deletes EcsVSwitchRemote
func (cli *ZSClient) DeleteEcsVSwitchRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vswitch/remote", uuid, string(deleteMode))
}

// CreateEcsVSwitchRemote creates EcsVSwitchRemote
func (cli *ZSClient) CreateEcsVSwitchRemote(params param.CreateEcsVSwitchRemoteParam) (*view.EcsVSwitchInventoryView, error) {
	resp := view.EcsVSwitchInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/vswitch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEcsVSwitchFromLocal queries EcsVSwitchFromLocal list
func (cli *ZSClient) QueryEcsVSwitchFromLocal(params *param.QueryParam) ([]view.EcsVSwitchInventoryView, error) {
	var resp []view.EcsVSwitchInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vswitch", params, &resp)
}

func (cli *ZSClient) GetEcsVSwitchFromLocal(uuid string) (*view.EcsVSwitchInventoryView, error) {
	var resp view.EcsVSwitchInventoryView
	if err := cli.Get("v1/hybrid/aliyun/vswitch", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEcsVSwitchFromLocal Pagination
func (cli *ZSClient) PageEcsVSwitchFromLocal(params *param.QueryParam) ([]view.EcsVSwitchInventoryView, int, error) {
	var ecsVSwitchFromLocals []view.EcsVSwitchInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/vswitch", params, &ecsVSwitchFromLocals)
	return ecsVSwitchFromLocals, total, err
}
