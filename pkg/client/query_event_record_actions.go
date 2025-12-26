// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEventRecord queries EventRecord list
func (cli *ZSClient) QueryEventRecord(params *param.QueryParam) ([]view.EventRecordsInventoryView, error) {
	var resp []view.EventRecordsInventoryView
	return resp, cli.List("v1/zwatch/event-records", params, &resp)
}
