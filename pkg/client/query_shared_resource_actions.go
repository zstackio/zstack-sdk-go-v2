// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySharedResource queries SharedResource list
func (cli *ZSClient) QuerySharedResource(params *param.QueryParam) ([]view.SharedResourceInventoryView, error) {
	var resp []view.SharedResourceInventoryView
	return resp, cli.List("v1/accounts/resources", params, &resp)
}
