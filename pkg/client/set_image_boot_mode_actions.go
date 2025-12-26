// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetImageBootMode operates on SetImageBootMode
func (cli *ZSClient) SetImageBootMode(uuid string, params param.SetImageBootModeParam) (*view.SetImageBootModeEventView, error) {
	resp := view.SetImageBootModeEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
