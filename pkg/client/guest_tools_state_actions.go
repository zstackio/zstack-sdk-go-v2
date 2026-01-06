// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestToolsState queries GuestToolsState list
func (cli *ZSClient) QueryGuestToolsState(params *param.QueryParam) ([]view.GuestToolsStateInventoryView, error) {
	var resp []view.GuestToolsStateInventoryView
	return resp, cli.List("v1/guesttools", params, &resp)
}
// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(uuid string, params param.UpdateGuestToolsStateParam) (*view.GuestToolsStateInventoryView, error) {
	var resp view.UpdateGuestToolsStateView
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/guesttools-state", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
