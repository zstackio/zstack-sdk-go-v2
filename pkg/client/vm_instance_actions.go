// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CloneVmInstance operates on VmInstance
func (cli *ZSClient) CloneVmInstance(vmInstanceUuid string, params param.CloneVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", vmInstanceUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeVmInstance operates on VmInstance
func (cli *ZSClient) ResumeVmInstance(uuid string, params param.ResumeVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.ResumeVmInstanceEventView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StartVmInstance starts VmInstance
func (cli *ZSClient) StartVmInstance(uuid string, params param.StartVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.StartVmInstanceEventView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StopVmInstance stops VmInstance
func (cli *ZSClient) StopVmInstance(uuid string, params param.StopVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.StopVmInstanceEventView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVmInstance queries VmInstance list
func (cli *ZSClient) QueryVmInstance(params *param.QueryParam) ([]view.VmInstanceInventoryView, error) {
	var resp []view.VmInstanceInventoryView
	return resp, cli.List("v1/vm-instances", params, &resp)
}

func (cli *ZSClient) GetVmInstance(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.Get("v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ExpungeVmInstance operates on VmInstance
func (cli *ZSClient) ExpungeVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), string(deleteMode))
}
// RebootVmInstance operates on VmInstance
func (cli *ZSClient) RebootVmInstance(uuid string, params param.RebootVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.RebootVmInstanceEventView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateVmInstance updates VmInstance
func (cli *ZSClient) UpdateVmInstance(uuid string, params param.UpdateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.UpdateVmInstanceEventView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DestroyVmInstance destroys VmInstance
func (cli *ZSClient) DestroyVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vm-instances", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// CreateVmInstance creates VmInstance
func (cli *ZSClient) CreateVmInstance(params param.CreateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.CreateVmInstanceEventView
	if err := cli.Post("v1/vm-instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// RecoverVmInstance operates on VmInstance
func (cli *ZSClient) RecoverVmInstance(uuid string, params param.RecoverVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	var resp view.RecoverVmInstanceEventView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
