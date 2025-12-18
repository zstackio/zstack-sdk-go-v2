// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEcsSecurityGroup 更新EcsSecurityGroup
func (cli *ZSClient) UpdateEcsSecurityGroup(uuid string, params param.UpdateEcsSecurityGroupParam) (*view.UpdateEcsSecurityGroupEventView, error) {
	resp := view.UpdateEcsSecurityGroupEventView{}
	if err := cli.Put("v1/hybrid/aliyun/security-group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

