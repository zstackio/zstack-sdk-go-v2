// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2ProjectRole 创建IAM2ProjectRole
func (cli *ZSClient) CreateIAM2ProjectRole(params param.CreateIAM2ProjectRoleParam) (*view.CreateRoleEventView, error) {
	resp := view.CreateRoleEventView{}
	if err := cli.Post("v1/iam2/project-roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

