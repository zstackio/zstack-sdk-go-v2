// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateXskyBlockVolume 更新XskyBlockVolume
func (cli *ZSClient) UpdateXskyBlockVolume(uuid string, params param.UpdateXskyBlockVolumeParam) (*view.UpdateBlockVolumeEventView, error) {
	resp := view.UpdateBlockVolumeEventView{}
	if err := cli.Put("v1/xsky/block-volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

