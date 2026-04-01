// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
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

// PageMetricRuleTemplate Pagination
func (cli *ZSClient) PageMetricRuleTemplate(params *param.QueryParam) ([]view.MetricRuleTemplateInventoryView, int, error) {
	var metricRuleTemplates []view.MetricRuleTemplateInventoryView
	total, err := cli.Page("v1/zwatch/monitortemplates/metricrules", params, &metricRuleTemplates)
	return metricRuleTemplates, total, err
}
// DeleteMetricRuleTemplate deletes MetricRuleTemplate
func (cli *ZSClient) DeleteMetricRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/metricrules", uuid, string(deleteMode))
}
// UpdateMetricRuleTemplate updates MetricRuleTemplate
func (cli *ZSClient) UpdateMetricRuleTemplate(uuid string, params param.UpdateMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	resp := view.MetricRuleTemplateInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/monitortemplates/metricrules", uuid, "", map[string]interface{}{
		"updateMetricRuleTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddMetricRuleTemplate adds MetricRuleTemplate
func (cli *ZSClient) AddMetricRuleTemplate(monitorTemplateUuid string, params param.AddMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	resp := view.MetricRuleTemplateInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/monitortemplates/%s/metricrules", monitorTemplateUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
