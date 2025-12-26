// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAlertDataAck queries AlertDataAck list
func (cli *ZSClient) QueryAlertDataAck(params *param.QueryParam) ([]view.AlertDataAckInventoryView, error) {
	var resp []view.AlertDataAckInventoryView
	return resp, cli.List("v1/zwatch/alert-histories/acknowledgments", params, &resp)
}
