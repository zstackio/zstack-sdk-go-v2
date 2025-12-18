// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// TerminateVirtualBorderRouterRemote 操作TerminateVirtualBorderRouterRemote
func (cli *ZSClient) TerminateVirtualBorderRouterRemote(uuid string, params param.TerminateVirtualBorderRouterRemoteParam) (*view.TerminateVirtualBorderRouterRemoteEventView, error) {
	resp := view.TerminateVirtualBorderRouterRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

