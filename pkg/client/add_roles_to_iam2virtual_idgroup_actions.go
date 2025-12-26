// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddRolesToIAM2VirtualIDGroup adds RolesToIAM2VirtualIDGroup
func (cli *ZSClient) AddRolesToIAM2VirtualIDGroup(params param.AddRolesToIAM2VirtualIDGroupParam) (*view.AddRolesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{groupUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
