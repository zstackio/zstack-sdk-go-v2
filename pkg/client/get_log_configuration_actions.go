// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLogConfiguration gets LogConfiguration by uuid
func (cli *ZSClient) GetLogConfiguration(uuid string) (*view.GetLogConfigurationView, error) {
	var resp view.GetLogConfigurationView
	if err := cli.Get("v1/log/configurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
