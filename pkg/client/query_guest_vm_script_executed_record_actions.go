// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGuestVmScriptExecutedRecord queries GuestVmScriptExecutedRecord list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecord(params param.QueryParam) ([]view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordInventoryView
	return resp, cli.List("v1/scripts/records", &params, &resp)
}
