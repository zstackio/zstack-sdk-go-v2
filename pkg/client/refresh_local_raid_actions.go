// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshLocalRaid operates on RefreshLocalRaid
func (cli *ZSClient) RefreshLocalRaid(uuid string, params param.RefreshLocalRaidParam) (*view.RefreshLocalRaidEventView, error) {
	resp := view.RefreshLocalRaidEventView{}
	if err := cli.Put("v1/storage-devices/local-raid/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
