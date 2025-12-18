// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddVmToAffinityGroup 操作AddVmToAffinityGroup
func (cli *ZSClient) AddVmToAffinityGroup(params param.AddVmToAffinityGroupParam) (*view.AddVmToAffinityGroupEventView, error) {
	resp := view.AddVmToAffinityGroupEventView{}
	if err := cli.Post("v1/affinity-groups/{affinityGroupUuid}/vm-instances/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

