// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetExternalServices gets ExternalServices by uuid
func (cli *ZSClient) GetExternalServices(uuid string) (*view.GetExternalServicesView, error) {
	var resp view.GetExternalServicesView
	if err := cli.Get("v1/external/services", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
