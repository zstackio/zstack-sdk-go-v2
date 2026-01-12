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
	return cli.DeleteWithSpec("v1/security-machines", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateSecurityMachine updates SecurityMachine
func (cli *ZSClient) UpdateSecurityMachine(uuid string, params param.UpdateSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	var resp view.UpdateSecurityMachineEventView
	err := cli.PutWithSpec("v1/security-machines", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySecurityMachine queries SecurityMachine list
func (cli *ZSClient) QuerySecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, error) {
	var resp []view.SecurityMachineInventoryView
	return resp, cli.List("v1/security-machines", params, &resp)
}

func (cli *ZSClient) GetSecurityMachine(uuid string) (*view.SecurityMachineInventoryView, error) {
	var resp view.SecurityMachineInventoryView
	if err := cli.Get("v1/security-machines", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
