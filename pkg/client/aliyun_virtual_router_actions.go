// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// SyncAliyunVirtualRouterFromRemote operates on AliyunVirtualRouterFromRemote
func (cli *ZSClient) SyncAliyunVirtualRouterFromRemote(vpcUuid string, params param.SyncAliyunVirtualRouterFromRemoteParam) (*view.VpcVirtualRouterInventoryView, error) {
	resp := view.VpcVirtualRouterInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/vrouter", vpcUuid, "", map[string]interface{}{
		"syncAliyunVirtualRouterFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVirtualRouterLocal deletes VirtualRouterLocal
func (cli *ZSClient) DeleteVirtualRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/vrouter", uuid, string(deleteMode))
}

// UpdateAliyunVirtualRouter updates AliyunVirtualRouter
func (cli *ZSClient) UpdateAliyunVirtualRouter(uuid string, params param.UpdateAliyunVirtualRouterParam) (*view.VpcVirtualRouterInventoryView, error) {
	resp := view.VpcVirtualRouterInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/vrouter", uuid, "", map[string]interface{}{
		"updateAliyunVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAliyunVirtualRouterFromLocal queries AliyunVirtualRouterFromLocal list
func (cli *ZSClient) QueryAliyunVirtualRouterFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouterInventoryView, error) {
	var resp []view.VpcVirtualRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/vrouter", params, &resp)
}

func (cli *ZSClient) GetAliyunVirtualRouterFromLocal(uuid string) (*view.VpcVirtualRouterInventoryView, error) {
	var resp view.VpcVirtualRouterInventoryView
	if err := cli.Get("v1/hybrid/aliyun/vrouter", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunVirtualRouterFromLocal Pagination
func (cli *ZSClient) PageAliyunVirtualRouterFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouterInventoryView, int, error) {
	var aliyunVirtualRouterFromLocals []view.VpcVirtualRouterInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/vrouter", params, &aliyunVirtualRouterFromLocals)
	return aliyunVirtualRouterFromLocals, total, err
}
