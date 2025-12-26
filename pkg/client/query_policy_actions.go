// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicy queries Policy list
func (cli *ZSClient) QueryPolicy(params *param.QueryParam) ([]view.PolicyInventoryView, error) {
	var resp []view.PolicyInventoryView
	return resp, cli.List("v1/accounts/policies", params, &resp)
}
