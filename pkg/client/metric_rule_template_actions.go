// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMetricRuleTemplate queries MetricRuleTemplate list
func (cli *ZSClient) QueryMetricRuleTemplate(params *param.QueryParam) ([]view.MetricRuleTemplateInventoryView, error) {
	var resp []view.MetricRuleTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates/metricrules", params, &resp)
}
// DeleteMetricRuleTemplate deletes MetricRuleTemplate
func (cli *ZSClient) DeleteMetricRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/metricrules/{uuid}", uuid, string(deleteMode))
}
// UpdateMetricRuleTemplate updates MetricRuleTemplate
func (cli *ZSClient) UpdateMetricRuleTemplate(uuid string, params param.UpdateMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	var resp view.UpdateMetricRuleTemplateEventView
	if err := cli.Put("v1/zwatch/monitortemplates/metricrules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddMetricRuleTemplate adds MetricRuleTemplate
func (cli *ZSClient) AddMetricRuleTemplate(params param.AddMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	var resp view.AddMetricRuleTemplateEventView
	if err := cli.Post("v1/zwatch/monitortemplates/{monitorTemplateUuid}/metricrules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
