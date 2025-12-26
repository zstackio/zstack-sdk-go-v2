// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmSecurityLevel operates on SetVmSecurityLevel
func (cli *ZSClient) SetVmSecurityLevel(uuid string, params param.SetVmSecurityLevelParam) (*view.SetVmSecurityLevelEventView, error) {
	resp := view.SetVmSecurityLevelEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
