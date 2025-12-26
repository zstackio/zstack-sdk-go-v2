// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEcsSecurityGroupFromLocal queries EcsSecurityGroupFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupInventoryView, error) {
	var resp []view.EcsSecurityGroupInventoryView
	return resp, cli.List("v1/hybrid/aliyun/security-group", params, &resp)
}
