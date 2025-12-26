// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateZone updates Zone
func (cli *ZSClient) UpdateZone(uuid string, params param.UpdateZoneParam) (*view.UpdateZoneEventView, error) {
	resp := view.UpdateZoneEventView{}
	if err := cli.Put("v1/zones/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
