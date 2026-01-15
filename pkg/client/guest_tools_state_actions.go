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

// PageGuestToolsState Pagination
func (cli *ZSClient) PageGuestToolsState(params *param.QueryParam) ([]view.GuestToolsStateInventoryView, int, error) {
	var guestToolsStates []view.GuestToolsStateInventoryView
	total, err := cli.Page("v1/guesttools", params, &guestToolsStates)
	return guestToolsStates, total, err
}
// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(vmInstanceUuid string, params param.UpdateGuestToolsStateParam) (*view.GuestToolsStateInventoryView, error) {
	resp := view.GuestToolsStateInventoryView{}
	if err := cli.Put("v1/vm-instances", vmInstanceUuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
