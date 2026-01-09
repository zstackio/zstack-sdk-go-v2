// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryActiveAlarmTemplate queries ActiveAlarmTemplate list
func (cli *ZSClient) QueryActiveAlarmTemplate(params *param.QueryParam) ([]view.ActiveAlarmTemplateInventoryView, error) {
	var resp []view.ActiveAlarmTemplateInventoryView
	return resp, cli.List("v1/zwatch/activealarms/templates", params, &resp)
}

func (cli *ZSClient) GetActiveAlarmTemplate(uuid string) (*view.ActiveAlarmTemplateInventoryView, error) {
	var resp view.ActiveAlarmTemplateInventoryView
	if err := cli.Get("v1/zwatch/activealarms/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateActiveAlarmTemplate updates ActiveAlarmTemplate
func (cli *ZSClient) UpdateActiveAlarmTemplate(uuid string, params param.UpdateActiveAlarmTemplateParam) (*view.ActiveAlarmTemplateInventoryView, error) {
	var resp view.UpdateActiveAlarmTemplateEventView
	if err := cli.Put("v1/zwatch/activealarms/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
