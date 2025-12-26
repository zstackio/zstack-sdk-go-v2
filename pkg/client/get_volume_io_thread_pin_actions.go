// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVolumeIoThreadPin gets VolumeIoThreadPin by uuid
func (cli *ZSClient) GetVolumeIoThreadPin(uuid string) (*view.GetVolumeIoThreadPinView, error) {
	var resp view.GetVolumeIoThreadPinView
	if err := cli.Get("v1/volumes/{uuid}/io-thread-pin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
