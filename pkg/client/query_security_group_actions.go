// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySecurityGroup queries SecurityGroup list
func (cli *ZSClient) QuerySecurityGroup(params *param.QueryParam) ([]view.SecurityGroupInventoryView, error) {
	var resp []view.SecurityGroupInventoryView
	return resp, cli.List("v1/security-groups", params, &resp)
}
