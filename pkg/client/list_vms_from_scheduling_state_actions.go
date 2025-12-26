// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ListVmsFromSchedulingState operates on ListVmsFromSchedulingState
func (cli *ZSClient) ListVmsFromSchedulingState(params param.ListVmsFromSchedulingStateParam) (*view.ListVmsFromSchedulingStateView, error) {
	resp := view.ListVmsFromSchedulingStateView{}
	if err := cli.Post("v1/list/vms/from/executeState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
