// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateResourceStack updates ResourceStack
func (cli *ZSClient) UpdateResourceStack(uuid string, params param.UpdateResourceStackParam) (*view.UpdateResourceStackEventView, error) {
	resp := view.UpdateResourceStackEventView{}
	if err := cli.Put("v1/cloudformation/stack/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
