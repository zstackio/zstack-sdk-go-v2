// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateAliyunVpcVirtualRouterEntryRemote creates AliyunVpcVirtualRouterEntryRemote
func (cli *ZSClient) CreateAliyunVpcVirtualRouterEntryRemote(params param.CreateAliyunVpcVirtualRouterEntryRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	resp := view.VpcVirtualRouteEntryInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/route-entry", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAliyunRouteEntryRemote deletes AliyunRouteEntryRemote
func (cli *ZSClient) DeleteAliyunRouteEntryRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/route-entry", uuid, string(deleteMode))
}

// QueryAliyunRouteEntryFromLocal queries AliyunRouteEntryFromLocal list
func (cli *ZSClient) QueryAliyunRouteEntryFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouteEntryInventoryView, error) {
	var resp []view.VpcVirtualRouteEntryInventoryView
	return resp, cli.List("v1/hybrid/aliyun/route-entry", params, &resp)
}

func (cli *ZSClient) GetAliyunRouteEntryFromLocal(uuid string) (*view.VpcVirtualRouteEntryInventoryView, error) {
	var resp view.VpcVirtualRouteEntryInventoryView
	if err := cli.Get("v1/hybrid/aliyun/route-entry", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunRouteEntryFromLocal Pagination
func (cli *ZSClient) PageAliyunRouteEntryFromLocal(params *param.QueryParam) ([]view.VpcVirtualRouteEntryInventoryView, int, error) {
	var aliyunRouteEntryFromLocals []view.VpcVirtualRouteEntryInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/route-entry", params, &aliyunRouteEntryFromLocals)
	return aliyunRouteEntryFromLocals, total, err
}

// SyncAliyunRouteEntryFromRemote operates on AliyunRouteEntryFromRemote
func (cli *ZSClient) SyncAliyunRouteEntryFromRemote(vRouterUuid string, params param.SyncAliyunRouteEntryFromRemoteParam) (*view.VpcVirtualRouteEntryInventoryView, error) {
	resp := view.VpcVirtualRouteEntryInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/route-entry", vRouterUuid, "", map[string]interface{}{
		"syncAliyunRouteEntryFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
