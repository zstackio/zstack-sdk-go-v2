// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEcsInstanceVncPassword updates EcsInstanceVncPassword
func (cli *ZSClient) UpdateEcsInstanceVncPassword(uuid string, params param.UpdateEcsInstanceVncPasswordParam) (*view.UpdateEcsInstanceVncPasswordEventView, error) {
	resp := view.UpdateEcsInstanceVncPasswordEventView{}
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/ecs-vnc", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
