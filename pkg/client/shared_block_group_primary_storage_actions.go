// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySharedBlockGroupPrimaryStorage 查询SharedBlockGroupPrimaryStorage列表
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorage(params param.QueryParam) ([]view.QuerySharedBlockGroupPrimaryStorageView, error) {
	var resp []view.QuerySharedBlockGroupPrimaryStorageView
	return resp, cli.List("v1/primary-storage/sharedblockgroup", &params, &resp)
}

