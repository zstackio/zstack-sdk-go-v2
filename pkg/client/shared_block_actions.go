// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySharedBlock 查询SharedBlock列表
func (cli *ZSClient) QuerySharedBlock(params param.QueryParam) ([]view.QuerySharedBlockView, error) {
	var resp []view.QuerySharedBlockView
	return resp, cli.List("v1/sharedblock-group/sharedblocks", &params, &resp)
}

