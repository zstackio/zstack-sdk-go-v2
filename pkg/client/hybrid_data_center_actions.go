// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// SyncDataCenterFromRemote operates on DataCenterFromRemote
func (cli *ZSClient) SyncDataCenterFromRemote(uuid string) (*view.SyncDataCenterFromRemoteEventView, error) {
	var resp view.SyncDataCenterFromRemoteEventView
	if err := cli.GetWithRespKey("v1/hybrid/data-center", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataCenterFromRemote gets DataCenterFromRemote by uuid
func (cli *ZSClient) GetDataCenterFromRemote() (*view.DataCenterPropertyView, error) {
	var resp view.DataCenterPropertyView
	if err := cli.GetWithRespKey("v1/hybrid/data-center/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDataCenterFromRemote adds DataCenterFromRemote
func (cli *ZSClient) AddDataCenterFromRemote(params param.AddDataCenterFromRemoteParam) (*view.DataCenterInventoryView, error) {
	resp := view.DataCenterInventoryView{}
	if err := cli.Post("v1/hybrid/data-center", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAllEcsInstancesFromDataCenter deletes AllEcsInstancesFromDataCenter
func (cli *ZSClient) DeleteAllEcsInstancesFromDataCenter(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/dc-ecs", uuid, string(deleteMode))
}

// QueryDataCenterFromLocal queries DataCenterFromLocal list
func (cli *ZSClient) QueryDataCenterFromLocal(params *param.QueryParam) ([]view.DataCenterInventoryView, error) {
	var resp []view.DataCenterInventoryView
	return resp, cli.List("v1/hybrid/data-center", params, &resp)
}

func (cli *ZSClient) GetDataCenterFromLocal(uuid string) (*view.DataCenterInventoryView, error) {
	var resp view.DataCenterInventoryView
	if err := cli.Get("v1/hybrid/data-center", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDataCenterFromLocal Pagination
func (cli *ZSClient) PageDataCenterFromLocal(params *param.QueryParam) ([]view.DataCenterInventoryView, int, error) {
	var dataCenterFromLocals []view.DataCenterInventoryView
	total, err := cli.Page("v1/hybrid/data-center", params, &dataCenterFromLocals)
	return dataCenterFromLocals, total, err
}

// DeleteDataCenterInLocal deletes DataCenterInLocal
func (cli *ZSClient) DeleteDataCenterInLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/data-center", uuid, string(deleteMode))
}
