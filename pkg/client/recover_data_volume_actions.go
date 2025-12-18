// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverDataVolume operates on DataVolume
func (cli *ZSClient) RecoverDataVolume(uuid string, params param.RecoverDataVolumeParam) (*view.RecoverDataVolumeEventView, error) {
	resp := view.RecoverDataVolumeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
