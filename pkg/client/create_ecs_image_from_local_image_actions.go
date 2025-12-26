// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEcsImageFromLocalImage creates EcsImageFromLocalImage
func (cli *ZSClient) CreateEcsImageFromLocalImage(params param.CreateEcsImageFromLocalImageParam) (*view.CreateEcsImageFromLocalImageEventView, error) {
	resp := view.CreateEcsImageFromLocalImageEventView{}
	if err := cli.Post("v1/hybrid/aliyun/image", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
