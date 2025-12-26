// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVirtualBorderRouterFromRemote operates on SyncVirtualBorderRouterFromRemote
func (cli *ZSClient) SyncVirtualBorderRouterFromRemote(uuid string, params param.SyncVirtualBorderRouterFromRemoteParam) (*view.SyncVirtualBorderRouterFromRemoteEventView, error) {
	resp := view.SyncVirtualBorderRouterFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
