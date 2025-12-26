// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RebootEcsInstance operates on EcsInstance
func (cli *ZSClient) RebootEcsInstance(uuid string, params param.RebootEcsInstanceParam) (*view.RebootEcsInstanceEventView, error) {
	resp := view.RebootEcsInstanceEventView{}
	if err := cli.Put("v1/hybrid/aliyun/ecs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
