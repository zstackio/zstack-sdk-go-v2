// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetEcsInstanceVncUrl gets EcsInstanceVncUrl by uuid
func (cli *ZSClient) GetEcsInstanceVncUrl(uuid string) (*view.GetEcsInstanceVncUrlView, error) {
	var resp view.GetEcsInstanceVncUrlView
	if err := cli.Get("v1/hybrid/aliyun/ecs-vnc/{uuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
