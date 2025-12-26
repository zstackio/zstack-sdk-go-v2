// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIscsiServer queries IscsiServer list
func (cli *ZSClient) QueryIscsiServer(params *param.QueryParam) ([]view.IscsiServerInventoryView, error) {
	var resp []view.IscsiServerInventoryView
	return resp, cli.List("v1/storage-devices/iscsi/servers", params, &resp)
}
