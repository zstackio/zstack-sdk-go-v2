// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySNSEndpointThirdpartyAlertHistory queries SNSEndpointThirdpartyAlertHistory list
func (cli *ZSClient) QuerySNSEndpointThirdpartyAlertHistory(params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, error) {
	var resp []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	return resp, cli.List("v1/zwatch/third-party/alert-publish-histories", params, &resp)
}
