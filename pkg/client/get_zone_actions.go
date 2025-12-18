// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetZone gets Zone by uuid
func (cli *ZSClient) GetZone(uuid string) (*view.GetZoneView, error) {
	var resp view.GetZoneView
	if err := cli.Get("v1/zones/{uuid}/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
