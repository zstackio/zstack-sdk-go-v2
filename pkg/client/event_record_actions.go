// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventRecord 查询EventRecord列表
func (cli *ZSClient) QueryEventRecord(params param.QueryParam) ([]view.QueryEventRecordView, error) {
	var resp []view.QueryEventRecordView
	return resp, cli.List("v1/zwatch/event-records", &params, &resp)
}

