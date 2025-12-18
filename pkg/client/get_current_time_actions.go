// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCurrentTime gets CurrentTime by uuid
func (cli *ZSClient) GetCurrentTime(uuid string) (*view.GetCurrentTimeView, error) {
	var resp view.GetCurrentTimeView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
