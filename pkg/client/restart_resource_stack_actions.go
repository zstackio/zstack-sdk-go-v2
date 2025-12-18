// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RestartResourceStack 操作RestartResourceStack
func (cli *ZSClient) RestartResourceStack(uuid string, params param.RestartResourceStackParam) (*view.RestartResourceStackEventView, error) {
	resp := view.RestartResourceStackEventView{}
	if err := cli.Put("v1/cloudformation/stack/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

