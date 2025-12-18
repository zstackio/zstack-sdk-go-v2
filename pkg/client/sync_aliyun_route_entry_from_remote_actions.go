// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncAliyunRouteEntryFromRemote 操作SyncAliyunRouteEntryFromRemote
func (cli *ZSClient) SyncAliyunRouteEntryFromRemote(uuid string, params param.SyncAliyunRouteEntryFromRemoteParam) (*view.SyncAliyunRouteEntryFromRemoteEventView, error) {
	resp := view.SyncAliyunRouteEntryFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/route-entry/{vRouterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

