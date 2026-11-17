// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryConnectionAccessPointFromLocal queries ConnectionAccessPointFromLocal list
func (cli *ZSClient) QueryConnectionAccessPointFromLocal(params *param.QueryParam) ([]view.ConnectionAccessPointInventoryView, error) {
	var resp []view.ConnectionAccessPointInventoryView
	return resp, cli.List("v1/hybrid/aliyun/access-point", params, &resp)
}

func (cli *ZSClient) GetConnectionAccessPointFromLocal(uuid string) (*view.ConnectionAccessPointInventoryView, error) {
	var resp view.ConnectionAccessPointInventoryView
	if err := cli.Get("v1/hybrid/aliyun/access-point", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageConnectionAccessPointFromLocal Pagination
func (cli *ZSClient) PageConnectionAccessPointFromLocal(params *param.QueryParam) ([]view.ConnectionAccessPointInventoryView, int, error) {
	var connectionAccessPointFromLocals []view.ConnectionAccessPointInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/access-point", params, &connectionAccessPointFromLocals)
	return connectionAccessPointFromLocals, total, err
}

// GetConnectionAccessPointFromRemote gets ConnectionAccessPointFromRemote by uuid
func (cli *ZSClient) GetConnectionAccessPointFromRemote(uuid string) (*view.ConnectionAccessPointInventoryView, error) {
	var resp view.ConnectionAccessPointInventoryView
	if err := cli.GetWithRespKey("v1/hybrid/aliyun/access-point", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddConnectionAccessPointFromRemote adds ConnectionAccessPointFromRemote
func (cli *ZSClient) AddConnectionAccessPointFromRemote(params param.AddConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	resp := view.ConnectionAccessPointInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/access-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteConnectionAccessPointLocal deletes ConnectionAccessPointLocal
func (cli *ZSClient) DeleteConnectionAccessPointLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/access-point", uuid, string(deleteMode))
}

// SyncConnectionAccessPointFromRemote operates on ConnectionAccessPointFromRemote
func (cli *ZSClient) SyncConnectionAccessPointFromRemote(dataCenterUuid string, params param.SyncConnectionAccessPointFromRemoteParam) (*view.ConnectionAccessPointInventoryView, error) {
	resp := view.ConnectionAccessPointInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/access-point", dataCenterUuid, "", map[string]interface{}{
		"syncConnectionAccessPointFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
