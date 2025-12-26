// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetEcsInstanceType gets EcsInstanceType by uuid
func (cli *ZSClient) GetEcsInstanceType(uuid string) (*view.GetEcsInstanceTypeView, error) {
	var resp view.GetEcsInstanceTypeView
	if err := cli.Get("v1/hybrid/ecs/type", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
