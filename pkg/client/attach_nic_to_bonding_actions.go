// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachNicToBonding operates on NicToBonding
func (cli *ZSClient) AttachNicToBonding(uuid string, params param.AttachNicToBondingParam) (*view.AttachNicToBondingEventView, error) {
	resp := view.AttachNicToBondingEventView{}
	if err := cli.Put("v1/hosts/bondings/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
