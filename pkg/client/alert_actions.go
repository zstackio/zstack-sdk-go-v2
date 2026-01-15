// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAlert queries Alert list
func (cli *ZSClient) QueryAlert(params *param.QueryParam) ([]view.AlertInventoryView, error) {
	var resp []view.AlertInventoryView
	return resp, cli.List("v1/monitoring/alerts", params, &resp)
}

// PageAlert Pagination
func (cli *ZSClient) PageAlert(params *param.QueryParam) ([]view.AlertInventoryView, int, error) {
	var alerts []view.AlertInventoryView
	total, err := cli.Page("v1/monitoring/alerts", params, &alerts)
	return alerts, total, err
}
// DeleteAlert deletes Alert
func (cli *ZSClient) DeleteAlert(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/alerts", uuid, string(deleteMode))
}
