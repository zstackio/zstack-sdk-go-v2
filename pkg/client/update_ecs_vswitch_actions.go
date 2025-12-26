// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEcsVSwitch updates EcsVSwitch
func (cli *ZSClient) UpdateEcsVSwitch(uuid string, params param.UpdateEcsVSwitchParam) (*view.UpdateEcsVSwitchEventView, error) {
	resp := view.UpdateEcsVSwitchEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vswitch/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
