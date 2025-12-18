// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestVmScriptExecutedRecordDetail 查询GuestVmScriptExecutedRecordDetail列表
func (cli *ZSClient) QueryGuestVmScriptExecutedRecordDetail(params param.QueryParam) ([]view.QueryGuestVmScriptExecutedRecordDetailView, error) {
	var resp []view.QueryGuestVmScriptExecutedRecordDetailView
	return resp, cli.List("v1/scripts/records/details", &params, &resp)
}

