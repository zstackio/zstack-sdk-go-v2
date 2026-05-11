// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CloneVmInstance operates on VmInstance
func (cli *ZSClient) CloneVmInstance(ctx context.Context, vmInstanceUuid string, params param.CloneVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"cloneVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeVmInstance operates on VmInstance
func (cli *ZSClient) ResumeVmInstance(ctx context.Context, uuid string, params param.ResumeVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"resumeVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartVmInstance starts VmInstance
func (cli *ZSClient) StartVmInstance(ctx context.Context, uuid string, params param.StartVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"startVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopVmInstance stops VmInstance
func (cli *ZSClient) StopVmInstance(ctx context.Context, uuid string, params param.StopVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"stopVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmInstance queries VmInstance list
func (cli *ZSClient) QueryVmInstance(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceInventoryView, error) {
	var resp []view.VmInstanceInventoryView
	return resp, cli.List(ctx, "v1/vm-instances", params, &resp)
}

func (cli *ZSClient) GetVmInstance(ctx context.Context, uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.Get(ctx, "v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstance Pagination
func (cli *ZSClient) PageVmInstance(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceInventoryView, int, error) {
	var vmInstances []view.VmInstanceInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances", params, &vmInstances)
	return vmInstances, total, err
}
// ExpungeVmInstance operates on VmInstance
func (cli *ZSClient) ExpungeVmInstance(ctx context.Context, uuid string) error {
	params := map[string]interface{}{
		"expungeVmInstance": map[string]interface{}{},
	}
	return cli.Put(ctx, "v1/vm-instances", uuid, params, nil)
}
// RebootVmInstance operates on VmInstance
func (cli *ZSClient) RebootVmInstance(ctx context.Context, uuid string, params param.RebootVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"rebootVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmInstance updates VmInstance
func (cli *ZSClient) UpdateVmInstance(ctx context.Context, uuid string, params param.UpdateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"updateVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DestroyVmInstance destroys VmInstance
func (cli *ZSClient) DestroyVmInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}
// CreateVmInstance creates VmInstance
func (cli *ZSClient) CreateVmInstance(ctx context.Context, params param.CreateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/vm-instances", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RecoverVmInstance operates on VmInstance
func (cli *ZSClient) RecoverVmInstance(ctx context.Context, uuid string, params param.RecoverVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances", uuid, "", map[string]interface{}{
		"recoverVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
