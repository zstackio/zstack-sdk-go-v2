// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGCJob 查询GCJob列表
func (cli *ZSClient) QueryGCJob(params param.QueryParam) ([]view.QueryGCJobView, error) {
	var resp []view.QueryGCJobView
	return resp, cli.List("v1/gc-jobs", &params, &resp)
}

