// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PowerOffHost operates on PowerOffHost
func (cli *ZSClient) PowerOffHost(uuid string, params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.Put("v1/hosts/power-off/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
