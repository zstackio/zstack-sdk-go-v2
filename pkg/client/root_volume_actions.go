// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResizeRootVolume 操作RootVolume
func (cli *ZSClient) ResizeRootVolume(uuid string, params param.ResizeRootVolumeParam) (*view.ResizeRootVolumeEventView, error) {
	resp := view.ResizeRootVolumeEventView{}
	if err := cli.Put("v1/volumes/resize/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

