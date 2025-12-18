// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncVirtualBorderRouterFromRemote operates on SyncVirtualBorderRouterFromRemote
func (cli *ZSClient) SyncVirtualBorderRouterFromRemote(uuid string, params param.SyncVirtualBorderRouterFromRemoteParam) (*view.SyncVirtualBorderRouterFromRemoteEventView, error) {
	resp := view.SyncVirtualBorderRouterFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
