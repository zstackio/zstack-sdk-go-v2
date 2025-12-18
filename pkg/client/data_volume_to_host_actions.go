// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachDataVolumeToHost 操作DataVolumeToHost
func (cli *ZSClient) AttachDataVolumeToHost(params param.AttachDataVolumeToHostParam) (*view.AttachDataVolumeToHostEventView, error) {
	resp := view.AttachDataVolumeToHostEventView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

