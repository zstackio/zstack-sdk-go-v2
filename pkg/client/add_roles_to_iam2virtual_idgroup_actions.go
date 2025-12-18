// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddRolesToIAM2VirtualIDGroup 操作AddRolesToIAM2VirtualIDGroup
func (cli *ZSClient) AddRolesToIAM2VirtualIDGroup(params param.AddRolesToIAM2VirtualIDGroupParam) (*view.AddRolesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{groupUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

