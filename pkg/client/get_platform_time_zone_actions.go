// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPlatformTimeZone gets PlatformTimeZone by uuid
func (cli *ZSClient) GetPlatformTimeZone(uuid string) (*view.GetPlatformTimeZoneView, error) {
	var resp view.GetPlatformTimeZoneView
	if err := cli.Get("v1/management-nodes/platform-timezone", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
