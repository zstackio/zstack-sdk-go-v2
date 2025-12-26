// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHost updates Host
func (cli *ZSClient) UpdateHost(uuid string, params param.UpdateHostParam) (*view.UpdateHostEventView, error) {
	resp := view.UpdateHostEventView{}
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
