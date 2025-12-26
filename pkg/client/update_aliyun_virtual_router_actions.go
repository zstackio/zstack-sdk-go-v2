// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunVirtualRouter updates AliyunVirtualRouter
func (cli *ZSClient) UpdateAliyunVirtualRouter(uuid string, params param.UpdateAliyunVirtualRouterParam) (*view.UpdateAliyunVirtualRouterEventView, error) {
	resp := view.UpdateAliyunVirtualRouterEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
