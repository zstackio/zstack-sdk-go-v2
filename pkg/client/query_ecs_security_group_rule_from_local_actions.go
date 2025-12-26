// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEcsSecurityGroupRuleFromLocal queries EcsSecurityGroupRuleFromLocal list
func (cli *ZSClient) QueryEcsSecurityGroupRuleFromLocal(params *param.QueryParam) ([]view.EcsSecurityGroupRuleInventoryView, error) {
	var resp []view.EcsSecurityGroupRuleInventoryView
	return resp, cli.List("v1/hybrid/aliyun/security-group-rule", params, &resp)
}
