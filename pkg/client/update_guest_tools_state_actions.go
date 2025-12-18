// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateGuestToolsState updates GuestToolsState
func (cli *ZSClient) UpdateGuestToolsState(uuid string, params param.UpdateGuestToolsStateParam) (*view.UpdateGuestToolsStateView, error) {
	resp := view.UpdateGuestToolsStateView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/guesttools-state", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
