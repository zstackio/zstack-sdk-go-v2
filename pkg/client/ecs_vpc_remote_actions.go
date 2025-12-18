// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEcsVpcRemote 创建EcsVpcRemote
func (cli *ZSClient) CreateEcsVpcRemote(params param.CreateEcsVpcRemoteParam) (*view.CreateEcsVpcRemoteEventView, error) {
	resp := view.CreateEcsVpcRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

