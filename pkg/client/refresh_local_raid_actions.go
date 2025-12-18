// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshLocalRaid 操作RefreshLocalRaid
func (cli *ZSClient) RefreshLocalRaid(uuid string, params param.RefreshLocalRaidParam) (*view.RefreshLocalRaidEventView, error) {
	resp := view.RefreshLocalRaidEventView{}
	if err := cli.Put("v1/storage-devices/local-raid/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

