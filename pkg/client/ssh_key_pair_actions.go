// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySshKeyPair 查询SshKeyPair列表
func (cli *ZSClient) QuerySshKeyPair(params param.QueryParam) ([]view.QuerySshKeyPairView, error) {
	var resp []view.QuerySshKeyPairView
	return resp, cli.List("v1/ssh-key-pair", &params, &resp)
}

