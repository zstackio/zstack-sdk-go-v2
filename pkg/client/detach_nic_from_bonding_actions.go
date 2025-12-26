// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachNicFromBonding operates on NicFromBonding
func (cli *ZSClient) DetachNicFromBonding(uuid string, params param.DetachNicFromBondingParam) (*view.DetachNicFromBondingEventView, error) {
	resp := view.DetachNicFromBondingEventView{}
	if err := cli.Put("v1/hosts/bondings/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
