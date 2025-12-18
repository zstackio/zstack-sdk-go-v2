// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAlertDataAck queries AlertDataAck list
func (cli *ZSClient) QueryAlertDataAck(params param.QueryParam) ([]view.AlertDataAckInventoryView, error) {
	var resp []view.AlertDataAckInventoryView
	return resp, cli.List("v1/zwatch/alert-histories/acknowledgments", &params, &resp)
}
