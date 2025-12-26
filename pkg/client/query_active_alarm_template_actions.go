// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryActiveAlarmTemplate queries ActiveAlarmTemplate list
func (cli *ZSClient) QueryActiveAlarmTemplate(params *param.QueryParam) ([]view.ActiveAlarmTemplateInventoryView, error) {
	var resp []view.ActiveAlarmTemplateInventoryView
	return resp, cli.List("v1/zwatch/activealarms/templates", params, &resp)
}
