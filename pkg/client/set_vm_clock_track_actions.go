// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmClockTrack 操作SetVmClockTrack
func (cli *ZSClient) SetVmClockTrack(uuid string, params param.SetVmClockTrackParam) (*view.SetVmClockTrackEventView, error) {
	resp := view.SetVmClockTrackEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

