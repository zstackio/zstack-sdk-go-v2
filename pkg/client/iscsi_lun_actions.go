// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIscsiLun 查询IscsiLun列表
func (cli *ZSClient) QueryIscsiLun(params param.QueryParam) ([]view.QueryIscsiLunView, error) {
	var resp []view.QueryIscsiLunView
	return resp, cli.List("v1/storage-devices/iscsi/luns", &params, &resp)
}

