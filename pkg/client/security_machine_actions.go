// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSecurityMachine deletes SecurityMachine
func (cli *ZSClient) DeleteSecurityMachine(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-machines", uuid, string(deleteMode))
}
// UpdateSecurityMachine updates SecurityMachine
func (cli *ZSClient) UpdateSecurityMachine(uuid string, params param.UpdateSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Put("v1/security-machines", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySecurityMachine queries SecurityMachine list
func (cli *ZSClient) QuerySecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, error) {
	var resp []view.SecurityMachineInventoryView
	return resp, cli.List("v1/security-machines", params, &resp)
}

// PageSecurityMachine Pagination
func (cli *ZSClient) PageSecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, int, error) {
	var securityMachines []view.SecurityMachineInventoryView
	total, err := cli.Page("v1/security-machines", params, &securityMachines)
	return securityMachines, total, err
}
