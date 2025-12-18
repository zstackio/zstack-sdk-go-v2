// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestVmScriptExecutedRecord 查询GuestVmScriptExecutedRecord列表
func (cli *ZSClient) QueryGuestVmScriptExecutedRecord(params param.QueryParam) ([]view.QueryGuestVmScriptExecutedRecordView, error) {
	var resp []view.QueryGuestVmScriptExecutedRecordView
	return resp, cli.List("v1/scripts/records", &params, &resp)
}

