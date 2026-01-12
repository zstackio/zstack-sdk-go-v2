// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddSanSecSecurityMachine adds SanSecSecurityMachine
func (cli *ZSClient) AddSanSecSecurityMachine(params param.AddSanSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.AddSecurityMachineEventView
	if err := cli.Post("v1/security-machine/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSanSecSecurityMachine updates SanSecSecurityMachine
func (cli *ZSClient) UpdateSanSecSecurityMachine(uuid string, params param.UpdateSanSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	if err := cli.Put("v1/security-machines/sanSec", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
