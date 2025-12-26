// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAffinityGroupState changes AffinityGroupState
func (cli *ZSClient) ChangeAffinityGroupState(uuid string, params param.ChangeAffinityGroupStateParam) (*view.ChangeAffinityGroupStateEventView, error) {
	resp := view.ChangeAffinityGroupStateEventView{}
	if err := cli.Put("v1/affinity-groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
