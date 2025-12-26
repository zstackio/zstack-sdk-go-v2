// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryGuestVmScriptExecutedRecord queries GuestVmScriptExecutedRecord list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecord(params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordInventoryView
	return resp, cli.List("v1/scripts/records", params, &resp)
}
