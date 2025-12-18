// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterPrimaryStorage 查询VCenterPrimaryStorage列表
func (cli *ZSClient) QueryVCenterPrimaryStorage(params param.QueryParam) ([]view.QueryVCenterPrimaryStorageView, error) {
	var resp []view.QueryVCenterPrimaryStorageView
	return resp, cli.List("v1/vcenters/primary-storage", &params, &resp)
}

