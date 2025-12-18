// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncAliyunVirtualRouterFromRemote 操作SyncAliyunVirtualRouterFromRemote
func (cli *ZSClient) SyncAliyunVirtualRouterFromRemote(uuid string, params param.SyncAliyunVirtualRouterFromRemoteParam) (*view.SyncAliyunVirtualRouterFromRemoteEventView, error) {
	resp := view.SyncAliyunVirtualRouterFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{vpcUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

