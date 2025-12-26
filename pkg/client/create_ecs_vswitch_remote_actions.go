// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEcsVSwitchRemote creates EcsVSwitchRemote
func (cli *ZSClient) CreateEcsVSwitchRemote(params param.CreateEcsVSwitchRemoteParam) (*view.CreateEcsVSwitchRemoteEventView, error) {
	resp := view.CreateEcsVSwitchRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/vswitch", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
