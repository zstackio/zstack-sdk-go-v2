// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySystemTag 查询SystemTag列表
func (cli *ZSClient) QuerySystemTag(params param.QueryParam) ([]view.QuerySystemTagView, error) {
	var resp []view.QuerySystemTagView
	return resp, cli.List("v1/system-tags", &params, &resp)
}

