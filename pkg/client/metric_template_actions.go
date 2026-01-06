// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteMetricTemplate deletes MetricTemplate
func (cli *ZSClient) DeleteMetricTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics/httpreceivers/templates/{uuid}", uuid, string(deleteMode))
}
// QueryMetricTemplate queries MetricTemplate list
func (cli *ZSClient) QueryMetricTemplate(params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List("v1/zwatch/metrics/httpreceivers/templates", params, &resp)
}
// CreateMetricTemplate creates MetricTemplate
func (cli *ZSClient) CreateMetricTemplate(params param.CreateMetricTemplateParam) (*view.MetricTemplateInventoryView, error) {
	var resp view.CreateMetricTemplateEventView
	if err := cli.Post("v1/zwatch/metrics/httpreceivers/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
