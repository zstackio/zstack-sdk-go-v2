// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEventLog queries EventLog list
func (cli *ZSClient) QueryEventLog(params *param.QueryParam) ([]view.EventLogInventoryView, error) {
	var resp []view.EventLogInventoryView
	return resp, cli.List("v1/eventlogs", params, &resp)
}
