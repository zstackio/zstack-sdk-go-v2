// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateNfvInstGroup creates NfvInstGroup
func (cli *ZSClient) CreateNfvInstGroup(ctx context.Context, params param.CreateNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	resp := view.NfvInstGroupInventoryView{}
	if err := cli.Post(ctx, "v1/nfvinstgroup/group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SyncNfvInstGroup operates on NfvInstGroup
func (cli *ZSClient) SyncNfvInstGroup(ctx context.Context, uuid string, params param.SyncNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	resp := view.NfvInstGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/nfvinstgroup/group", uuid, "", map[string]interface{}{
		"syncNfvInstGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateNfvInstGroup updates NfvInstGroup
func (cli *ZSClient) UpdateNfvInstGroup(ctx context.Context, uuid string, params param.UpdateNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	resp := view.NfvInstGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/nfvinstgroup/group", uuid, "", map[string]interface{}{
		"updateNfvInstGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteNfvInstGroup deletes NfvInstGroup
func (cli *ZSClient) DeleteNfvInstGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/nfvinstgroup/group", uuid, string(deleteMode))
}
// QueryNfvInstGroup queries NfvInstGroup list
func (cli *ZSClient) QueryNfvInstGroup(ctx context.Context, params *param.QueryParam) ([]view.NfvInstGroupInventoryView, error) {
	var resp []view.NfvInstGroupInventoryView
	return resp, cli.List(ctx, "v1/nfvinstgroup/group", params, &resp)
}

func (cli *ZSClient) GetNfvInstGroup(ctx context.Context, uuid string) (*view.NfvInstGroupInventoryView, error) {
	var resp view.NfvInstGroupInventoryView
	if err := cli.Get(ctx, "v1/nfvinstgroup/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNfvInstGroup Pagination
func (cli *ZSClient) PageNfvInstGroup(ctx context.Context, params *param.QueryParam) ([]view.NfvInstGroupInventoryView, int, error) {
	var nfvInstGroups []view.NfvInstGroupInventoryView
	total, err := cli.Page(ctx, "v1/nfvinstgroup/group", params, &nfvInstGroups)
	return nfvInstGroups, total, err
}
