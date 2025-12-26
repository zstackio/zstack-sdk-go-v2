// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAccessControlRule queries AccessControlRule list
func (cli *ZSClient) QueryAccessControlRule(params *param.QueryParam) ([]view.AccessControlRuleInventoryView, error) {
	var resp []view.AccessControlRuleInventoryView
	return resp, cli.List("v1/login-control/access-control/rules", params, &resp)
}
