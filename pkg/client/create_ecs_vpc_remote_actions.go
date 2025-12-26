// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEcsVpcRemote creates EcsVpcRemote
func (cli *ZSClient) CreateEcsVpcRemote(params param.CreateEcsVpcRemoteParam) (*view.CreateEcsVpcRemoteEventView, error) {
	resp := view.CreateEcsVpcRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/vpc", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
