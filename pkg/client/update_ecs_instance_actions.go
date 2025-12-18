// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEcsInstance updates EcsInstance
func (cli *ZSClient) UpdateEcsInstance(uuid string, params param.UpdateEcsInstanceParam) (*view.UpdateEcsInstanceEventView, error) {
	resp := view.UpdateEcsInstanceEventView{}
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/ecs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
