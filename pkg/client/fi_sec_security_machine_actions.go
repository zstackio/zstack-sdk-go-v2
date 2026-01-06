// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFiSecSecurityMachine updates FiSecSecurityMachine
func (cli *ZSClient) UpdateFiSecSecurityMachine(uuid string, params param.UpdateFiSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	if err := cli.Put("v1/security-machines/fiSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddFiSecSecurityMachine adds FiSecSecurityMachine
func (cli *ZSClient) AddFiSecSecurityMachine(params param.AddFiSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.AddSecurityMachineEventView
	if err := cli.Post("v1/security-machine/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
