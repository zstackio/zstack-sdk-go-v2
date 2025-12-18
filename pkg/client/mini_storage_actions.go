// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMiniStorage 查询MiniStorage列表
func (cli *ZSClient) QueryMiniStorage(params param.QueryParam) ([]view.QueryMiniStorageView, error) {
	var resp []view.QueryMiniStorageView
	return resp, cli.List("v1/primary-storage/mini", &params, &resp)
}

