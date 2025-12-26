// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeHostState changes HostState
func (cli *ZSClient) ChangeHostState(uuid string, params param.ChangeHostStateParam) (*view.ChangeHostStateEventView, error) {
	resp := view.ChangeHostStateEventView{}
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
