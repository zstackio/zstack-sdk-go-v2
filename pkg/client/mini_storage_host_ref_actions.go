// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMiniStorageHostRef 查询MiniStorageHostRef列表
func (cli *ZSClient) QueryMiniStorageHostRef(params param.QueryParam) ([]view.QueryMiniStorageHostRefView, error) {
	var resp []view.QueryMiniStorageHostRefView
	return resp, cli.List("v1/primary-storage/mini/host-refs", &params, &resp)
}

