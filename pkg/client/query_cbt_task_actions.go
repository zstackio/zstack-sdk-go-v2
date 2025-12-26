// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCbtTask queries CbtTask list
func (cli *ZSClient) QueryCbtTask(params *param.QueryParam) ([]view.CbtTaskInventoryView, error) {
	var resp []view.CbtTaskInventoryView
	return resp, cli.List("v1/cbt-task", params, &resp)
}
