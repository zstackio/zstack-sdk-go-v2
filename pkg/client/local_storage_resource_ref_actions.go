// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLocalStorageResourceRef 查询LocalStorageResourceRef列表
func (cli *ZSClient) QueryLocalStorageResourceRef(params param.QueryParam) ([]view.QueryLocalStorageResourceRefView, error) {
	var resp []view.QueryLocalStorageResourceRefView
	return resp, cli.List("v1/primary-storage/local-storage/resource-refs", &params, &resp)
}

