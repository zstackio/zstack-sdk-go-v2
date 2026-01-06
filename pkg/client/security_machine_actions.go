// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSecurityMachine deletes SecurityMachine
func (cli *ZSClient) DeleteSecurityMachine(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-machines/{uuid}", uuid, string(deleteMode))
}
// UpdateSecurityMachine updates SecurityMachine
func (cli *ZSClient) UpdateSecurityMachine(uuid string, params param.UpdateSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	if err := cli.Put("v1/security-machines/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySecurityMachine queries SecurityMachine list
func (cli *ZSClient) QuerySecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, error) {
	var resp []view.SecurityMachineInventoryView
	return resp, cli.List("v1/security-machines", params, &resp)
}
