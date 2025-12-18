// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoveryVirtualBorderRouterRemote 操作yVirtualBorderRouterRemote
func (cli *ZSClient) RecoveryVirtualBorderRouterRemote(uuid string, params param.RecoveryVirtualBorderRouterRemoteParam) (*view.RecoveryVirtualBorderRouterRemoteEventView, error) {
	resp := view.RecoveryVirtualBorderRouterRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

