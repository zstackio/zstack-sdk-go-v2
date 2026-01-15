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

// PageActiveAlarmTemplate Pagination
func (cli *ZSClient) PageActiveAlarmTemplate(params *param.QueryParam) ([]view.ActiveAlarmTemplateInventoryView, int, error) {
	var activeAlarmTemplates []view.ActiveAlarmTemplateInventoryView
	total, err := cli.Page("v1/zwatch/activealarms/templates", params, &activeAlarmTemplates)
	return activeAlarmTemplates, total, err
}
// UpdateActiveAlarmTemplate updates ActiveAlarmTemplate
func (cli *ZSClient) UpdateActiveAlarmTemplate(uuid string, params param.UpdateActiveAlarmTemplateParam) (*view.ActiveAlarmTemplateInventoryView, error) {
	resp := view.ActiveAlarmTemplateInventoryView{}
	if err := cli.Put("v1/zwatch/activealarms/templates", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
