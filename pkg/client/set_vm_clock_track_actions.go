// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmClockTrack operates on SetVmClockTrack
func (cli *ZSClient) SetVmClockTrack(uuid string, params param.SetVmClockTrackParam) (*view.SetVmClockTrackEventView, error) {
	resp := view.SetVmClockTrackEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
