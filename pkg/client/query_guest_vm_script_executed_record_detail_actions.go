// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestVmScriptExecutedRecordDetail queries GuestVmScriptExecutedRecordDetail list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecordDetail(params param.QueryParam) ([]view.GuestVmScriptExecutedRecordDetailInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordDetailInventoryView
	return resp, cli.List("v1/scripts/records/details", &params, &resp)
}
