// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeIPSecConnectionState changes IPSecConnectionState
func (cli *ZSClient) ChangeIPSecConnectionState(uuid string, params param.ChangeIPSecConnectionStateParam) (*view.ChangeIPSecConnectionStateEventView, error) {
	resp := view.ChangeIPSecConnectionStateEventView{}
	if err := cli.Put("v1/ipsec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
