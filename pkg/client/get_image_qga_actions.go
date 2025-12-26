// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetImageQga gets ImageQga by uuid
func (cli *ZSClient) GetImageQga(uuid string) (*view.GetImageQgaView, error) {
	var resp view.GetImageQgaView
	if err := cli.Get("v1/images/{uuid}/qga", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
