// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAccessKey creates AccessKey
func (cli *ZSClient) CreateAccessKey(params param.CreateAccessKeyParam) (*view.AccessKeyInventoryView, error) {
	var resp view.CreateAccessKeyEventView
	if err := cli.Post("v1/accesskeys", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAccessKey queries AccessKey list
func (cli *ZSClient) QueryAccessKey(params *param.QueryParam) ([]view.AccessKeyInventoryView, error) {
	var resp []view.AccessKeyInventoryView
	return resp, cli.List("v1/accesskeys", params, &resp)
}

func (cli *ZSClient) GetAccessKey(uuid string) (*view.AccessKeyInventoryView, error) {
	var resp view.AccessKeyInventoryView
	if err := cli.Get("v1/accesskeys", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccessKey deletes AccessKey
func (cli *ZSClient) DeleteAccessKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accesskeys", uuid, string(deleteMode))
}
