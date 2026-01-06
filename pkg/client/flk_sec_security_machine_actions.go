// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFlkSecSecurityMachine updates FlkSecSecurityMachine
func (cli *ZSClient) UpdateFlkSecSecurityMachine(uuid string, params param.UpdateFlkSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	if err := cli.Put("v1/security-machines/flkSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddFlkSecSecurityMachine adds FlkSecSecurityMachine
func (cli *ZSClient) AddFlkSecSecurityMachine(params param.AddFlkSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.AddSecurityMachineEventView
	if err := cli.Post("v1/security-machine/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
