// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachDataVolumeToHost operates on DataVolumeToHost
func (cli *ZSClient) AttachDataVolumeToHost(params param.AttachDataVolumeToHostParam) (*view.AttachDataVolumeToHostEventView, error) {
	resp := view.AttachDataVolumeToHostEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
