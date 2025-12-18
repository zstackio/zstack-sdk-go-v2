// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEcsInstanceVncPassword 更新EcsInstanceVncPassword
func (cli *ZSClient) UpdateEcsInstanceVncPassword(uuid string, params param.UpdateEcsInstanceVncPasswordParam) (*view.UpdateEcsInstanceVncPasswordEventView, error) {
	resp := view.UpdateEcsInstanceVncPasswordEventView{}
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/ecs-vnc", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

