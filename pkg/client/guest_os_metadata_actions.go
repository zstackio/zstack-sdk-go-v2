// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetGuestOsMetadata 获取GuestOsMetadata详情
func (cli *ZSClient) GetGuestOsMetadata(uuid string) (*view.GetGuestOsMetadataView, error) {
	var resp view.GetGuestOsMetadataView
	if err := cli.Get("v1/guest-os/metadata", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

