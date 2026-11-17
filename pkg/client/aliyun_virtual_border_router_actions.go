// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// TerminateVirtualBorderRouterRemote operates on TerminateVirtualBorderRouterRemote
func (cli *ZSClient) TerminateVirtualBorderRouterRemote(uuid string, params param.TerminateVirtualBorderRouterRemoteParam) (*view.TerminateVirtualBorderRouterRemoteEventView, error) {
	resp := view.TerminateVirtualBorderRouterRemoteEventView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/border-router", uuid, "", map[string]interface{}{
		"terminateVirtualBorderRouterRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVirtualBorderRouterFromLocal queries VirtualBorderRouterFromLocal list
func (cli *ZSClient) QueryVirtualBorderRouterFromLocal(params *param.QueryParam) ([]view.VirtualBorderRouterInventoryView, error) {
	var resp []view.VirtualBorderRouterInventoryView
	return resp, cli.List("v1/hybrid/aliyun/border-router", params, &resp)
}

func (cli *ZSClient) GetVirtualBorderRouterFromLocal(uuid string) (*view.VirtualBorderRouterInventoryView, error) {
	var resp view.VirtualBorderRouterInventoryView
	if err := cli.Get("v1/hybrid/aliyun/border-router", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualBorderRouterFromLocal Pagination
func (cli *ZSClient) PageVirtualBorderRouterFromLocal(params *param.QueryParam) ([]view.VirtualBorderRouterInventoryView, int, error) {
	var virtualBorderRouterFromLocals []view.VirtualBorderRouterInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/border-router", params, &virtualBorderRouterFromLocals)
	return virtualBorderRouterFromLocals, total, err
}

// DeleteVirtualBorderRouterLocal deletes VirtualBorderRouterLocal
func (cli *ZSClient) DeleteVirtualBorderRouterLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/border-router", uuid, string(deleteMode))
}

// UpdateVirtualBorderRouterRemote updates VirtualBorderRouterRemote
func (cli *ZSClient) UpdateVirtualBorderRouterRemote(uuid string, params param.UpdateVirtualBorderRouterRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	resp := view.VirtualBorderRouterInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/border-router", uuid, "", map[string]interface{}{
		"updateVirtualBorderRouterRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryVirtualBorderRouterRemote operates on yVirtualBorderRouterRemote
func (cli *ZSClient) RecoveryVirtualBorderRouterRemote(uuid string, params param.RecoveryVirtualBorderRouterRemoteParam) (*view.RecoveryVirtualBorderRouterRemoteEventView, error) {
	resp := view.RecoveryVirtualBorderRouterRemoteEventView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/border-router", uuid, "", map[string]interface{}{
		"recoveryVirtualBorderRouterRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVirtualBorderRouterFromRemote operates on VirtualBorderRouterFromRemote
func (cli *ZSClient) SyncVirtualBorderRouterFromRemote(dataCenterUuid string, params param.SyncVirtualBorderRouterFromRemoteParam) (*view.VirtualBorderRouterInventoryView, error) {
	resp := view.VirtualBorderRouterInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/border-router", dataCenterUuid, "", map[string]interface{}{
		"syncVirtualBorderRouterFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
