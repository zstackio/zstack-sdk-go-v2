// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddInfoSecSecurityMachine adds InfoSecSecurityMachine
func (cli *ZSClient) AddInfoSecSecurityMachine(params param.AddInfoSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.AddSecurityMachineEventView
	if err := cli.Post("v1/security-machine/infoSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateInfoSecSecurityMachine updates InfoSecSecurityMachine
func (cli *ZSClient) UpdateInfoSecSecurityMachine(uuid string, params param.UpdateInfoSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	err := cli.PutWithSpec("v1/security-machines/infoSec", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
