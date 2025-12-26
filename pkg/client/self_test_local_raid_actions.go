// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SelfTestLocalRaid operates on SelfTestLocalRaid
func (cli *ZSClient) SelfTestLocalRaid(uuid string, params param.SelfTestLocalRaidParam) (*view.SelfTestLocalRaidEventView, error) {
	resp := view.SelfTestLocalRaidEventView{}
	if err := cli.Put("v1/storage-devices/local-raid/physical-drives/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
