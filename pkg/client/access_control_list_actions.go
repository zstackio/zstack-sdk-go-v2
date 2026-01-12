// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAccessControlList creates AccessControlList
func (cli *ZSClient) CreateAccessControlList(params param.CreateAccessControlListParam) (*view.AccessControlListInventoryView, error) {
	var resp view.CreateAccessControlListEventView
	if err := cli.Post("v1/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAccessControlList updates AccessControlList
func (cli *ZSClient) UpdateAccessControlList(uuid string, params param.UpdateAccessControlListParam) (*view.AccessControlListInventoryView, error) {
	var resp view.UpdateAccessControlListEventView
	err := cli.PutWithSpec("v1/access-control-lists", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAccessControlList queries AccessControlList list
func (cli *ZSClient) QueryAccessControlList(params *param.QueryParam) ([]view.AccessControlListInventoryView, error) {
	var resp []view.AccessControlListInventoryView
	return resp, cli.List("v1/access-control-lists", params, &resp)
}

func (cli *ZSClient) GetAccessControlList(uuid string) (*view.AccessControlListInventoryView, error) {
	var resp view.AccessControlListInventoryView
	if err := cli.Get("v1/access-control-lists", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccessControlList deletes AccessControlList
func (cli *ZSClient) DeleteAccessControlList(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/access-control-lists", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
