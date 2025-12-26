// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RestartResourceStack operates on RestartResourceStack
func (cli *ZSClient) RestartResourceStack(uuid string, params param.RestartResourceStackParam) (*view.RestartResourceStackEventView, error) {
	resp := view.RestartResourceStackEventView{}
	if err := cli.Put("v1/cloudformation/stack/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
