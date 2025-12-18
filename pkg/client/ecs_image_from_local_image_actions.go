// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEcsImageFromLocalImage 创建EcsImageFromLocalImage
func (cli *ZSClient) CreateEcsImageFromLocalImage(params param.CreateEcsImageFromLocalImageParam) (*view.CreateEcsImageFromLocalImageEventView, error) {
	resp := view.CreateEcsImageFromLocalImageEventView{}
	if err := cli.Post("v1/hybrid/aliyun/image", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

