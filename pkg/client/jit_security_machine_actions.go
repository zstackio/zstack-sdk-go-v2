// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateJitSecurityMachine updates JitSecurityMachine
func (cli *ZSClient) UpdateJitSecurityMachine(uuid string, params param.UpdateJitSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	if err := cli.Put("v1/security-machines/jida/auth-gateway/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddJitSecurityMachine adds JitSecurityMachine
func (cli *ZSClient) AddJitSecurityMachine(params param.AddJitSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.AddSecurityMachineEventView
	if err := cli.Post("v1/security-machine/jida/auth-gateway", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
