// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncVmClock operates on SyncVmClock
func (cli *ZSClient) SyncVmClock(uuid string, params param.SyncVmClockParam) (*view.SyncVmClockEventView, error) {
	resp := view.SyncVmClockEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
