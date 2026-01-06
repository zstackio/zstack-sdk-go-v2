// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestVmScriptExecutedRecordDetail queries GuestVmScriptExecutedRecordDetail list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecordDetail(params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordDetailInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordDetailInventoryView
	return resp, cli.List("v1/scripts/records/details", params, &resp)
}
