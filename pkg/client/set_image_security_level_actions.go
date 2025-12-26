// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetImageSecurityLevel operates on SetImageSecurityLevel
func (cli *ZSClient) SetImageSecurityLevel(uuid string, params param.SetImageSecurityLevelParam) (*view.SetImageSecurityLevelEventView, error) {
	resp := view.SetImageSecurityLevelEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
