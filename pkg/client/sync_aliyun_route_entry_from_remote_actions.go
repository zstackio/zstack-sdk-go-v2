// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncAliyunRouteEntryFromRemote operates on SyncAliyunRouteEntryFromRemote
func (cli *ZSClient) SyncAliyunRouteEntryFromRemote(uuid string, params param.SyncAliyunRouteEntryFromRemoteParam) (*view.SyncAliyunRouteEntryFromRemoteEventView, error) {
	resp := view.SyncAliyunRouteEntryFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/route-entry/{vRouterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
