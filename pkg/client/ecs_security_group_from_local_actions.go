// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEcsSecurityGroupFromLocal 查询EcsSecurityGroupFromLocal列表
func (cli *ZSClient) QueryEcsSecurityGroupFromLocal(params param.QueryParam) ([]view.QueryEcsSecurityGroupFromLocalView, error) {
	var resp []view.QueryEcsSecurityGroupFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/security-group", &params, &resp)
}

