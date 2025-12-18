// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ListVmsFromSchedulingState 操作ListVmsFromSchedulingState
func (cli *ZSClient) ListVmsFromSchedulingState(params param.ListVmsFromSchedulingStateParam) (*view.ListVmsFromSchedulingStateView, error) {
	resp := view.ListVmsFromSchedulingStateView{}
	if err := cli.Post("v1/list/vms/from/executeState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

