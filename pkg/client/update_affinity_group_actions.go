// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAffinityGroup updates AffinityGroup
func (cli *ZSClient) UpdateAffinityGroup(uuid string, params param.UpdateAffinityGroupParam) (*view.UpdateAffinityGroupEventView, error) {
	resp := view.UpdateAffinityGroupEventView{}
	if err := cli.Put("v1/affinity-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
