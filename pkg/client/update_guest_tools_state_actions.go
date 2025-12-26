// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(uuid string, params param.UpdateGuestToolsStateParam) (*view.UpdateGuestToolsStateView, error) {
	resp := view.UpdateGuestToolsStateView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/guesttools-state", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
