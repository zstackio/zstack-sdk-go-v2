// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectBareMetal2Instance operates on ReconnectBareMetal2Instance
func (cli *ZSClient) ReconnectBareMetal2Instance(uuid string, params param.ReconnectBareMetal2InstanceParam) (*view.ReconnectBareMetal2InstanceEventView, error) {
	resp := view.ReconnectBareMetal2InstanceEventView{}
	if err := cli.Put("v1/baremetal2/bm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
