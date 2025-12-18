// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetEventData gets EventData by uuid
func (cli *ZSClient) GetEventData(uuid string) (*view.GetEventDataView, error) {
	var resp view.GetEventDataView
	if err := cli.Get("v1/zwatch/events", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
