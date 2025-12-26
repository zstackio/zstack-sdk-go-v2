// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCdpTask queries CdpTask list
func (cli *ZSClient) QueryCdpTask(params *param.QueryParam) ([]view.CdpTaskInventoryView, error) {
	var resp []view.CdpTaskInventoryView
	return resp, cli.List("v1/cdp-task", params, &resp)
}
