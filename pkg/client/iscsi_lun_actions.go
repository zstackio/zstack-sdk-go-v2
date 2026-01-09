// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIscsiLun queries IscsiLun list
func (cli *ZSClient) QueryIscsiLun(params *param.QueryParam) ([]view.IscsiLunInventoryView, error) {
	var resp []view.IscsiLunInventoryView
	return resp, cli.List("v1/storage-devices/iscsi/luns", params, &resp)
}

func (cli *ZSClient) GetIscsiLun(uuid string) (*view.IscsiLunInventoryView, error) {
	var resp view.IscsiLunInventoryView
	if err := cli.Get("v1/storage-devices/iscsi/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
