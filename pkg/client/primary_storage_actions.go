// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPrimaryStorage 查询PrimaryStorage列表
func (cli *ZSClient) QueryPrimaryStorage(params param.QueryParam) ([]view.QueryPrimaryStorageView, error) {
	var resp []view.QueryPrimaryStorageView
	return resp, cli.List("v1/primary-storage", &params, &resp)
}

