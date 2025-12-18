// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmNicInSecurityGroup 查询VmNicInSecurityGroup列表
func (cli *ZSClient) QueryVmNicInSecurityGroup(params param.QueryParam) ([]view.QueryVmNicInSecurityGroupView, error) {
	var resp []view.QueryVmNicInSecurityGroupView
	return resp, cli.List("v1/security-groups/vm-instances/nics", &params, &resp)
}

