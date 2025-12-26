// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StartEcsInstance starts EcsInstance
func (cli *ZSClient) StartEcsInstance(uuid string, params param.StartEcsInstanceParam) (*view.StartEcsInstanceEventView, error) {
	resp := view.StartEcsInstanceEventView{}
	if err := cli.Put("v1/hybrid/aliyun/ecs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
