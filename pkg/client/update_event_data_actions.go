// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEventData updates EventData
func (cli *ZSClient) UpdateEventData(uuid string, params param.UpdateEventDataParam) (*view.UpdateEventDataEventView, error) {
	resp := view.UpdateEventDataEventView{}
	if err := cli.Put("v1/zwatch/events/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
