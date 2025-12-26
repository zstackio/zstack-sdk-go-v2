// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetEventData gets EventData by uuid
func (cli *ZSClient) GetEventData(uuid string) (*view.GetEventDataView, error) {
	var resp view.GetEventDataView
	if err := cli.Get("v1/zwatch/events", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
