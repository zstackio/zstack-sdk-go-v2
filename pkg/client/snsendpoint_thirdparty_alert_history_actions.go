// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSEndpointThirdpartyAlertHistory queries SNSEndpointThirdpartyAlertHistory list
func (cli *ZSClient) QuerySNSEndpointThirdpartyAlertHistory(params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, error) {
	var resp []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	return resp, cli.List("v1/zwatch/third-party/alert-publish-histories", params, &resp)
}

// PageSNSEndpointThirdpartyAlertHistory Pagination
func (cli *ZSClient) PageSNSEndpointThirdpartyAlertHistory(params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, int, error) {
	var sNSEndpointThirdpartyAlertHistories []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	total, err := cli.Page("v1/zwatch/third-party/alert-publish-histories", params, &sNSEndpointThirdpartyAlertHistories)
	return sNSEndpointThirdpartyAlertHistories, total, err
}
