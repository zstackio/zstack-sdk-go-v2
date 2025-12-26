// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncAliyunVirtualRouterFromRemote operates on SyncAliyunVirtualRouterFromRemote
func (cli *ZSClient) SyncAliyunVirtualRouterFromRemote(uuid string, params param.SyncAliyunVirtualRouterFromRemoteParam) (*view.SyncAliyunVirtualRouterFromRemoteEventView, error) {
	resp := view.SyncAliyunVirtualRouterFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{vpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
