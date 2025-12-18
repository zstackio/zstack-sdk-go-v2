// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEcsVSwitch 更新EcsVSwitch
func (cli *ZSClient) UpdateEcsVSwitch(uuid string, params param.UpdateEcsVSwitchParam) (*view.UpdateEcsVSwitchEventView, error) {
	resp := view.UpdateEcsVSwitchEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vswitch/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

