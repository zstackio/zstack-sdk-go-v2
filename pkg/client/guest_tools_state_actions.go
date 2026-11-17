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

// PageGuestToolsState Pagination
func (cli *ZSClient) PageGuestToolsState(params *param.QueryParam) ([]view.GuestToolsStateInventoryView, int, error) {
	var guestToolsStates []view.GuestToolsStateInventoryView
	total, err := cli.Page("v1/guesttools", params, &guestToolsStates)
	return guestToolsStates, total, err
}
// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(vmInstanceUuid string, params param.UpdateGuestToolsStateParam) (*view.GuestToolsStateInventoryView, error) {
	resp := view.GuestToolsStateInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"updateGuestToolsState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmGuestToolsInfo gets VmGuestToolsInfo by uuid
func (cli *ZSClient) GetVmGuestToolsInfo(uuid string) (*view.GetVmGuestToolsInfoView, error) {
	var resp view.GetVmGuestToolsInfoView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachGuestToolsIsoToVm operates on GuestToolsIsoToVm
func (cli *ZSClient) AttachGuestToolsIsoToVm(uuid string, params param.AttachGuestToolsIsoToVmParam) (*view.AttachGuestToolsIsoToVmEventView, error) {
	resp := view.AttachGuestToolsIsoToVmEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"attachGuestToolsIsoToVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLatestGuestToolsForVm gets LatestGuestToolsForVm by uuid
func (cli *ZSClient) GetLatestGuestToolsForVm(uuid string) (*view.GuestToolsInventoryView, error) {
	var resp view.GetLatestGuestToolsForVmView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "inventory", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
