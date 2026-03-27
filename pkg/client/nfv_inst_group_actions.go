// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateNfvInstGroup creates NfvInstGroup
func (cli *ZSClient) CreateNfvInstGroup(params param.CreateNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	resp := view.NfvInstGroupInventoryView{}
	if err := cli.Post("v1/nfvinstgroup/group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SyncNfvInstGroup operates on NfvInstGroup
func (cli *ZSClient) SyncNfvInstGroup(uuid string, params param.SyncNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	resp := view.NfvInstGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/nfvinstgroup/group", uuid, "", map[string]interface{}{
		"syncNfvInstGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateNfvInstGroup updates NfvInstGroup
func (cli *ZSClient) UpdateNfvInstGroup(uuid string, params param.UpdateNfvInstGroupParam) (*view.NfvInstGroupInventoryView, error) {
	resp := view.NfvInstGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/nfvinstgroup/group", uuid, "", map[string]interface{}{
		"updateNfvInstGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteNfvInstGroup deletes NfvInstGroup
func (cli *ZSClient) DeleteNfvInstGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/nfvinstgroup/group", uuid, string(deleteMode))
}
// QueryNfvInstGroup queries NfvInstGroup list
func (cli *ZSClient) QueryNfvInstGroup(params *param.QueryParam) ([]view.NfvInstGroupInventoryView, error) {
	var resp []view.NfvInstGroupInventoryView
	return resp, cli.List("v1/nfvinstgroup/group", params, &resp)
}

func (cli *ZSClient) GetNfvInstGroup(uuid string) (*view.NfvInstGroupInventoryView, error) {
	var resp view.NfvInstGroupInventoryView
	if err := cli.Get("v1/nfvinstgroup/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNfvInstGroup Pagination
func (cli *ZSClient) PageNfvInstGroup(params *param.QueryParam) ([]view.NfvInstGroupInventoryView, int, error) {
	var nfvInstGroups []view.NfvInstGroupInventoryView
	total, err := cli.Page("v1/nfvinstgroup/group", params, &nfvInstGroups)
	return nfvInstGroups, total, err
}
