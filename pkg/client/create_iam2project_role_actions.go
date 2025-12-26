// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateIAM2ProjectRole creates IAM2ProjectRole
func (cli *ZSClient) CreateIAM2ProjectRole(params param.CreateIAM2ProjectRoleParam) (*view.CreateRoleEventView, error) {
	resp := view.CreateRoleEventView{}
	if err := cli.Post("v1/iam2/project-roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
