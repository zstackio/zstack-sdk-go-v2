// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMetricDataHttpReceiver queries MetricDataHttpReceiver list
func (cli *ZSClient) QueryMetricDataHttpReceiver(params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List("v1/zwatch/metrics/httpreceivers", params, &resp)
}

func (cli *ZSClient) GetMetricDataHttpReceiver(uuid string) (*view.MetricDataHttpReceiverInventoryView, error) {
	var resp view.MetricDataHttpReceiverInventoryView
	if err := cli.Get("v1/zwatch/metrics/httpreceivers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMetricDataHttpReceiver Pagination
func (cli *ZSClient) PageMetricDataHttpReceiver(params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, int, error) {
	var metricDataHttpReceivers []view.MetricDataHttpReceiverInventoryView
	total, err := cli.Page("v1/zwatch/metrics/httpreceivers", params, &metricDataHttpReceivers)
	return metricDataHttpReceivers, total, err
}
// DeleteMetricDataHttpReceiver deletes MetricDataHttpReceiver
func (cli *ZSClient) DeleteMetricDataHttpReceiver(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics/httpreceivers", uuid, string(deleteMode))
}
// CreateMetricDataHttpReceiver creates MetricDataHttpReceiver
func (cli *ZSClient) CreateMetricDataHttpReceiver(params param.CreateMetricDataHttpReceiverParam) (*view.MetricDataHttpReceiverInventoryView, error) {
	resp := view.MetricDataHttpReceiverInventoryView{}
	if err := cli.Post("v1/zwatch/metrics/httpreceivers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMetricData requires additional arguments.
//
// Deprecated: use DeleteMetricDataWithParams to supply the required fields.
func (cli *ZSClient) DeleteMetricData(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteMetricData requires additional arguments; use DeleteMetricDataWithParams")
}

// DeleteMetricDataWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteMetricDataWithParams(deleteMode param.DeleteMode, params param.DeleteMetricDataParamDetail) error {
	query := url.Values{
		"namespace":  {params.Namespace},
		"metricName": {params.MetricName},
		"labels":     params.Labels,
	}
	return cli.deleteActionWithQuery("v1/zwatch/metrics", "", "", deleteMode, query, "namespace", "metricName")
}

// PutMetricData operates on PutMetricData
func (cli *ZSClient) PutMetricData(params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.Post("v1/zwatch/metrics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
