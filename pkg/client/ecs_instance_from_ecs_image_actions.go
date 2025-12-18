// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEcsInstanceFromEcsImage 创建EcsInstanceFromEcsImage
func (cli *ZSClient) CreateEcsInstanceFromEcsImage(params param.CreateEcsInstanceFromEcsImageParam) (*view.CreateEcsInstanceFromEcsImageEventView, error) {
	resp := view.CreateEcsInstanceFromEcsImageEventView{}
	if err := cli.Post("v1/hybrid/aliyun/ecs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

