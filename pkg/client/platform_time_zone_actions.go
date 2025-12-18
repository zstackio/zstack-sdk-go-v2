// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPlatformTimeZone 获取PlatformTimeZone详情
func (cli *ZSClient) GetPlatformTimeZone(uuid string) (*view.GetPlatformTimeZoneView, error) {
	var resp view.GetPlatformTimeZoneView
	if err := cli.Get("v1/management-nodes/platform-timezone", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

