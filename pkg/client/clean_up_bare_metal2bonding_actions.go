// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CleanUpBareMetal2Bonding operates on CleanUpBareMetal2Bonding
func (cli *ZSClient) CleanUpBareMetal2Bonding(uuid string, params param.CleanUpBareMetal2BondingParam) (*view.CleanUpBaremetal2BondingEventView, error) {
	resp := view.CleanUpBaremetal2BondingEventView{}
	if err := cli.Put("v1/baremetal2/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
