// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmCleanTraffic operates on SetVmCleanTraffic
func (cli *ZSClient) SetVmCleanTraffic(uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
