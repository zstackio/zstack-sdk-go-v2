// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVolumeCapabilities 获取VolumeCapabilities详情
func (cli *ZSClient) GetVolumeCapabilities(uuid string) (*view.GetVolumeCapabilitiesView, error) {
	var resp view.GetVolumeCapabilitiesView
	if err := cli.Get("v1/volumes/{uuid}/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

