// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVolumeFormat gets VolumeFormat by uuid
func (cli *ZSClient) GetVolumeFormat(uuid string) (*view.GetVolumeFormatView, error) {
	var resp view.GetVolumeFormatView
	if err := cli.Get("v1/volumes/formats", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
