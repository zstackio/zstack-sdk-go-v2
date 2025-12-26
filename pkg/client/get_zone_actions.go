// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetZone gets Zone by uuid
func (cli *ZSClient) GetZone(uuid string) (*view.GetZoneView, error) {
	var resp view.GetZoneView
	if err := cli.Get("v1/zones/{uuid}/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
