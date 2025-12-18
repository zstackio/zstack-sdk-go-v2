// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCephPrimaryStorage 查询CephPrimaryStorage列表
func (cli *ZSClient) QueryCephPrimaryStorage(params param.QueryParam) ([]view.QueryPrimaryStorageView, error) {
	var resp []view.QueryPrimaryStorageView
	return resp, cli.List("v1/primary-storage/ceph", &params, &resp)
}

