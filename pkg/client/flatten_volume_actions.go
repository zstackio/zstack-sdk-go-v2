// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// FlattenVolume operates on FlattenVolume
func (cli *ZSClient) FlattenVolume(uuid string, params param.FlattenVolumeParam) (*view.FlattenVolumeEventView, error) {
	resp := view.FlattenVolumeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
