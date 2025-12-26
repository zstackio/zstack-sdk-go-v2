// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PowerOnHost operates on PowerOnHost
func (cli *ZSClient) PowerOnHost(uuid string, params param.PowerOnHostParam) (*view.PowerOnHostEventView, error) {
	resp := view.PowerOnHostEventView{}
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
