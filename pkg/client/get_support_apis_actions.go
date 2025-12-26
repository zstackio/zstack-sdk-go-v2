// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSupportAPIs gets SupportAPIs by uuid
func (cli *ZSClient) GetSupportAPIs(uuid string) (*view.GetSupportAPIsView, error) {
	var resp view.GetSupportAPIsView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
