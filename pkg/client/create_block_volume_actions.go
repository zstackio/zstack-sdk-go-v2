// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateBlockVolume creates BlockVolume
func (cli *ZSClient) CreateBlockVolume(params param.CreateBlockVolumeParam) (*view.CreateBlockVolumeEventView, error) {
	resp := view.CreateBlockVolumeEventView{}
	if err := cli.Post("v1/block-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
