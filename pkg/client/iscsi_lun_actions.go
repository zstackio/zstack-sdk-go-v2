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

// PageIscsiLun Pagination
func (cli *ZSClient) PageIscsiLun(params *param.QueryParam) ([]view.IscsiLunInventoryView, int, error) {
	var iscsiLuns []view.IscsiLunInventoryView
	total, err := cli.Page("v1/storage-devices/iscsi/luns", params, &iscsiLuns)
	return iscsiLuns, total, err
}
