// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverDataVolume operates on DataVolume
func (cli *ZSClient) RecoverDataVolume(uuid string, params param.RecoverDataVolumeParam) (*view.RecoverDataVolumeEventView, error) {
	resp := view.RecoverDataVolumeEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
