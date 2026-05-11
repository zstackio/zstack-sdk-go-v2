// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestVmScript queries GuestVmScript list
func (cli *ZSClient) QueryGuestVmScript(ctx context.Context, params *param.QueryParam) ([]view.GuestVmScriptInventoryView, error) {
	var resp []view.GuestVmScriptInventoryView
	return resp, cli.List(ctx, "v1/scripts", params, &resp)
}

func (cli *ZSClient) GetGuestVmScript(ctx context.Context, uuid string) (*view.GuestVmScriptInventoryView, error) {
	var resp view.GuestVmScriptInventoryView
	if err := cli.Get(ctx, "v1/scripts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGuestVmScript Pagination
func (cli *ZSClient) PageGuestVmScript(ctx context.Context, params *param.QueryParam) ([]view.GuestVmScriptInventoryView, int, error) {
	var guestVmScripts []view.GuestVmScriptInventoryView
	total, err := cli.Page(ctx, "v1/scripts", params, &guestVmScripts)
	return guestVmScripts, total, err
}
// DeleteGuestVmScript deletes GuestVmScript
func (cli *ZSClient) DeleteGuestVmScript(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/scripts", uuid, string(deleteMode))
}
// CreateGuestVmScript creates GuestVmScript
func (cli *ZSClient) CreateGuestVmScript(ctx context.Context, params param.CreateGuestVmScriptParam) (*view.GuestVmScriptInventoryView, error) {
	resp := view.GuestVmScriptInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/scripts", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateGuestVmScript updates GuestVmScript
func (cli *ZSClient) UpdateGuestVmScript(ctx context.Context, uuid string, params param.UpdateGuestVmScriptParam) (*view.GuestVmScriptInventoryView, error) {
	resp := view.GuestVmScriptInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/scripts", uuid, "", map[string]interface{}{
		"updateGuestVmScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
