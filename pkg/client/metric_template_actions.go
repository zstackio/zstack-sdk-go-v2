// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteMetricTemplate deletes MetricTemplate
func (cli *ZSClient) DeleteMetricTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/metrics/httpreceivers/templates", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryMetricTemplate queries MetricTemplate list
func (cli *ZSClient) QueryMetricTemplate(params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List("v1/zwatch/metrics/httpreceivers/templates", params, &resp)
}

func (cli *ZSClient) GetMetricTemplate(uuid string) (*view.MetricDataHttpReceiverInventoryView, error) {
	var resp view.MetricDataHttpReceiverInventoryView
	if err := cli.Get("v1/zwatch/metrics/httpreceivers/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateMetricTemplate creates MetricTemplate
func (cli *ZSClient) CreateMetricTemplate(params param.CreateMetricTemplateParam) (*view.MetricTemplateInventoryView, error) {
	var resp view.CreateMetricTemplateEventView
	if err := cli.Post("v1/zwatch/metrics/httpreceivers/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
