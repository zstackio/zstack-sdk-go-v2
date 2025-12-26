// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSecurityMachineState changes SecurityMachineState
func (cli *ZSClient) ChangeSecurityMachineState(uuid string, params param.ChangeSecurityMachineStateParam) (*view.ChangeSecurityMachineStateEventView, error) {
	resp := view.ChangeSecurityMachineStateEventView{}
	if err := cli.Put("v1/security-machines/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
