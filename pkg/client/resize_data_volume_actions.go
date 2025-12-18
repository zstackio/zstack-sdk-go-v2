// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResizeDataVolume operates on DataVolume
func (cli *ZSClient) ResizeDataVolume(uuid string, params param.ResizeDataVolumeParam) (*view.ResizeDataVolumeEventView, error) {
	resp := view.ResizeDataVolumeEventView{}
	if err := cli.Put("v1/volumes/data/resize/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
