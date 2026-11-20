// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

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

func (cli *ZSClient) GetAlert(uuid string) (*view.AlertInventoryView, error) {
	var resp view.AlertInventoryView
	if err := cli.Get("v1/monitoring/alerts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlert Pagination
func (cli *ZSClient) PageAlert(params *param.QueryParam) ([]view.AlertInventoryView, int, error) {
	var alerts []view.AlertInventoryView
	total, err := cli.Page("v1/monitoring/alerts", params, &alerts)
	return alerts, total, err
}
// DeleteAlert sends one UUID to the collection endpoint.
func (cli *ZSClient) DeleteAlert(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteAlertWithParams(deleteMode, param.DeleteAlertParamDetail{Uuids: []string{uuid}})
}

// DeleteAlertWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteAlertWithParams(deleteMode param.DeleteMode, params param.DeleteAlertParamDetail) error {
	query := url.Values{
		"uuids": params.Uuids,
	}
	return cli.deleteActionWithQuery("v1/monitoring/alerts", "", "", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "uuids")
}
