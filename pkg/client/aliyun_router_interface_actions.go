// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// SyncAliyunRouterInterfaceFromRemote operates on AliyunRouterInterfaceFromRemote
func (cli *ZSClient) SyncAliyunRouterInterfaceFromRemote(dataCenterUuid string, params param.SyncAliyunRouterInterfaceFromRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/router-interface", dataCenterUuid, "", map[string]interface{}{
		"syncAliyunRouterInterfaceFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAliyunRouterInterfaceRemote creates AliyunRouterInterfaceRemote
func (cli *ZSClient) CreateAliyunRouterInterfaceRemote(params param.CreateAliyunRouterInterfaceRemoteParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/router-interface", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartConnectionBetweenAliyunRouterInterface starts ConnectionBetweenAliyunRouterInterface
func (cli *ZSClient) StartConnectionBetweenAliyunRouterInterface(vbrInterfaceUuid string, params param.StartConnectionBetweenAliyunRouterInterfaceParam) (*view.AliyunRouterInterfaceInventoryView, error) {
	resp := view.AliyunRouterInterfaceInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/router-interface", vbrInterfaceUuid, "", map[string]interface{}{
		"startConnectionBetweenAliyunRouterInterface": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunRouterInterfaceFromLocal queries AliyunRouterInterfaceFromLocal list
func (cli *ZSClient) QueryAliyunRouterInterfaceFromLocal(params *param.QueryParam) ([]view.AliyunRouterInterfaceInventoryView, error) {
	var resp []view.AliyunRouterInterfaceInventoryView
	return resp, cli.List("v1/hybrid/aliyun/router-interface", params, &resp)
}

func (cli *ZSClient) GetAliyunRouterInterfaceFromLocal(uuid string) (*view.AliyunRouterInterfaceInventoryView, error) {
	var resp view.AliyunRouterInterfaceInventoryView
	if err := cli.Get("v1/hybrid/aliyun/router-interface", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunRouterInterfaceFromLocal Pagination
func (cli *ZSClient) PageAliyunRouterInterfaceFromLocal(params *param.QueryParam) ([]view.AliyunRouterInterfaceInventoryView, int, error) {
	var aliyunRouterInterfaceFromLocals []view.AliyunRouterInterfaceInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/router-interface", params, &aliyunRouterInterfaceFromLocals)
	return aliyunRouterInterfaceFromLocals, total, err
}

// DeleteAliyunRouterInterfaceRemote deletes AliyunRouterInterfaceRemote
func (cli *ZSClient) DeleteAliyunRouterInterfaceRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface/remote", uuid, string(deleteMode))
}

// UpdateAliyunRouteInterfaceRemote updates AliyunRouteInterfaceRemote
func (cli *ZSClient) UpdateAliyunRouteInterfaceRemote(uuid string, params param.UpdateAliyunRouteInterfaceRemoteParam) (*view.UpdateAliyunRouteInterfaceRemoteEventView, error) {
	resp := view.UpdateAliyunRouteInterfaceRemoteEventView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/router-interface", uuid, "", map[string]interface{}{
		"updateAliyunRouteInterfaceRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunRouterInterfaceLocal deletes AliyunRouterInterfaceLocal
func (cli *ZSClient) DeleteAliyunRouterInterfaceLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/router-interface", uuid, string(deleteMode))
}
