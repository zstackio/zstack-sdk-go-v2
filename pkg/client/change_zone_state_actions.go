// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeZoneState changes ZoneState
func (cli *ZSClient) ChangeZoneState(uuid string, params param.ChangeZoneStateParam) (*view.ChangeZoneStateEventView, error) {
	resp := view.ChangeZoneStateEventView{}
	if err := cli.Put("v1/zones/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
