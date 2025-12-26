// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVmToAffinityGroup adds VmToAffinityGroup
func (cli *ZSClient) AddVmToAffinityGroup(params param.AddVmToAffinityGroupParam) (*view.AddVmToAffinityGroupEventView, error) {
	resp := view.AddVmToAffinityGroupEventView{}
	if err := cli.Post("v1/affinity-groups/{affinityGroupUuid}/vm-instances/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
