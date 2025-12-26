// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StopEcsInstance stops EcsInstance
func (cli *ZSClient) StopEcsInstance(uuid string, params param.StopEcsInstanceParam) (*view.StopEcsInstanceEventView, error) {
	resp := view.StopEcsInstanceEventView{}
	if err := cli.Put("v1/hybrid/aliyun/ecs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
