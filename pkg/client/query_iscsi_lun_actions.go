// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIscsiLun queries IscsiLun list
func (cli *ZSClient) QueryIscsiLun(params *param.QueryParam) ([]view.IscsiLunInventoryView, error) {
	var resp []view.IscsiLunInventoryView
	return resp, cli.List("v1/storage-devices/iscsi/luns", params, &resp)
}
