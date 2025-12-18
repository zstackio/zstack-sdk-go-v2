// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetEcsInstanceVncUrl 获取EcsInstanceVncUrl详情
func (cli *ZSClient) GetEcsInstanceVncUrl(uuid string) (*view.GetEcsInstanceVncUrlView, error) {
	var resp view.GetEcsInstanceVncUrlView
	if err := cli.Get("v1/hybrid/aliyun/ecs-vnc/{uuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

