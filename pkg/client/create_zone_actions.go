// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateZone creates Zone
func (cli *ZSClient) CreateZone(params param.CreateZoneParam) (*view.CreateZoneEventView, error) {
	resp := view.CreateZoneEventView{}
	if err := cli.Post("v1/zones", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
