// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StartBareMetal2Instance starts BareMetal2Instance
func (cli *ZSClient) StartBareMetal2Instance(uuid string, params param.StartBareMetal2InstanceParam) (*view.StartBareMetal2InstanceEventView, error) {
	resp := view.StartBareMetal2InstanceEventView{}
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/action", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
