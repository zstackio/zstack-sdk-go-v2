// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEcsInstanceFromEcsImage creates EcsInstanceFromEcsImage
func (cli *ZSClient) CreateEcsInstanceFromEcsImage(params param.CreateEcsInstanceFromEcsImageParam) (*view.CreateEcsInstanceFromEcsImageEventView, error) {
	resp := view.CreateEcsInstanceFromEcsImageEventView{}
	if err := cli.Post("v1/hybrid/aliyun/ecs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
