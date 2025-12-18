// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunVirtualRouter 更新AliyunVirtualRouter
func (cli *ZSClient) UpdateAliyunVirtualRouter(uuid string, params param.UpdateAliyunVirtualRouterParam) (*view.UpdateAliyunVirtualRouterEventView, error) {
	resp := view.UpdateAliyunVirtualRouterEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vrouter/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

