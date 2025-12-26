// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeL3NetworkState changes L3NetworkState
func (cli *ZSClient) ChangeL3NetworkState(uuid string, params param.ChangeL3NetworkStateParam) (*view.ChangeL3NetworkStateEventView, error) {
	resp := view.ChangeL3NetworkStateEventView{}
	if err := cli.Put("v1/l3-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
