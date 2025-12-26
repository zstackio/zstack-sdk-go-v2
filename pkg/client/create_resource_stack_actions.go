// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateResourceStack creates ResourceStack
func (cli *ZSClient) CreateResourceStack(params param.CreateResourceStackParam) (*view.CreateResourceStackEventView, error) {
	resp := view.CreateResourceStackEventView{}
	if err := cli.Post("v1/cloudformation/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
