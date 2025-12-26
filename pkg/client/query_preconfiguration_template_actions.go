// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPreconfigurationTemplate queries PreconfigurationTemplate list
func (cli *ZSClient) QueryPreconfigurationTemplate(params *param.QueryParam) ([]view.PreconfigurationTemplateInventoryView, error) {
	var resp []view.PreconfigurationTemplateInventoryView
	return resp, cli.List("v1/baremetal/preconfigurations", params, &resp)
}
