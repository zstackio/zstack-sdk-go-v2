// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySharedBlockGroupPrimaryStorageHostRef 查询SharedBlockGroupPrimaryStorageHostRef列表
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorageHostRef(params param.QueryParam) ([]view.QuerySharedBlockGroupPrimaryStorageHostRefView, error) {
	var resp []view.QuerySharedBlockGroupPrimaryStorageHostRefView
	return resp, cli.List("v1/sharedblock-group/host-refs", &params, &resp)
}

