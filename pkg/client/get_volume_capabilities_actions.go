// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVolumeCapabilities gets VolumeCapabilities by uuid
func (cli *ZSClient) GetVolumeCapabilities(uuid string) (*view.GetVolumeCapabilitiesView, error) {
	var resp view.GetVolumeCapabilitiesView
	if err := cli.Get("v1/volumes/{uuid}/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
