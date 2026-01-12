// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMetricRuleTemplate queries MetricRuleTemplate list
func (cli *ZSClient) QueryMetricRuleTemplate(params *param.QueryParam) ([]view.MetricRuleTemplateInventoryView, error) {
	var resp []view.MetricRuleTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates/metricrules", params, &resp)
}

func (cli *ZSClient) GetMetricRuleTemplate(uuid string) (*view.MetricRuleTemplateInventoryView, error) {
	var resp view.MetricRuleTemplateInventoryView
	if err := cli.Get("v1/zwatch/monitortemplates/metricrules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMetricRuleTemplate deletes MetricRuleTemplate
func (cli *ZSClient) DeleteMetricRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/monitortemplates/metricrules", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateMetricRuleTemplate updates MetricRuleTemplate
func (cli *ZSClient) UpdateMetricRuleTemplate(uuid string, params param.UpdateMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	var resp view.UpdateMetricRuleTemplateEventView
	err := cli.PutWithSpec("v1/zwatch/monitortemplates/metricrules", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
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
