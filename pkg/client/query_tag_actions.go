// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTag queries Tag list
func (cli *ZSClient) QueryTag(params *param.QueryParam) ([]view.TagPatternInventoryView, error) {
	var resp []view.TagPatternInventoryView
	return resp, cli.List("v1/tags", params, &resp)
}
