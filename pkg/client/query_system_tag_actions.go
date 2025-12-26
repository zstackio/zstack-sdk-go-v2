// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySystemTag queries SystemTag list
func (cli *ZSClient) QuerySystemTag(params *param.QueryParam) ([]view.SystemTagInventoryView, error) {
	var resp []view.SystemTagInventoryView
	return resp, cli.List("v1/system-tags", params, &resp)
}
