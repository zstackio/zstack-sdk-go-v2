// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryActiveAlarmTemplate queries ActiveAlarmTemplate list
func (cli *ZSClient) QueryActiveAlarmTemplate(params *param.QueryParam) ([]view.ActiveAlarmTemplateInventoryView, error) {
	var resp []view.ActiveAlarmTemplateInventoryView
	return resp, cli.List("v1/zwatch/activealarms/templates", params, &resp)
}
// UpdateActiveAlarmTemplate updates ActiveAlarmTemplate
func (cli *ZSClient) UpdateActiveAlarmTemplate(uuid string, params param.UpdateActiveAlarmTemplateParam) (*view.ActiveAlarmTemplateInventoryView, error) {
	var resp view.UpdateActiveAlarmTemplateEventView
	if err := cli.Put("v1/zwatch/activealarms/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
