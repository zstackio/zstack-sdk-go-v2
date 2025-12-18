// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateZone creates Zone
func (cli *ZSClient) CreateZone(params param.CreateZoneParam) (*view.CreateZoneEventView, error) {
	resp := view.CreateZoneEventView{}
	if err := cli.Post("v1/zones", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
