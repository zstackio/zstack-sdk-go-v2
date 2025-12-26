// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVSwitchTypes gets VSwitchTypes by uuid
func (cli *ZSClient) GetVSwitchTypes(uuid string) (*view.GetVSwitchTypesView, error) {
	var resp view.GetVSwitchTypesView
	if err := cli.Get("v1/l2-networks/vSwitchTypes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
