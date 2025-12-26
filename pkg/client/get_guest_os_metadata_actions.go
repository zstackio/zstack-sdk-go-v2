// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetGuestOsMetadata gets GuestOsMetadata by uuid
func (cli *ZSClient) GetGuestOsMetadata(uuid string) (*view.GetGuestOsMetadataView, error) {
	var resp view.GetGuestOsMetadataView
	if err := cli.Get("v1/guest-os/metadata", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
