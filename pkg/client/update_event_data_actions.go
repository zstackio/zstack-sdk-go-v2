// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEventData updates EventData
func (cli *ZSClient) UpdateEventData(uuid string, params param.UpdateEventDataParam) (*view.UpdateEventDataEventView, error) {
	resp := view.UpdateEventDataEventView{}
	if err := cli.Put("v1/zwatch/events/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
