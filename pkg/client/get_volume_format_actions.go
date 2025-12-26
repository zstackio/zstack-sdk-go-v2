// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVolumeFormat gets VolumeFormat by uuid
func (cli *ZSClient) GetVolumeFormat(uuid string) (*view.GetVolumeFormatView, error) {
	var resp view.GetVolumeFormatView
	if err := cli.Get("v1/volumes/formats", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
