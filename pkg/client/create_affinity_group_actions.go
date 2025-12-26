// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAffinityGroup creates AffinityGroup
func (cli *ZSClient) CreateAffinityGroup(params param.CreateAffinityGroupParam) (*view.CreateAffinityGroupEventView, error) {
	resp := view.CreateAffinityGroupEventView{}
	if err := cli.Post("v1/affinity-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
