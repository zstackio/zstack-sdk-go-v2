// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddRolesToIAM2VirtualID adds RolesToIAM2VirtualID
func (cli *ZSClient) AddRolesToIAM2VirtualID(params param.AddRolesToIAM2VirtualIDParam) (*view.AddRolesToIAM2VirtualIDEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
