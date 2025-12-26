// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVolumeQos gets VolumeQos by uuid
func (cli *ZSClient) GetVolumeQos(uuid string) (*view.GetVolumeQosView, error) {
	var resp view.GetVolumeQosView
	if err := cli.Get("v1/volumes/{uuid}/qos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
