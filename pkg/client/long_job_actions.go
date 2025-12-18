// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLongJob 查询LongJob列表
func (cli *ZSClient) QueryLongJob(params param.QueryParam) ([]view.QueryLongJobView, error) {
	var resp []view.QueryLongJobView
	return resp, cli.List("v1/longjobs", &params, &resp)
}

