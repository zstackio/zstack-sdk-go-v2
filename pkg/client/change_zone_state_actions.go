// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeZoneState changes ZoneState
func (cli *ZSClient) ChangeZoneState(uuid string, params param.ChangeZoneStateParam) (*view.ChangeZoneStateEventView, error) {
	resp := view.ChangeZoneStateEventView{}
	if err := cli.Put("v1/zones/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
