// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddKVMHost adds KVMHost
func (cli *ZSClient) AddKVMHost(params param.AddKVMHostParam) (*view.HostInventoryView, error) {
	var resp view.AddHostEventView
	if err := cli.Post("v1/hosts/kvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateKVMHost updates KVMHost
func (cli *ZSClient) UpdateKVMHost(uuid string, params param.UpdateKVMHostParam) (*view.HostInventoryView, error) {
	var resp view.UpdateHostEventView
	if err := cli.Put("v1/hosts/kvm/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
