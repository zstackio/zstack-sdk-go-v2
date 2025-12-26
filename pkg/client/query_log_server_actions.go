// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLogServer queries LogServer list
func (cli *ZSClient) QueryLogServer(params *param.QueryParam) ([]view.LogServerInventoryView, error) {
	var resp []view.LogServerInventoryView
	return resp, cli.List("v1/log/servers", params, &resp)
}
