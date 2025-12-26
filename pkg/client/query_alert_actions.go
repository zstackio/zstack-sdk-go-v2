// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAlert queries Alert list
func (cli *ZSClient) QueryAlert(params *param.QueryParam) ([]view.AlertInventoryView, error) {
	var resp []view.AlertInventoryView
	return resp, cli.List("v1/monitoring/alerts", params, &resp)
}
