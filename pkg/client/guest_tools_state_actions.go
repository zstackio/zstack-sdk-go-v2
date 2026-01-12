// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestToolsState queries GuestToolsState list
func (cli *ZSClient) QueryGuestToolsState(params *param.QueryParam) ([]view.GuestToolsStateInventoryView, error) {
	var resp []view.GuestToolsStateInventoryView
	return resp, cli.List("v1/guesttools", params, &resp)
}

func (cli *ZSClient) GetGuestToolsState(uuid string) (*view.GuestToolsStateInventoryView, error) {
	var resp view.GuestToolsStateInventoryView
	if err := cli.Get("v1/guesttools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(vmInstanceUuid string, params param.UpdateGuestToolsStateParam) (*view.GuestToolsStateInventoryView, error) {
	var resp view.UpdateGuestToolsStateView
	err := cli.PutWithSpec("v1/vm-instances", fmt.Sprintf(\"%s/guesttools-state\", vmInstanceUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
