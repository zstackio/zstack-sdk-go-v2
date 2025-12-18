// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEcsSecurityGroupRuleFromLocal 查询EcsSecurityGroupRuleFromLocal列表
func (cli *ZSClient) QueryEcsSecurityGroupRuleFromLocal(params param.QueryParam) ([]view.QueryEcsSecurityGroupRuleFromLocalView, error) {
	var resp []view.QueryEcsSecurityGroupRuleFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/security-group-rule", &params, &resp)
}

