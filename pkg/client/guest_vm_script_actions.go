// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestVmScript queries GuestVmScript list
func (cli *ZSClient) QueryGuestVmScript(params *param.QueryParam) ([]view.GuestVmScriptInventoryView, error) {
	var resp []view.GuestVmScriptInventoryView
	return resp, cli.List("v1/scripts", params, &resp)
}

func (cli *ZSClient) GetGuestVmScript(uuid string) (*view.GuestVmScriptInventoryView, error) {
	var resp view.GuestVmScriptInventoryView
	if err := cli.Get("v1/scripts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteGuestVmScript deletes GuestVmScript
func (cli *ZSClient) DeleteGuestVmScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/scripts", uuid, string(deleteMode))
}
// CreateGuestVmScript creates GuestVmScript
func (cli *ZSClient) CreateGuestVmScript(params param.CreateGuestVmScriptParam) (*view.GuestVmScriptInventoryView, error) {
	var resp view.CreateGuestVmScriptEventView
	if err := cli.Post("v1/scripts", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateGuestVmScript updates GuestVmScript
func (cli *ZSClient) UpdateGuestVmScript(uuid string, params param.UpdateGuestVmScriptParam) (*view.GuestVmScriptInventoryView, error) {
	var resp view.UpdateGuestVmScriptEventView
	if err := cli.Put("v1/scripts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
