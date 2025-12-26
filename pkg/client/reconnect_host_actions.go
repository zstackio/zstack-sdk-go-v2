// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectHost operates on ReconnectHost
func (cli *ZSClient) ReconnectHost(uuid string, params param.ReconnectHostParam) (*view.ReconnectHostEventView, error) {
	resp := view.ReconnectHostEventView{}
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
